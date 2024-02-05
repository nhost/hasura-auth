import { RequestHandler } from 'express';
import twilio from 'twilio';
import { ReasonPhrases } from 'http-status-codes';

import { UserRegistrationOptions } from '@/types';
import {
  gqlSdk,
  getNewOneTimePasswordData,
  getUserByPhoneNumber,
  insertUser,
  ENV,
} from '@/utils';
import { sendError } from '@/errors';
import { Joi, phoneNumber, registrationOptions } from '@/validation';
import { isTestingPhoneNumber, isVerifySid } from '@/utils/twilio';
import { logger } from '@/logger';
import { renderTemplate } from '@/templates';
import Dysmsapi20170525, * as $Dysmsapi20170525 from '@alicloud/dysmsapi20170525';
import * as OpenApi from '@alicloud/openapi-client';

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

  // check if email already exist
  let user = await getUserByPhoneNumber({ phoneNumber });
  const userExists = !!user;

  // if no user exists, create the user
  if (!userExists) {
    user = await insertUser({
      disabled: ENV.AUTH_DISABLE_NEW_USERS,
      displayName,
      avatarUrl: '',
      phoneNumber,
      locale,
      defaultRole,
      roles: {
        // restructure user roles to be inserted in GraphQL mutation
        data: allowedRoles.map((role: string) => ({ role })),
      },
      metadata,
    });
  }

  if (user.disabled) {
    return sendError(res, 'disabled-user');
  }

  // set otp for user that will be sent in the email
  const { otp, otpHash, otpHashExpiresAt } = await getNewOneTimePasswordData();

  await gqlSdk.updateUser({
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


  try {
    switch (ENV.AUTH_SMS_PROVIDER) {
      case "twilio":
        await sendSmsWithTwilio(phoneNumber, user, otp);
        break;
      case "alicloud":
        await sendSmsWithAlicloud(phoneNumber, user, otp);
        break;
    }

  } catch (error: any) {
    logger.error('Error sending sms');
    logger.error(error);

    // delete user that was inserted because we were not able to send the SMS
    if (!userExists) {
      await gqlSdk.deleteUser({
        userId: user.id,
      });
    }
    return sendError(res, 'cannot-send-sms');
  }

  return res.json(ReasonPhrases.OK);
};



async function sendSmsWithTwilio(phoneNumber: string, user: any, otp: string) {
  const twilioClient = twilio(
    ENV.AUTH_SMS_TWILIO_ACCOUNT_SID,
    ENV.AUTH_SMS_TWILIO_AUTH_TOKEN
  );
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
}





async function sendSmsWithAlicloud(phoneNumber: string, user: any, otp: string) {
  let config = new OpenApi.Config({
    accessKeyId: ENV.AUTH_SMS_ALICLOUD_ACCESS_KEY_ID,
    accessKeySecret: ENV.AUTH_SMS_ALICLOUD_ACCESS_KEY_SECRET,
    endpoint: ENV.AUTH_SMS_ALICLOUD_ENDPOINT
  });

  let locale = user.locale ?? ENV.AUTH_LOCALE_DEFAULT;
  let template = ENV.AUTH_SMS_ALICLOUD_TEMPLATE_LOCALE[locale] ?? {
    templateCode: ENV.AUTH_SMS_ALICLOUD_TEMPLATE_CODE_DEFAULT,
    signName: ENV.AUTH_SMS_ALICLOUD_SIGN_NAME_DEFAULT
  };

  const client = new Dysmsapi20170525(config);
  const request = new $Dysmsapi20170525.SendSmsRequest({
    phoneNumbers: phoneNumber,
    templateParam: JSON.stringify({ "code": otp }),
    ...template
  });

  const response = await client.sendSms(request);

  if (response.body.code !== "OK") {
    logger.error(response)
    throw Error(response.body.message)
  }

}