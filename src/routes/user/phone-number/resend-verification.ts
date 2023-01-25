import { RequestHandler } from 'express';
import { ReasonPhrases } from 'http-status-codes';

import { ENV, pgClient, getNewOneTimePasswordData } from '@/utils';
import { sendError } from '@/errors';
import { isTestingPhoneNumber, isVerifySid } from '@/utils/twilio';
import { renderTemplate } from '@/templates';
import { logger } from '@/logger';
import twilio from 'twilio';

export const userPhoneNumberResendVerificationChange: RequestHandler<
  {},
  {},
  {
    phoneNumber: string;
  }
> = async (req, res) => {
  const { phoneNumber } = req.body;

  const { userId } = req.auth as RequestAuth;

  const userById = await pgClient.getUserById(userId);

  if (!userById) {
    return sendError(res, 'user-not-found');
  }

  if (userById.disabled) {
    return sendError(res, 'disabled-user');
  }

  if (userById.isAnonymous) {
    return sendError(res, 'forbidden-anonymous');
  }

  if (
    phoneNumber === userById.newPhoneNumber ||
    (phoneNumber === userById.phoneNumber && !userById.phoneNumberVerified)
  ) {
    const { otp, otpHash, otpHashExpiresAt } =
      await getNewOneTimePasswordData();

    const user = await pgClient.updateUser({
      id: userId,
      user: {
        otpMethodLastUsed: 'sms',
        otpHash,
        otpHashExpiresAt,
      },
    });

    if (!user) {
      return sendError(res, 'user-not-found');
    }

    if (isTestingPhoneNumber(phoneNumber)) {
      const template = 'signin-passwordless-sms';
      const message =
        (await renderTemplate(`${template}/text`, {
          locale: user.locale ?? ENV.AUTH_LOCALE_DEFAULT,
          displayName: user.displayName,
          code: otp,
        })) ?? `Your code is ${otp}`;

      logger.info(`Message to ${phoneNumber}: ${message}`);
      return res.json(ReasonPhrases.OK);
    }

    if (!ENV.AUTH_SMS_PROVIDER) {
      throw Error('No sms provider set');
    }

    const twilioClient = twilio(
      ENV.AUTH_SMS_TWILIO_ACCOUNT_SID,
      ENV.AUTH_SMS_TWILIO_AUTH_TOKEN
    );

    try {
      const messagingServiceSid = ENV.AUTH_SMS_TWILIO_MESSAGING_SERVICE_ID;

      if (isVerifySid(messagingServiceSid)) {
        await twilioClient.verify
          .services(messagingServiceSid)
          .verifications.create({
            channel: 'sms',
            to: phoneNumber,
          });
      } else {
        const template = 'signin-passwordless-sms';
        const message = await renderTemplate(`${template}/text`, {
          locale: user.locale ?? ENV.AUTH_LOCALE_DEFAULT,
          displayName: user.displayName,
          code: otp,
        });

        await twilioClient.messages.create({
          body: message ?? `Your code is ${otp}`,
          from: ENV.AUTH_SMS_TWILIO_MESSAGING_SERVICE_ID,
          to: phoneNumber,
        });
      }
    } catch (error: any) {
      logger.error('Error sending sms');
      logger.error(error);

      return sendError(res, 'cannot-send-sms');
    }

    return res.json(ReasonPhrases.OK);
  } else {
    return sendError(res, 'phone-number-already-in-use');
  }
};
