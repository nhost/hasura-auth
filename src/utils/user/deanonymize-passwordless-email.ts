import { Response } from 'express';
import { v4 as uuidv4 } from 'uuid';

import { gqlSdk } from '../gqlSDK';
import { getUserByEmail } from '@/helpers';
import { ENV } from '../env';
import { emailClient } from '@/email';
import { generateTicketExpiresAt } from '../ticket';

export type BodyTypePasswordlessEmail = {
  signInMethod: 'passwordless';
  connection: 'email';
  email: string;
  options: {
    allowedRoles: string[];
    defaultRole: string;
    redirectTo: string;
  };
};

export const handleDeanonymizeUserPasswordlessEmail = async (
  body: BodyTypePasswordlessEmail,
  userId: string,
  res: Response
): Promise<unknown> => {
  const { user } = await gqlSdk.user({
    id: userId,
  });

  if (user?.isAnonymous !== true) {
    return res.boom.badRequest('Logged in user is not anonymous');
  }

  const {
    email,
    options: { redirectTo, defaultRole, allowedRoles },
  } = body;

  // check if email already in use by some other user
  if (await getUserByEmail(email)) {
    return res.boom.conflict('Email already in use');
  }

  // delete existing (anonymous) user roles
  await gqlSdk.deleteUserRolesByUserId({
    userId,
  });

  // insert new user roles (userRoles)
  await gqlSdk.insertUserRoles({
    userRoles: allowedRoles.map((role: string) => ({ role, userId })),
  });

  const ticket = `verifyEmail:${uuidv4()}`;
  const ticketExpiresAt = generateTicketExpiresAt(60 * 60);

  await gqlSdk.updateUser({
    id: userId,
    user: {
      emailVerified: false,
      email,
      defaultRole,
      ticket,
      ticketExpiresAt,
      isAnonymous: false,
    },
  });

  // Delete old refresh token and send email if email must be verified
  if (ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED) {
    // delete old refresh tokens for user
    await gqlSdk.deleteUserRefreshTokens({
      userId,
    });
    // create ticket
    const ticket = `passwordlessEmail:${uuidv4()}`;
    const ticketExpiresAt = generateTicketExpiresAt(60 * 60);

    await gqlSdk.updateUser({
      id: userId,
      user: {
        ticket,
        ticketExpiresAt,
      },
    });

    const template = 'signin-passwordless';
    await emailClient.send({
      template,
      message: {
        to: email,
        headers: {
          'x-ticket': {
            prepared: true,
            value: ticket,
          },
          'x-redirect-to': {
            prepared: true,
            value: redirectTo,
          },
          'x-email-template': {
            prepared: true,
            value: template,
          },
        },
      },
      locals: {
        link: `${ENV.AUTH_SERVER_URL}/verify?&ticket=${ticket}&type=signinPasswordless&redirectTo=${redirectTo}`,
        displayName: user.displayName,
        email,
        ticket,
        redirectTo,
        locale: user.locale ?? ENV.AUTH_LOCALE_DEFAULT,
        serverUrl: ENV.AUTH_SERVER_URL,
        clientUrl: ENV.AUTH_CLIENT_URL,
      },
    });
  }
};
