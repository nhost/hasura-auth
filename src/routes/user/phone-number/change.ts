import { RequestHandler } from 'express';
import { ReasonPhrases } from 'http-status-codes';

import {
  ENV,
  pgClient,
  getUserByPhoneNumber,
  getNewOneTimePasswordData,
} from '@/utils';
import { Joi } from '@/validation';
import { sendError } from '@/errors';
import { isTestingPhoneNumber, isVerifySid } from '@/utils/twilio';
import { renderTemplate } from '@/templates';
import { logger } from '@/logger';
import twilio from 'twilio';

export const userPhoneNumberChangeSchema = Joi.object({
  newPhoneNumber: Joi.string,
}).meta({ className: 'UserPhoneNumberChangeSchema' });

export const userPhoneNumberChange: RequestHandler<
  {},
  {},
  {
    newPhoneNumber: string;
  }
> = async (req, res) => {
  const { newPhoneNumber } = req.body;

  const { userId } = req.auth as RequestAuth;

  // Send an error if the new phone number is already used by another user
  if (await getUserByPhoneNumber({ phoneNumber: newPhoneNumber })) {
    return sendError(res, 'phone-number-already-in-use');
  }

  const { otp, otpHash, otpHashExpiresAt } = await getNewOneTimePasswordData();

  // set newPhoneNumber and otp data for user
  const user = await pgClient.updateUser({
    id: userId,
    user: {
      newPhoneNumber,
      otpMethodLastUsed: 'sms',
      otpHash,
      otpHashExpiresAt,
    },
  });

  if (!user) {
    return sendError(res, 'user-not-found');
  }

  if (user.isAnonymous) {
    return sendError(res, 'forbidden-anonymous');
  }

  if (isTestingPhoneNumber(newPhoneNumber)) {
    const template = 'signin-passwordless-sms';
    const message =
      (await renderTemplate(`${template}/text`, {
        locale: user.locale ?? ENV.AUTH_LOCALE_DEFAULT,
        displayName: user.displayName,
        code: otp,
      })) ?? `Your code is ${otp}`;

    logger.info(`Message to ${newPhoneNumber}: ${message}`);
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
          to: newPhoneNumber,
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
        to: newPhoneNumber,
      });
    }
  } catch (error: any) {
    logger.error('Error sending sms');
    logger.error(error);

    return sendError(res, 'cannot-send-sms');
  }

  return res.json(ReasonPhrases.OK);
};
