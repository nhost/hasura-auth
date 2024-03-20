import { RequestHandler } from 'express';
import bcrypt from 'bcryptjs';

import { getSignInResponse, getUserByEmail, ENV } from '@/utils';
import { logger } from '@/logger';
import { sendError } from '@/errors';
import { Joi, email, password } from '@/validation';
import {ClaimValueType} from "@/types";

export const signInEmailPasswordSchema = Joi.object({
  email: email.required(),
  password: password.required(),
  extraClaims: Joi.any().optional(),
}).meta({ className: 'SignInEmailPasswordSchema' });

export const signInEmailPasswordHandler: RequestHandler<
  {},
  {},
  {
    email: string;
    password: string;
    extraClaims?: { [key: string]: ClaimValueType };
  }
> = async (req, res) => {
  const { email, password, extraClaims } = req.body;
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
    extraClaims,
  });

  return res.send(signInTokens);
};
