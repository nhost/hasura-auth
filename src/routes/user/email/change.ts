import { RequestHandler } from 'express';
import { v4 as uuidv4 } from 'uuid';
import { ReasonPhrases } from 'http-status-codes';

import {
  gqlSdk,
  generateTicketExpiresAt,
  ENV,
  createEmailRedirectionLink,
  getUserByEmail,
} from '@/utils';
import { emailClient } from '@/email';
import { Joi, email, redirectTo } from '@/validation';
import { EMAIL_TYPES } from '@/types';
import { sendError } from '@/errors';

export const userEmailChangeSchema = Joi.object({
  newEmail: email,
  options: Joi.object({
    redirectTo,
  }).default(),
}).meta({ className: 'UserEmailChangeSchema' });

export const userEmailChange: RequestHandler<
  {},
  {},
  {
    newEmail: string;
    options: {
      redirectTo: string;
    };
  }
> = async (req, res) => {
  const {
    newEmail,
    options: { redirectTo },
  } = req.body;

  const { userId } = req.auth as RequestAuth;

  const ticket = `${EMAIL_TYPES.CONFIRM_CHANGE}:${uuidv4()}`;
  const ticketExpiresAt = generateTicketExpiresAt(60 * 60); // 1 hour

  // * Send an error if the new email is already used by another user
  if (await getUserByEmail(newEmail)) {
    return sendError(res, 'email-already-in-use');
  }

  // set newEmail for user
  const updatedUserResponse = await gqlSdk.updateUser({
    id: userId,
    user: {
      ticket,
      ticketExpiresAt,
      newEmail,
    },
  });

  const user = updatedUserResponse.updateUser;

  if (!user) {
    return sendError(res, 'user-not-found');
  }

  if (user.isAnonymous) {
    return sendError(res, 'forbidden-anonymous');
  }

  const template = 'email-confirm-change';
  const link = createEmailRedirectionLink(
    EMAIL_TYPES.CONFIRM_CHANGE,
    ticket,
    redirectTo
  );
  await emailClient.send({
    template,
    locals: {
      link,
      displayName: user.displayName,
      email: user.email,
      newEmail,
      ticket,
      redirectTo: encodeURIComponent(redirectTo),
      locale: user.locale ?? ENV.AUTH_LOCALE_DEFAULT,
      serverUrl: ENV.AUTH_SERVER_URL,
      clientUrl: ENV.AUTH_CLIENT_URL,
    },
    message: {
      to: newEmail,
      headers: {
        'x-ticket': {
          prepared: true,
          value: ticket,
        },
        'x-redirect-to': {
          prepared: true,
          value: redirectTo,
        },
        'x-email-template': {
          prepared: true,
          value: template,
        },
        'x-link': {
          prepared: true,
          value: link,
        },
      },
    },
  });

  return res.send(ReasonPhrases.OK);
};
