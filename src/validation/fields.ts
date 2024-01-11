import { UserRegistrationOptionsWithRedirect } from '@/types';
import { pwnedPassword } from 'hibp';
import { isValidPhoneNumber, parsePhoneNumber } from 'libphonenumber-js';
import micromatch from 'micromatch';

import { ENV } from '../utils/env';

import { Joi } from './joi';
import { EmailValidator } from './validators';

export const uuidRegex =
  /\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b/;

export const password = Joi.string().example('Str0ngPassw#ord-94|%');

export const passwordInsert = password
  .min(ENV.AUTH_PASSWORD_MIN_LENGTH)
  .description(
    `A password of minimum ${ENV.AUTH_PASSWORD_MIN_LENGTH} characters`
  )
  .external(async (value) => {
    if (ENV.AUTH_PASSWORD_HIBP_ENABLED && (await pwnedPassword(value))) {
      const message = 'Password is too weak (it has been pwned)';
      throw new Joi.ValidationError(
        message,
        [
          {
            message,
            path: ['password'],
            type: 'string.min',
          },
        ],
        value
      );
    } else return value;
  }, `When HIBP is enabled, will check if the password is too weak`);

export const email = Joi.string()
  .email()
  .custom(EmailValidator)
  .example('john.smith@nhost.io')
  .description('A valid email');

export const locale = Joi.string()
  .length(2)
  // TODO reactivate ENV.AUTH_LOCALE_ALLOWED_LOCALES check
  // * For the moment, the Nhost console does not allow selecting activated locales.
  // * Once it will be possible, we will need to reactivate this check to make sure an undeclared locale falls back to the default one, but tries it all the way
  // * As an example: if 'fr' is not in the allowed locales, it is currently passed on all the way.
  // * In reactivating the following code, it will then fall back to the default locale (en) and for instance use the english email templates
  // .valid(...ENV.AUTH_LOCALE_ALLOWED_LOCALES)
  // .failover(ENV.AUTH_LOCALE_DEFAULT)
  .default(ENV.AUTH_LOCALE_DEFAULT)
  .example(ENV.AUTH_LOCALE_DEFAULT)
  .description(`A two-characters locale`);

export const defaultRole = Joi.string()
  .default(ENV.AUTH_USER_DEFAULT_ROLE)
  .example(ENV.AUTH_USER_DEFAULT_ROLE)
  .valid(...ENV.AUTH_USER_DEFAULT_ALLOWED_ROLES);

export const allowedRoles = Joi.array()
  .items(...ENV.AUTH_USER_DEFAULT_ALLOWED_ROLES)
  .default(ENV.AUTH_USER_DEFAULT_ALLOWED_ROLES)
  .example(ENV.AUTH_USER_DEFAULT_ALLOWED_ROLES);

export const displayName = Joi.string().example('John Smith');

export const metadata = Joi.object().default({}).example({
  firstName: 'John',
  lastName: 'Smith',
});

export const redirectTo = Joi.string()
  .messages({
    redirectTo: 'The value of {{#label}} is not allowed.',
  })
  .default(ENV.AUTH_CLIENT_URL)
  .custom((value, helper) => {
    // * If no client url is set, we allow any valid url
    if (!ENV.AUTH_CLIENT_URL) {
      try {
        new URL(value);
        return value;
      } catch {
        return helper.error('redirectTo');
      }
    }

    // * We allow any sub-path of the client url
    // * With optional hash and query params
    if (
      new RegExp(`^${ENV.AUTH_CLIENT_URL}(\/.*)?([?].*)?([#].*)?$`).test(value)
    ) {
      return value;
    }

    // * Check if the value's hostname matches any allowed hostname
    // * Required to avoid shadowing domains
    const hostnames = ENV.AUTH_ACCESS_CONTROL_ALLOWED_REDIRECT_URLS.map(
      (allowed) => {
        return new URL(allowed).hostname;
      }
    );

    const valueUrl = new URL(value);
    if (!micromatch.isMatch(valueUrl.hostname, hostnames, { nocase: true })) {
      return helper.error('redirectTo');
    }

    // * We allow any sub-path of the allowed redirect urls.
    // * Allowed redirect urls also accepts wildcards and other micromatch patterns
    const expressions = ENV.AUTH_ACCESS_CONTROL_ALLOWED_REDIRECT_URLS.map(
      (allowed) => {
        // * Replace all the `.` by `/` so micromatch will understand `.` as a path separator
        allowed = allowed.replace(/[.]/g, '/');
        // * Append `/**` to the end of the allowed URL to allow for any subpath
        if (allowed.endsWith('/**')) {
          return allowed;
        }
        if (allowed.endsWith('/*')) {
          return `${allowed}*`;
        }
        if (allowed.endsWith('/')) {
          return `${allowed}**`;
        }
        return `${allowed}/**`;
      }
    );

    try {
      // * Don't take the query parameters into account
      // * And replace `.` with `/` because of micromatch
      const urlWithoutParams = `${valueUrl.origin}${valueUrl.pathname}`.replace(
        /[.]/g,
        '/'
      );
      const match = micromatch.isMatch(urlWithoutParams, expressions, {
        nocase: true,
      });
      if (match) return value;
      return helper.error('redirectTo');
    } catch {
      // * value is not a valid URL
      return helper.error('redirectTo');
    }
  })
  .example(`${ENV.AUTH_CLIENT_URL}/catch-redirection`);

export const uuid = Joi.string()
  .regex(uuidRegex)
  .example('2c35b6f3-c4b9-48e3-978a-d4d0f1d42e24')
  .description('A valid UUID');

export const jwt = Joi.string()
  .regex(/^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_.+/=]*$/)
  .example(
    'eyJhbGciOiJIUzI1NiJ9.eyJodHRwczovL2hhc3VyYS5pby9qd3QvY2xhaW1zIjp7IngtaGFzdXJhLWFsbG93ZWQtcm9sZXMiOlsibWUiLCJ1c2VyIl0sIngtaGFzdXJhLWRlZmF1bHQtcm9sZSI6InVzZXIiLCJ4LWhhc3VyYS11c2VyLWlkIjoiODAwYjA2ZWYtNGMyYi00NjQwLWIyMjAtNWZlNjk3ZWNjZGM2IiwieC1oYXN1cmEtdXNlci1pcy1hbm9ueW1vdXMiOiJmYWxzZSJ9LCJzdWIiOiI4MDBiMDZlZi00YzJiLTQ2NDAtYjIyMC01ZmU2OTdlY2NkYzYiLCJpc3MiOiJoYXN1cmEtYXV0aCIsImlhdCI6MTY1MTg2NTkwMCwiZXhwIjoxNjUxODY2ODAwfQ.IvFIMXOe6J21fyEfPkP9Caim3C_uAD2qimK4oGpNm44'
  )
  .description('A valid JWT token');

export const userId = uuid.description('Id of the user');

export const refreshToken = uuid
  .required()
  .description(
    'Refresh token during authentication or when refreshing the JWT'
  );

export const personalAccessToken = uuid
  .required()
  .description('Personal access token');

export const token = jwt.optional().description('Access token');

export const registrationOptions =
  Joi.object<UserRegistrationOptionsWithRedirect>({
    locale,
    defaultRole,
    allowedRoles,
    displayName,
    metadata,
    redirectTo,
  })
    .default()
    .custom((value, helper) => {
      const { allowedRoles, defaultRole } = value;
      if (!allowedRoles.includes(defaultRole)) {
        return helper.error('Default role must be part of allowed roles');
      }
      // check if allowedRoles is a subset of allowed user roles
      if (
        !allowedRoles.every((role: string) =>
          ENV.AUTH_USER_DEFAULT_ALLOWED_ROLES.includes(role)
        )
      ) {
        return helper.error('Allowed roles must be a subset of allowedRoles');
      }
      return value;
    });

export const mfaTotpTicketPattern = new RegExp(`mfaTotp:${uuidRegex.source}`);

export const activeMfaType = Joi.alternatives()
  .try(
    Joi.string().valid('').empty('').default(null), // accept only empty strings and convert those to null
    Joi.string().valid('totp')
  )
  .example('totp')
  .description(
    'Multi-factor authentication type. A null value deactivates MFA'
  );

export const phoneNumber = Joi.string()
  .required()
  .custom((value: string) => {
    // * Replace '00' prefix by '+'
    if (value.startsWith('00')) {
      value = value.replace('00', '+');
    }
    if (isValidPhoneNumber(value)) {
      return parsePhoneNumber(value).format('E.164');
    } else {
      throw new Error('invalid phone number');
    }
  }, 'valid phone number');
