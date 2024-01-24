import { sendError } from '@/errors';
import { ENV, getUserByEmail, performWebAuthn, verifyWebAuthn } from '@/utils';
import { RequestHandler } from 'express';

import { SignInResponse } from '@/types';
import { Joi, email } from '@/validation';

import {
  AuthenticationResponseJSON,
  PublicKeyCredentialRequestOptionsJSON,
} from '@simplewebauthn/types';

export type SignInWebAuthnRequestBody = { email: string };
export type SignInWebAuthnResponseBody = PublicKeyCredentialRequestOptionsJSON;

export const signInWebauthnSchema = Joi.object<SignInWebAuthnRequestBody>({
  email: email.required(),
}).meta({ className: 'SignInWebauthnSchema' });

export const signInWebauthnHandler: RequestHandler<
  {},
  SignInWebAuthnResponseBody,
  SignInWebAuthnRequestBody
> = async (req, res) => {
  if (!ENV.AUTH_WEBAUTHN_ENABLED) {
    return sendError(res, 'disabled-endpoint');
  }

  const { body } = req;
  const { email } = body;

  const user = await getUserByEmail(email);

  // ? Do we know to let anyone know if the user doesn't exist?
  if (!user) {
    return sendError(res, 'user-not-found');
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  if (ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED && !user.emailVerified) {
    return sendError(res, 'unverified-user');
  }

  const options = await performWebAuthn(user.id);

  return res.send(options);
};

export type SignInVerifyWebAuthnRequestBody = {
  credential: AuthenticationResponseJSON;
  email: string;
};

export type SignInVerifyWebAuthnResponseBody = SignInResponse;

export const signInVerifyWebauthnSchema =
  Joi.object<SignInVerifyWebAuthnRequestBody>({
    email: email.required(),
    credential: Joi.object().required(),
  }).meta({ className: 'SignInVerifyWebauthnSchema' });

export const signInVerifyWebauthnHandler: RequestHandler<
  {},
  SignInVerifyWebAuthnResponseBody,
  SignInVerifyWebAuthnRequestBody
> = async (req, res) => {
  if (!ENV.AUTH_WEBAUTHN_ENABLED) {
    return sendError(res, 'disabled-endpoint');
  }

  const { credential, email } = req.body;

  const user = await getUserByEmail(email);

  if (!user) {
    return sendError(res, 'user-not-found');
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  if (ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED && !user.emailVerified) {
    return sendError(res, 'unverified-user');
  }

  await verifyWebAuthn(
    user.id,
    credential,
    (code, payload) => sendError(res, code, payload),
    (signInResponse) => res.send(signInResponse)
  );
};
