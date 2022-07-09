import { sendError } from '@/errors';
import { getClaims } from '@/utils';
import { Joi, token } from '@/validation';
import { RequestHandler } from 'express';
import { ReasonPhrases } from 'http-status-codes';

export const verifyTokenSchema = Joi.object({
  token,
}).meta({ className: 'VerifyTokenSchema' });

export const verifyTokenHandler: RequestHandler<
  {},
  {},
  { token?: string }
> = async (req, res) => {
  const authorization = req.body.token || req.headers.authorization;

  try {
    await getClaims(authorization);
    return res.json(ReasonPhrases.OK);
  } catch (e) {
    return sendError(res, 'unauthenticated-user');
  }
};
