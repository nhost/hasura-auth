import { RequestHandler } from 'express';

import {
  getSignInResponse,
  ENV,
  hashPassword,
  getGravatarUrl,
  createVerifyEmailTicket,
  pgClient,
} from '@/utils';
import { User, UserRegistrationOptionsWithRedirect } from '@/types';
import { sendError } from '@/errors';
import { Joi, email, passwordInsert, registrationOptions } from '@/validation';
import { sendEmailIfNotVerified } from '@/utils/user/email-verification';
import { logger } from '@/logger';

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
    options: UserRegistrationOptionsWithRedirect;
  }
> = async (req, res) => {
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
  // hash password
  const passwordHash = password && (await hashPassword(password));

  // create ticket
  const { ticket, ticketExpiresAt } = createVerifyEmailTicket();

  // insert user
  let user: User;
  try {
    user = await pgClient.insertUser({
      disabled: ENV.AUTH_DISABLE_NEW_USERS,
      displayName,
      avatarUrl: getGravatarUrl(email),
      email,
      passwordHash,
      ticket,
      ticketExpiresAt,
      emailVerified: false,
      locale,
      defaultRole,
      roles: allowedRoles,
      metadata,
    });
  } catch (e) {
    const error = e as Error;
    if (error.message === 'email-already-in-use') {
      return sendError(res, 'email-already-in-use');
    }
    logger.warn('Error while inserting user', { error });
    return sendError(res, 'internal-error');
  }

  await sendEmailIfNotVerified({
    email,
    newEmail: user.newEmail,
    user,
    displayName,
    ticket,
    redirectTo,
  });

  // SIGNIN_EMAIL_VERIFIED_REQUIRED = true => User must verify their email before signing in.
  // SIGNIN_EMAIL_VERIFIED_REQUIRED = false => User don't have to verify their email before signin in.

  if (!ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED) {
    const signInResponse = await getSignInResponse({
      user,
      checkMFA: false,
    });

    // return logged in session because user does not have to verify their email
    // to sign in
    return res.send(signInResponse);
  }

  return res.send({ session: null, mfa: null });
};
