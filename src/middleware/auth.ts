import { RequestHandler } from 'express';
import { getPermissionVariables } from '@/utils';
import { sendError } from '@/errors';
import { checkRecaptchaChallenge } from '@/utils/recaptcha';

export const authMiddleware: RequestHandler = async (req, _, next) => {
  try {
    const permissionVariables = await getPermissionVariables(
      req.headers.authorization
    );
    req.auth = {
      userId: permissionVariables['user-id'],
      defaultRole: permissionVariables['default-role'],
      isAnonymous: permissionVariables['is-anonymous'] === true,
    };
  } catch (e) {
    req.auth = null;
  }
  next();
};

export const authenticationGate: RequestHandler = (req, res, next) => {
  if (!req.auth) {
    return sendError(res, 'unauthenticated-user');
  } else {
    next();
  }
};

export const verifyCaptcha: RequestHandler = async (req, res, next) => {
  const { body } = req;
  const { recaptchaChallenge } = body;

  try {
    const isValidChallenge = await checkRecaptchaChallenge(recaptchaChallenge);

    if (!isValidChallenge) {
      return sendError(res, 'missing-or-invalid-captcha');
    } else {
      next();
    }
  } catch (e: any) {
    sendError(
      res,
      e.message === 'missing-captcha-site-key'
        ? 'missing-captcha-site-key'
        : 'internal-error'
    );
  }
};

export const alwaysAllow: RequestHandler = (_, __, next) => next();
