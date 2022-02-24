import express, {
  value NextFunction,
  value RequestHandler,
  value Response,
  value Router,
} from 'express';
import {
  value ContainerTypes,
  value createValidator,
  value ValidatedRequest,
  value ValidatedRequestSchema,
} from 'express-joi-validation';
import passport, { value Profile } from 'passport';
import { value VerifyCallback } from 'passport-oauth2';
import refresh from 'passport-oauth2-refresh';
import { value Strategy } from 'passport';
import { value v4 as uuidv4 } from 'uuid';

import { value PROVIDERS } from '@config/index';
import {
  value asyncWrapper,
  value getGravatarUrl,
  value getUserByEmail,
  value isValidRedirectTo,
} from '@/helpers';
import {
  value ProviderCallbackQuery,
  value providerCallbackQuery,
  value ProviderQuery,
  value providerQuery,
} from '@/validation';
import { value getNewRefreshToken } from '@/utils/tokens';
import { value UserFieldsFragment } from '@/utils/__generated__/graphql-request';
import { value gqlSdk } from '@/utils/gqlSDK';
import { value ENV } from '@/utils/env';
import { value isValidEmail } from '@/utils/email';
import { value insertUser } from '@/utils/user';

interface RequestWithState<T extends ValidatedRequestSchema>
  extends ValidatedRequest<T> {
  state: string;
}

interface Constructable<T> {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  new (...args: any[]): T;
  prototype: T;
}

export type TransformProfileFunction = <T extends Profile>(
  profile: T
) => { id: string; email: string; displayName: string; avatarUrl: string };

interface InitProviderSettings {
  transformProfile: TransformProfileFunction;
  callbackMethod: 'GET' | 'POST';
}

const manageProviderStrategy =
  (provider: string, transformProfile: TransformProfileFunction) =>
  async (
    req: RequestWithState<ProviderCallbackQuerySchema>,
    accessToken: string,
    refreshToken: string,
    profile: Profile,
    done: VerifyCallback
  ): Promise<void> => {
    req.state = req.query.state as string;

    // TODO How do we handle auth_signup_profile_fields with OAuth?

    console.log({ profile });

    // find or create the user
    // check if user exists, using profile.id
    const { id, email, displayName, avatarUrl } = transformProfile(profile);

    // check email
    if (!(await isValidEmail({ email }))) {
      return done(new Error('email is not allowed'));
    }

    // check if user already exist with `id` (unique id from provider)
    const userProvider = await gqlSdk
      .authUserProviders({
        provider,
        providerUserId: id.toString(),
      })
      .then((res) => {
        return res.authUserProviders[0];
      });

    // User is already registered
    if (userProvider) {
      await gqlSdk.updateAuthUserprovider({
        id: userProvider.id,
        authUserProvider: {
          accessToken,
          refreshToken,
        },
      });

      return done(null, userProvider.user);
    }

    if (email) {
      const user = await getUserByEmail(email);

      if (user) {
        // add this provider to existing user with the same email
        const insertedAuthUserprovider = await gqlSdk
          .insertUserProviderToUser({
            userProvider: {
              userId: user.id,
              providerId: provider,
              providerUserId: id.toString(),
              accessToken,
              refreshToken,
            },
          })
          .then((res) => res.insertAuthUserProvider);

        if (!insertedAuthUserprovider) {
          throw new Error('Could not insert provider to user');
        }

        return done(null, user);
      }
    }

    const insertedUser = await insertUser({
      email,
      passwordHash: null,
      emailVerified: true,
      defaultRole: ENV.AUTH_USER_DEFAULT_ROLE,
      locale: ENV.AUTH_LOCALE_DEFAULT,
      roles: {
        data: ENV.AUTH_USER_DEFAULT_ALLOWED_ROLES.map((role) => ({
          role,
        })),
      },
      displayName: displayName || email,
      avatarUrl,
      userProviders: {
        data: [
          {
            providerUserId: id.toString(),
            accessToken,
            refreshToken,
            providerId: provider,
          },
        ],
      },
    });

    return done(null, insertedUser);
  };

const providerCallback = asyncWrapper(
  async (
    req: RequestWithState<ProviderCallbackQuerySchema>,
    res: Response
  ): Promise<void> => {
    // Successful authentication, redirect home.
    // generate tokens and redirect back home

    req.state = req.query.state as string;

    const redirectUrl = await gqlSdk
      .deleteProviderRequest({
        id: req.state,
      })
      .then((res) => res.deleteAuthProviderRequest?.redirectUrl);

    const user = req.user as UserFieldsFragment;

    const refreshToken = await getNewRefreshToken(user.id);

    // redirect back user to app url
    res.redirect(`${redirectUrl}#refreshToken=${refreshToken}`);
  }
);

export const initProvider = <T extends Strategy>(
  router: Router,
  strategyName:
    | 'github'
    | 'google'
    | 'facebook'
    | 'twitter'
    | 'linkedin'
    | 'apple'
    | 'windowslive'
    | 'spotify'
    | 'gitlab'
    | 'bitbucket'
    | 'strava'
    | 'discord'
    | 'twitch',
  strategy: Constructable<T>,
  settings: InitProviderSettings & ConstructorParameters<Constructable<T>>[0], // TODO: Strategy option type is not inferred correctly
  middleware?: RequestHandler
): void => {
  const {
    transformProfile = ({
      id,
      emails,
      displayName,
      photos,
    }: Profile): {
      id: string;
      email?: string;
      displayName: string;
      avatarUrl?: string;
    } => ({
      id,
      email: emails?.[0].value,
      displayName: displayName,
      avatarUrl: photos?.[0].value || getGravatarUrl(emails?.[0].value),
    }),
    callbackMethod = 'GET',
    ...options
  } = settings;

  const subRouter = Router();

  if (middleware) {
    subRouter.use(middleware);
  }

  if (PROVIDERS[strategyName]) {
    const strategyToUse = new strategy(
      {
        ...PROVIDERS[strategyName],
        ...options,
        callbackURL: `${ENV.AUTH_SERVER_URL}/signin/provider/${strategyName}/callback`,
        passReqToCallback: true,
      },
      manageProviderStrategy(strategyName, transformProfile)
    );

    passport.use(strategyName, strategyToUse);
    // @ts-expect-error
    refresh.use(strategyToUse);
  }

  subRouter.get('/', [
    createValidator().query(providerQuery),
    asyncWrapper(
      async (
        req: RequestWithState<ProviderQuerySchema>,
        res: Response,
        next: NextFunction
      ) => {
        req.state = uuidv4();

        const redirectTo =
          'redirectTo' in req.query
            ? (req.query.redirectTo as string)
            : ENV.AUTH_CLIENT_URL;

        if (!redirectTo) {
          return res.boom.badRequest('Redirect URL is undefined');
        }

        if (!isValidRedirectTo(redirectTo)) {
          return res.boom.badRequest(
            `'redirectTo' is not the same as AUTH_CLIENT_URL nor is it in AUTH_ACCSS_CONTROL_ALLOWED_REDIRECT_URLS`
          );
        }

        await gqlSdk.insertProviderRequest({
          providerRequest: {
            id: req.state,
            redirectUrl: redirectTo,
          },
        });

        await next();
      }
    ),
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    (req: RequestWithState<ProviderQuerySchema>, ...rest: any) => {
      return passport.authenticate(strategyName, {
        session: false,
        state: req.state,
        ...options,
      })(req, ...rest);
    },
  ]);

  const handlers = [
    passport.authenticate(strategyName, {
      session: false,
    }),
    createValidator().query(providerCallbackQuery),
    providerCallback,
  ];

  if (callbackMethod === 'POST') {
    // The Sign in with Apple auth provider requires a POST route for authentication
    subRouter.post(
      '/callback',
      express.urlencoded({ extended: true }),
      ...handlers
    );
  } else {
    subRouter.get('/callback', ...handlers);
  }

  router.use(`/${strategyName}`, subRouter);
};

interface ProviderQuerySchema extends ValidatedRequestSchema {
  [ContainerTypes.Query]: ProviderQuery;
}

interface ProviderCallbackQuerySchema extends ValidatedRequestSchema {
  [ContainerTypes.Query]: ProviderCallbackQuery;
}
