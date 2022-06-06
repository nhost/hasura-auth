import { RequestHandler } from 'express';
import { authenticator } from 'otplib';

import { createQR, ENV } from '@/utils';
import { sendError } from '@/errors';
import { postgres } from '@/utils/postgres';

export const mfatotpGenerateHandler: RequestHandler<
  {},
  { imageUrl: string; totpSecret: string },
  {}
> = async (req, res) => {
  if (!ENV.AUTH_MFA_ENABLED) {
    return sendError(res, 'disabled-endpoint');
  }
  const { userId } = req.auth as RequestAuth;

  const totpSecret = authenticator.generateSecret();
  const otpAuth = authenticator.keyuri(
    userId,
    ENV.AUTH_MFA_TOTP_ISSUER,
    totpSecret
  );

  await postgres.runSql('UPDATE users SET totp_secret = %L WHERE id = %L', [
    totpSecret,
    userId,
  ]);

  const imageUrl = await createQR(otpAuth);

  return res.send({ imageUrl, totpSecret });
};
