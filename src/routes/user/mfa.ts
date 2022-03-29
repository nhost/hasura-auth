import { sendError } from '@/errors';
import { gqlSdk } from '@/utils/gqlSDK';
import { RequestHandler } from 'express';
import { authenticator } from 'otplib';

export const userMFAHandler: RequestHandler<
  {},
  {},
  {
    code: string;
    activeMfaType: null | 'totp'; // | 'sms';
  }
> = async (req, res) => {
  // check if user is logged in
  if (!req.auth?.userId) {
    return sendError(res, 'unauthenticated-user');
  }

  const { code, activeMfaType } = req.body;

  // TODO joi validation
  if (activeMfaType && !['totp'].includes(activeMfaType)) {
    return sendError(res, 'invalid-mfa-type');
  }

  const { userId } = req.auth;

  const { user } = await gqlSdk.user({
    id: userId,
  });

  if (!user) {
    throw new Error('user could not be fetched');
  }

  if (!activeMfaType) {
    // user wants to deactivate any active MFA type
    if (!user.activeMfaType) {
      return sendError(res, 'mfa-type-not-found');
    }

    if (user.activeMfaType === 'totp') {
      if (!user.totpSecret) {
        return sendError(res, 'no-totp-secret');
      }

      if (!authenticator.check(code, user.totpSecret)) {
        return sendError(res, 'invalid-otp');
      }
    }

    // if (user.activeMfaType === 'sms') {
    // }

    await gqlSdk.updateUser({
      id: userId,
      user: {
        activeMfaType: null,
      },
    });

    return res.send('ok');
  }

  // activate MFA
  if (activeMfaType === 'totp') {
    if (user.activeMfaType === 'totp') {
      return sendError(res, 'totp-already-active');
    }

    if (!user.totpSecret) {
      return sendError(res, 'no-totp-secret');
    }

    if (!authenticator.check(code, user.totpSecret)) {
      return sendError(res, 'invalid-otp');
    }
  }
  // else if (activeMfaType === 'sms') {
  // }

  await gqlSdk.updateUser({
    id: userId,
    user: {
      activeMfaType,
    },
  });

  return res.send('OK');
};
