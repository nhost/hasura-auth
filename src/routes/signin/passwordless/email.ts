import { RequestHandler } from 'express';
import { v4 as uuidv4 } from 'uuid';
import { ReasonPhrases } from 'http-status-codes';

import {
  getUserByEmail,
  getGravatarUrl,
  generateTicketExpiresAt,
  ENV,
  createEmailRedirectionLink,
} from '@/utils';
import { emailClient } from '@/email';
import { DBUser, EMAIL_TYPES, PasswordLessEmailBody } from '@/types';
import { sendError } from '@/errors';
import { Joi, email, registrationOptions } from '@/validation';
import { postgres } from '@/utils/postgres';

export const signInPasswordlessEmailSchema = Joi.object({
  email: email.required(),
  options: registrationOptions,
}).meta({ className: 'SignInPasswordlessEmailSchema' });

export const signInPasswordlessEmailHandler: RequestHandler<
  {},
  {},
  PasswordLessEmailBody
> = async (req, res) => {
  if (!ENV.AUTH_EMAIL_PASSWORDLESS_ENABLED) {
    return sendError(res, 'disabled-endpoint');
  }

  const {
    email,
    options: {
      redirectTo,
      defaultRole,
      allowedRoles,
      displayName,
      locale,
      metadata,
    },
  } = req.body;

  // check if user with email already exist
  let user = await getUserByEmail(email);

  // if no user exists, create the user
  if (!user) {
    // insert user
    const { rows } = await postgres.runSqlParsed(
      `INSERT INTO auth.users (
        disabled, 
        display_name, 
        avatar_url,
        email,
        email_verified,
        locale,
        default_role,
        metadata
      ) VALUES (%L, %L, %L, %L, %L, %L, %L, %L) RETURNING id`,
      [
        ENV.AUTH_DISABLE_NEW_USERS,
        displayName || email,
        getGravatarUrl(email),
        email,
        false,
        locale,
        defaultRole,
        metadata,
      ]
    );

    if (!Array.isArray(rows[0])) {
      return;
    }

    // get inserted user id
    const insertedUserId = rows[0][0];

    // insert allowed roles for the user
    for (const role of allowedRoles) {
      await postgres.runSql(
        `INSERT INTO auth.user_roles (user_id, role) VALUES (%L, %L)`,
        [insertedUserId, role]
      );
    }

    // get users
    const { rows: users } = await postgres.runSqlParsed(
      `SELECT row_to_json(u) FROM auth.users u WHERE id = %L LIMIT 1`,
      [insertedUserId]
    );

    // this should never happen
    if (users.length === 0) {
      // TODO: better error
      console.log('user not found after insert');
      throw new Error('User not found');
    }

    user = users[0] as DBUser;
  }

  // Now, a user should already exist or we just created one.

  // this should never happen. Either there was an existing user or we just created one
  if (!user) {
    throw Error('no user');
  }

  if (user?.disabled) {
    return sendError(res, 'disabled-user');
  }

  // create ticket for magic link
  const ticket = `passwordlessEmail:${uuidv4()}`;
  const ticketExpiresAt = generateTicketExpiresAt(60 * 60);

  await postgres.runSql(
    `UPDATE auth.users SET ticket = %L, ticket_expires_at = %L WHERE id = %L`,
    [ticket, ticketExpiresAt, user.id]
  );

  const template = 'signin-passwordless';
  const link = createEmailRedirectionLink(
    EMAIL_TYPES.SIGNIN_PASSWORDLESS,
    ticket,
    redirectTo
  );
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
        'x-link': {
          prepared: true,
          value: link,
        },
      },
    },
    locals: {
      link,
      displayName: user.display_name,
      email,
      ticket,
      redirectTo: encodeURIComponent(redirectTo),
      locale: user.locale ?? ENV.AUTH_LOCALE_DEFAULT,
      serverUrl: ENV.AUTH_SERVER_URL,
      clientUrl: ENV.AUTH_CLIENT_URL,
    },
  });

  return res.send(ReasonPhrases.OK);
};
