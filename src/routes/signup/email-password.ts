import { RequestHandler } from 'express';
import { v4 as uuidv4 } from 'uuid';

import { emailClient } from '@/email';
import {
  getGravatarUrl,
  hashPassword,
  generateTicketExpiresAt,
  getSignInResponse,
  ENV,
  createEmailRedirectionLink,
  getUserByEmail,
} from '@/utils';
import { DBUser, EMAIL_TYPES, UserRegistrationOptions } from '@/types';
import { sendError } from '@/errors';
import { Joi, email, passwordInsert, registrationOptions } from '@/validation';
import { postgres } from '@/utils/postgres';

export const signUpEmailPasswordSchema = Joi.object({
  email: email.required(),
  password: passwordInsert.required(),
  options: registrationOptions,
}).meta({ className: 'SignUpEmailPasswordSchema' });

export const signUpEmailPasswordHandler: RequestHandler<
  {},
  {},
  {
    email: string;
    password: string;
    options: UserRegistrationOptions & {
      redirectTo: string;
    };
  }
> = async (req, res) => {
  console.log('sign up request email password');

  const { body } = req;
  const {
    email,
    password,
    options: {
      redirectTo,
      locale,
      defaultRole,
      allowedRoles,
      metadata,
      displayName = email,
    },
  } = body;

  // check if email already in use by some other user
  if (await getUserByEmail(email)) {
    return sendError(res, 'email-already-in-use');
  }

  // hash password
  const passwordHash = await hashPassword(password);

  // create ticket
  const ticket = `verifyEmail:${uuidv4()}`;
  const ticketExpiresAt = generateTicketExpiresAt(60 * 60 * 24 * 30); // 30 days

  // insert user
  const { rows } = await postgres.runSqlParsed(
    `INSERT INTO auth.users (
        disabled, 
        display_name, 
        avatar_url,
        email,
        password_hash,
        ticket,
        ticket_expires_at,
        email_verified,
        locale,
        default_role,
        metadata
      ) VALUES (%L, %L, %L, %L, %L, %L, %L, %L, %L, %L, %L) RETURNING id`,
    [
      ENV.AUTH_DISABLE_NEW_USERS,
      displayName,
      getGravatarUrl(email),
      email,
      passwordHash,
      ticket,
      ticketExpiresAt,
      false,
      locale,
      defaultRole,
      metadata,
    ]
  );

  if (!Array.isArray(rows[0])) {
    return;
  }

  const insertedUserId = rows[0][0];

  // insert allowed roles
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

  if (users.length === 0) {
    return null;
  }

  const user = users[0] as DBUser;

  console.log('user:');
  console.log(user);

  // user is now inserted. Continue sending out activation email
  if (
    !ENV.AUTH_DISABLE_NEW_USERS &&
    ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED
  ) {
    const template = 'email-verify';
    const link = createEmailRedirectionLink(
      EMAIL_TYPES.VERIFY,
      ticket,
      redirectTo
    );
    await emailClient.send({
      template,
      message: {
        to: email,
        headers: {
          /** @deprecated */
          'x-ticket': {
            prepared: true,
            value: ticket,
          },
          /** @deprecated */
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
        displayName,
        email,
        ticket,
        redirectTo: encodeURIComponent(redirectTo),
        locale: user.locale,
        serverUrl: ENV.AUTH_SERVER_URL,
        clientUrl: ENV.AUTH_CLIENT_URL,
      },
    });
  }

  // SIGNIN_EMAIL_VERIFIED_REQUIRED = true => Must verify email before sign in
  // SIGNIN_EMAIL_VERIFIED_REQUIRED = true => Don't have to verify email before
  // sign in

  if (!ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED) {
    const signInResponse = await getSignInResponse({
      userId: user.id,
      checkMFA: false,
    });

    // return logged in session because user does not have to verify their email
    // to sign in
    return res.send(signInResponse);
  }

  return res.send({ session: null, mfa: null });
};
