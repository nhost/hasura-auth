import { RequestHandler } from 'express';
import { getNewOrUpdateCurrentSession, pgClient } from '@/utils';
import { sendError } from '@/errors';
import { Joi, refreshToken } from '@/validation';
import NodeCache from 'node-cache';

const cache = new NodeCache({ stdTTL: 60 * 60 * 24 });

const tokenCache = new NodeCache({ stdTTL: 850 });

export const tokenSchema = Joi.object({
  refreshToken,
}).meta({ className: 'TokenSchema' });

export const tokenHandler: RequestHandler<{},
  {},
  { refreshToken: string }> = async (req, res) => {
  const { refreshToken } = req.body;

  if (cache.get(refreshToken) === true) {
    return sendError(res, 'invalid-refresh-token');
  }

  const user = await pgClient.getUserByRefreshToken(refreshToken);

  if (!user) {
    cache.set(refreshToken, true);
    return sendError(res, 'invalid-refresh-token');
  }

  if (tokenCache.get(user.id)) {
    return res.send(tokenCache.get(user.id));
  }

  // 1 in 10 request will delete expired refresh tokens
  // TODO: CRONJOB in the future.
  if (Math.random() < 0.001) {
    // no await
    await pgClient.deleteExpiredRefreshTokens();
  }

  const session = await getNewOrUpdateCurrentSession({
    user,
    currentRefreshToken: refreshToken,
  });

  tokenCache.set(user.id, session);

  return res.send(session);
};
