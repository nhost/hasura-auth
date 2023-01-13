import { RequestHandler } from 'express';

import { getSignInResponse, ENV, pgClient } from '@/utils';
import { logger } from '@/logger';
import { sendError } from '@/errors';
import { Joi, email, password } from '@/validation';

export const signInEmailPasswordSchema = Joi.object({
  email: email.required(),
  password: password.required(),
}).meta({ className: 'SignInEmailPasswordSchema' });

export const signInEmailPasswordHandler: RequestHandler<
  {},
  {},
  {
    email: string;
    password: string;
  }
> = async (req, res) => {
  const { email, password } = req.body;
  logger.debug(`Sign in with email: ${email}`);

  const user = await pgClient.getUserByEmail(email);

  if (!user) {
    return sendError(res, 'invalid-email-password');
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  if (!user.passwordHash) {
    return sendError(res, 'invalid-email-password');
  }

  const { compare } = await import('bcrypt');
  const isPasswordCorrect = await compare(password, user.passwordHash);

  if (!isPasswordCorrect) {
    return sendError(res, 'invalid-email-password');
  }

  if (ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED && !user.emailVerified) {
    return sendError(res, 'unverified-user');
  }

  const signInTokens = await getSignInResponse({
    user,
    checkMFA: true,
  });

  return res.send(signInTokens);
};
