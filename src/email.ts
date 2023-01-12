// TODO this library takes more than one third of the time required by hasura-auth to load
import Email, { EmailOptions } from 'email-templates';
import nodemailer from 'nodemailer';

import { ENV } from './utils/env';
import { EmailLocals, renderTemplate } from './templates';
import { logger } from './logger';

/**
 * SMTP transport.
 */
const transport = nodemailer.createTransport({
  host: ENV.AUTH_SMTP_HOST,
  port: Number(ENV.AUTH_SMTP_PORT),
  secure: Boolean(ENV.AUTH_SMTP_SECURE),
  auth: {
    pass: ENV.AUTH_SMTP_PASS,
    user: ENV.AUTH_SMTP_USER,
  },
  authMethod: ENV.AUTH_SMTP_AUTH_METHOD,
});

/**
 * Reusable email client.
 */
const emailClient = new Email<EmailLocals>({
  preview: false,
  transport,
  message: { from: ENV.AUTH_SMTP_SENDER },
  send: true,
  render: renderTemplate,
});

export const sendEmail = async (options: EmailOptions) => {
  try {
    await emailClient.send(options);
  } catch (err) {
    const error = err as Error;
    logger.warn(
      `SMTP error`,
      Object.entries(error).reduce(
        (acc, [key, value]) => ({
          ...acc,
          [key]: value,
        }),
        {}
      )
    );
    const { template, message } = options;
    logger.warn(`SMTP error context`, {
      template,
      to: message?.to,
    });
    throw err;
  }
};
