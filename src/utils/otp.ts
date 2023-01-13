import { generateTicketExpiresAt } from './ticket';

// * Performance: the library `random-number-csprng` was slow to import and slow to run, with no clear benefit over `Math.random()`.
const randomNumber = (min: number, max: number) =>
  Math.floor(Math.random() * (max - min + 1) + min);

export const getNewOneTimePasswordData = async () => {
  const pw = await randomNumber(0, 999999);
  const otp = pw.toString().padStart(6, '0');
  const { hash } = await import('bcrypt');
  const otpHash = await hash(otp, 10);
  const otpHashExpiresAt = generateTicketExpiresAt(5 * 60);

  return {
    otp,
    otpHash,
    otpHashExpiresAt,
  };
};
