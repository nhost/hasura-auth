import { RequestHandler } from 'express';
import twilio from 'twilio';
import { ReasonPhrases } from 'http-status-codes';

import { UserRegistrationOptions } from '@/types';
import {
  getNewOneTimePasswordData,
  getUserByPhoneNumber,
  insertUser,
  ENV,
  pgClient,
} from '@/utils';
import { sendError } from '@/errors';
import { Joi, phoneNumber, registrationOptions } from '@/validation';
import { isTestingPhoneNumber, isVerifySid } from '@/utils/twilio';
import { logger } from '@/logger';
import { renderTemplate } from '@/templates';

export type PasswordLessSmsRequestBody = {
  phoneNumber: string;
  options: UserRegistrationOptions;
};

export const signInPasswordlessSmsSchema =
  Joi.object<PasswordLessSmsRequestBody>({
    phoneNumber,
    options: registrationOptions,
  }).meta({ className: 'SignInPasswordlessSmsSchema' });

export const signInPasswordlessSmsHandler: RequestHandler<
  {},
  {},
  PasswordLessSmsRequestBody
> = async (req, res) => {
  if (!ENV.AUTH_SMS_PASSWORDLESS_ENABLED) {
    return sendError(res, 'disabled-endpoint');
  }

  const {
    phoneNumber,
    options: { defaultRole, allowedRoles, displayName, locale, metadata },
  } = req.body;

  // check if user with given phone number already exist
  let user = await getUserByPhoneNumber({ phoneNumber });
  const userExists = !!user;

  // if no user exists, create the user
  if (!user) {
    user = await insertUser({
      disabled: ENV.AUTH_DISABLE_NEW_USERS,
      displayName,
      avatarUrl: '',
      phoneNumber,
      locale,
      defaultRole,
      roles: allowedRoles,
      metadata,
    });
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  const { otp, otpHash, otpHashExpiresAt } = await getNewOneTimePasswordData();

  await pgClient.updateUser({
    id: user.id,
    user: {
      otpMethodLastUsed: 'sms',
      otpHash,
      otpHashExpiresAt,
    },
  });

  if (isTestingPhoneNumber(user.phoneNumber)) {
    const template = 'signin-passwordless-sms';
    const message =
      (await renderTemplate(`${template}/text`, {
        locale: user.locale ?? ENV.AUTH_LOCALE_DEFAULT,
        displayName: user.displayName,
        code: otp,
      })) ?? `Your code is ${otp}`;

    logger.info(`Message to ${user.phoneNumber}: ${message}`);
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

    // delete user that was inserted because we were not able to send the SMS
    if (!userExists) {
      await pgClient.deleteUser(user.id);
    }
    return sendError(res, 'cannot-send-sms');
  }

  return res.json(ReasonPhrases.OK);
};
