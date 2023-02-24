import { RequestHandler } from 'express';
import bcrypt from 'bcryptjs';

import { getSignInResponse, getUserByEmail, ENV } from '@/utils';
import { logger } from '@/logger';
import { sendError } from '@/errors';
import { Joi, email, password } from '@/validation';

export const signInEmailPasswordSchema = Joi.object({
  email: email.required(),
  password: password.required(),
}).meta({ className: 'SignInEmailPasswordSchema' });

export const signInEmailPasswordHandlerLegacy: RequestHandler<
  {},
  {},
  {
    email: string;
    password: string;
  }
> = async (req, res) => {
  const { email, password } = req.body;
  logger.debug(`Sign in with email: ${email}`);

  const user = await getUserByEmail(email);

  if (!user) {
    return sendError(res, 'invalid-email-password');
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  if (!user.passwordHash) {
    return sendError(res, 'invalid-email-password');
  }

  const isPasswordCorrect = await bcrypt.compare(password, user.passwordHash);

  if (!isPasswordCorrect) {
    return sendError(res, 'invalid-email-password');
  }

  if (ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED && !user.emailVerified) {
    return sendError(res, 'unverified-user');
  }

  const signInTokens = await getSignInResponse({
    userId: user.id,
    checkMFA: true,
  });

  const legacyResponse = {
    jwt_token: signInTokens.session?.accessToken,
    jwt_expires_in: signInTokens.session?.accessTokenExpiresIn,
    user: {
      id: signInTokens.session?.user?.id,
      display_name: signInTokens.session?.user?.displayName,
      email: signInTokens.session?.user?.email,
    },
  };

  return res.send(legacyResponse);
};
