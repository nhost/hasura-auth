import randomNumber from 'random-number-csprng';
import bcrypt from 'bcryptjs';
import { generateTicketExpiresAt } from './ticket';
import { ENV } from './env';

export const getNewOneTimePasswordData = async () => {
  const pw = await randomNumber(0, 999999);
  const otp = pw.toString().padStart(6, '0');
  const otpHash = await bcrypt.hash(otp, 10);
  const otpHashExpiresAt = generateTicketExpiresAt(
    ENV.AUTH_SMS_VERIFICATION_TIMEOUT
  );

  return {
    otp,
    otpHash,
    otpHashExpiresAt,
  };
};
