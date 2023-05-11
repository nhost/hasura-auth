import { RequestHandler } from 'express';
import { getNewOrUpdateCurrentSession, pgClient } from '@/utils';
import { sendError } from '@/errors';
import { Joi, refreshToken } from '@/validation';
import nr from 'newrelic';

export const tokenSchema = Joi.object({
  refreshToken,
}).meta({ className: 'TokenSchema' });

export const tokenHandler: RequestHandler<{},
  {},
  { refreshToken: string }> = async (req, res) => {
  const { refreshToken } = req.body;

  const user = await nr.startSegment('getUserByRefreshToken', true, async () => pgClient.getUserByRefreshToken(refreshToken));

  if (!user) {
    return sendError(res, 'invalid-refresh-token');
  }

  // 1 in 10 request will delete expired refresh tokens
  // TODO: CRONJOB in the future.
  if (Math.random() < 0.001) {
    // no await
    await nr.startSegment('deleteExpiredRefreshTokens', true, pgClient.deleteExpiredRefreshTokens);
  }

  const session = nr.startSegment('getNewOrUpdateCurrentSession', true, async () => await getNewOrUpdateCurrentSession({
    user,
    currentRefreshToken: refreshToken,
  }));

  return res.send(session);
};
