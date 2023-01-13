import { RequestHandler } from 'express';

import { getSignInResponse, getUserByTicket } from '@/utils';
import { sendError } from '@/errors';

export const signInMfaSmspHandler: RequestHandler<
  {},
  {},
  {
    ticket: string;
    otp: string;
  }
> = async (req, res) => {
  const { ticket, otp } = req.body;

  const user = await getUserByTicket(ticket);

  if (!user || !user.otpHash) {
    return sendError(res, 'invalid-otp');
  }
  const { compare } = await import('bcrypt');
  if (!(await compare(otp, user.otpHash))) {
    return sendError(res, 'invalid-otp');
  }

  const signInResponse = await getSignInResponse({
    user,
    checkMFA: false,
  });

  return res.send(signInResponse);
};
