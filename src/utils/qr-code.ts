/**
 * Create QR code.
 * @param secret Required OTP secret.
 */
export const createQR = async (secret: string): Promise<string> => {
  const { default: QRCode } = await import('qrcode');
  return QRCode.toDataURL(secret);
};
