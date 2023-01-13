import type { EmailOptions } from 'email-templates';

import { ENV } from './utils/env';
import { EmailLocals, renderTemplate } from './templates';
import { logger } from './logger';

export const sendEmail = async (options: EmailOptions) => {
  // TODO this library takes more than one third of the time required by hasura-auth to load
  const { default: Email } = await import('email-templates');
  const { createTransport } = await import('nodemailer');
  /**
   * SMTP transport.
   */
  const transport = createTransport({
    host: ENV.AUTH_SMTP_HOST,
    port: Number(ENV.AUTH_SMTP_PORT),
    secure: Boolean(ENV.AUTH_SMTP_SECURE),
    auth: {
      pass: ENV.AUTH_SMTP_PASS,
      user: ENV.AUTH_SMTP_USER,
    },
    authMethod: ENV.AUTH_SMTP_AUTH_METHOD,
  });

  const emailClient = new Email<EmailLocals>({
    preview: false,
    transport,
    message: { from: ENV.AUTH_SMTP_SENDER },
    send: true,
    render: renderTemplate,
  });
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
