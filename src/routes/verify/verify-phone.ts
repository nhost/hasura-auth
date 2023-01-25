import { RequestHandler } from 'express';
import bcrypt from 'bcryptjs';

import { ENV, pgClient } from '@/utils';
import { sendError } from '@/errors';
import { isTestingPhoneNumber, isVerifySid } from '@/utils/twilio';
import twilio from 'twilio';
import { OtpSmsRequestBody } from '@/types';
import { ReasonPhrases } from 'http-status-codes';

export const verifyPhoneChangeHandler: RequestHandler<
  {},
  {},
  OtpSmsRequestBody
> = async (req, res) => {
  if (!ENV.AUTH_SMS_PASSWORDLESS_ENABLED) {
    return sendError(res, 'disabled-endpoint');
  }

  const { body } = req;
  const { phoneNumber, otp } = body;
  const { userId } = req.auth as RequestAuth;

  const user = await pgClient.getUserById(userId);

  if (!user) {
    return sendError(res, 'user-not-found');
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  if (phoneNumber !== user.phoneNumber && phoneNumber !== user.newPhoneNumber) {
    return sendError(res, 'invalid-request');
  }

  if (!user.otpHash) {
    return sendError(res, 'invalid-otp');
  }

  async function updateUserPhoneNumber() {
    await pgClient.updateUser({
      id: userId,
      user: {
        otpHash: null,
        phoneNumberVerified: true,
        phoneNumber,
        newPhoneNumber: null,
      },
    });

    return res.json(ReasonPhrases.OK);
  }

  if (isTestingPhoneNumber(phoneNumber)) {
    if (await bcrypt.compare(otp, user.otpHash)) {
      return await updateUserPhoneNumber();
    } else {
      return sendError(res, 'invalid-otp');
    }
  }

  if (!ENV.AUTH_SMS_PROVIDER) {
    throw Error('No sms provider set');
  }

  const messagingServiceSid = ENV.AUTH_SMS_TWILIO_MESSAGING_SERVICE_ID;

  if (isVerifySid(messagingServiceSid)) {
    const twilioClient = twilio(
      ENV.AUTH_SMS_TWILIO_ACCOUNT_SID,
      ENV.AUTH_SMS_TWILIO_AUTH_TOKEN
    );

    try {
      const verificationCheck = await twilioClient.verify
        .services(messagingServiceSid)
        .verificationChecks.create({
          code: otp,
          to: user.phoneNumber ?? '',
        });

      if (!verificationCheck.valid || verificationCheck.status !== 'approved') {
        return sendError(res, 'invalid-otp');
      }
    } catch (error) {
      throw Error('Cannot veirfy otp');
    }
  } else if (!(await bcrypt.compare(otp, user.otpHash))) {
    return sendError(res, 'invalid-otp');
  }

  return await updateUserPhoneNumber();
};
