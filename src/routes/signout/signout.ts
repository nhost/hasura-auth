import { RequestHandler } from 'express';
import { ReasonPhrases } from 'http-status-codes';

import { sendError } from '@/errors';
import { Joi, refreshToken } from '@/validation';
import { postgres } from '@/utils/postgres';

export const signOutSchema = Joi.object({
  refreshToken,
  all: Joi.boolean()
    .default(false)
    .description('Sign out from all connected devices'),
}).meta({ className: 'SignOutSchema' });

export const signOutHandler: RequestHandler<
  {},
  {},
  {
    refreshToken: string;
    all: boolean;
  }
> = async (req, res) => {
  const { refreshToken, all } = req.body;

  if (all) {
    if (!req.auth?.userId) {
      return sendError(res, 'unauthenticated-user', {
        customMessage: 'User must be signed in to sign out from all sessions',
      });
    }

    const { userId } = req.auth;

    await postgres.runSql(
      `DELETE FROM auth.refresh_tokens WHERE user_id = %L`,
      [userId]
    );
  } else {
    // only sign out from the current session
    // delete current refresh token
    await postgres.runSql(
      `DELETE FROM auth.refresh_tokens WHERE refresh_token = %L`,
      [refreshToken]
    );
  }

  return res.send(ReasonPhrases.OK);
};
