import { Response } from 'express';
import {
  ContainerTypes,
  ValidatedRequest,
  ValidatedRequestSchema,
} from 'express-joi-validation';
import { v4 as uuidv4 } from 'uuid';

import { REGISTRATION } from '@config/registration';
import { getGravatarUrl, getUserByEmail, hashPassword } from '@/helpers';
import { gqlSdk } from '@/utils/gqlSDK';
import { AUTHENTICATION } from '@config/authentication';
import { APPLICATION } from '@config/application';
import { emailClient } from '@/email';
import { insertProfile, isProfileValid } from '@/utils/profile';
import { isValidEmail } from '@/utils/email';
import { isPasswordValid } from '@/utils/password';
import { isRolesValid } from '@/utils/roles';
import { ENV } from '@/utils/env';

type Profile = {
  [key: string]: string | number | boolean;
};

type BodyType = {
  email: string;
  password: string;
  locale: string;
  allowedRoles: string[] | null;
  defaultRole: string | null;
  displayName: string | null;
  profile: Profile | null;
};

interface Schema extends ValidatedRequestSchema {
  [ContainerTypes.Body]: BodyType;
}

export const signUpEmailPasswordHandler = async (
  req: ValidatedRequest<Schema>,
  res: Response
): Promise<unknown> => {
  const { body } = req;
  const { email, password, profile, locale = ENV.DEFAULT_LOCALE } = body;

  // check email
  if (!(await isValidEmail({ email, res }))) {
    // function send potential error via `res`
    return;
  }

  // check password
  if (!(await isPasswordValid({ password, res }))) {
    // function send potential error via `res`
    return;
  }

  // check profile
  if (!(await isProfileValid({ profile, res }))) {
    // function send potential error via `res`
    return;
  }

  // check roles
  const defaultRole = body.defaultRole ?? ENV.DEFAULT_USER_ROLE;
  const allowedRoles = body.allowedRoles ?? ENV.DEFAULT_ALLOWED_USER_ROLES;
  if (!(await isRolesValid({ defaultRole, allowedRoles, res }))) {
    return;
  }

  // check if email already in use by some other user
  if (await getUserByEmail(email)) {
    return res.boom.conflict('Email already in use');
  }

  // hash password
  const passwordHash = await hashPassword(password);

  // create ticket
  const ticket = `userActivate:${uuidv4()}`;
  const ticketExpiresAt = new Date(+new Date() + 60 * 60 * 1000).toISOString();

  // restructure user roles to be inserted in GraphQL mutation
  const userRoles = allowedRoles.map((role: string) => ({ role }));

  const displayName = body.displayName ?? email;
  const avatarUrl = getGravatarUrl(email);

  // insert user
  const user = await gqlSdk
    .insertUser({
      user: {
        displayName,
        avatarUrl,
        email,
        passwordHash,
        ticket,
        ticketExpiresAt,
        isActive: REGISTRATION.AUTO_ACTIVATE_NEW_USERS,
        emailVerified: false,
        locale,
        defaultRole,
        roles: {
          data: userRoles,
        },
      },
    })
    .then((res) => res.insertUser);

  if (!user) {
    throw new Error('Unable to insert new user');
  }

  await insertProfile({ userId: user.id, profile });

  // user is now inserted. Continue sending out activation email
  if (!REGISTRATION.AUTO_ACTIVATE_NEW_USERS && AUTHENTICATION.VERIFY_EMAILS) {
    if (!APPLICATION.EMAILS_ENABLED) {
      throw new Error('SMTP settings unavailable');
    }

    await emailClient.send({
      template: 'activate-user',
      message: {
        to: email,
        headers: {
          'x-ticket': {
            prepared: true,
            value: ticket,
          },
        },
      },
      locals: {
        displayName,
        ticket,
        url: APPLICATION.SERVER_URL,
        locale: user.locale,
      },
    });
  }

  return res.status(200).send('OK');
};