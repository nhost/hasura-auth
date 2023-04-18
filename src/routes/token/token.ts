import { sendError } from '@/errors';
import {
  createHasuraAccessToken,
  getNewOrUpdateCurrentSession,
  getUser,
  getUserByRefreshToken,
  gqlSdk,
} from '@/utils';
import { Joi, refreshToken } from '@/validation';
import { RequestHandler } from 'express';
import { v4 as uuidv4 } from 'uuid';

export const tokenSchema = Joi.object({
  refreshToken,
}).meta({ className: 'TokenSchema' });

export const createPATSchema = Joi.object({
  expiresAt: Joi.date().required(),
  metadata: Joi.object(),
}).meta({ className: 'CreatePATSchema' });

export const verifyPATSchema = Joi.object({
  personalAccessToken: refreshToken,
}).meta({ className: 'VerifyPATSchema' });

export const tokenHandler: RequestHandler<
  {},
  {},
  { refreshToken: string }
> = async (req, res) => {
  const { refreshToken } = req.body;

  const user = await getUserByRefreshToken(refreshToken);

  if (!user) {
    return sendError(res, 'invalid-refresh-token');
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  // 1 in 10 request will delete expired refresh tokens
  // TODO: CRONJOB in the future.
  if (Math.random() < 0.1) {
    // no await
    gqlSdk.deleteExpiredRefreshTokens();
  }

  const session = await getNewOrUpdateCurrentSession({
    user,
    currentRefreshToken: refreshToken,
  });

  return res.send(session);
};

export const createPATHandler: RequestHandler<
  {},
  {},
  { refreshToken: string; metadata: any; expiresAt: Date }
> = async (req, res) => {
  if (!req.auth) {
    return sendError(res, 'unauthenticated-user');
  }

  const { userId } = req.auth as RequestAuth;

  const user = await getUser({ userId });

  if (!user) {
    return sendError(res, 'invalid-refresh-token');
  }

  const { metadata, expiresAt } = req.body;

  // Expires at must be at least 7 days from now
  if (new Date(expiresAt).getTime() < Date.now() + 7 * 24 * 60 * 60 * 1000) {
    return sendError(res, 'invalid-expiry-date');
  }

  const { id } = user;

  const personalAccessToken = uuidv4();

  await gqlSdk.insertRefreshToken({
    refreshToken: {
      userId: id,
      refreshToken: personalAccessToken,
      expiresAt: new Date(expiresAt),
      metadata,
    },
  });

  return res.send({
    personalAccessToken,
  });
};

export const verifyPATHandler: RequestHandler<
  {},
  {},
  { personalAccessToken: string }
> = async (req, res) => {
  const user = await getUserByRefreshToken(req.body.personalAccessToken);

  if (!user) {
    return sendError(res, 'invalid-refresh-token');
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  const accessToken = await createHasuraAccessToken(user);

  return res.send({ accessToken });
};
