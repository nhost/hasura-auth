import { ENV } from '../env';
import { EMAIL_TYPES, User } from '@/types';
import { sendEmail } from '@/email';
import { createEmailRedirectionLink } from '../redirect';

export const sendEmailIfNotVerified = async ({
  email,
  newEmail,
  user,
  displayName,
  ticket,
  redirectTo,
}: {
  email: string;
  newEmail: string | null;
  user: NonNullable<User>;
  displayName: string;
  ticket?: string | null;
  redirectTo: string;
}) => {
  if (
    !ENV.AUTH_DISABLE_NEW_USERS &&
    ENV.AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED &&
    !user.emailVerified
  ) {
    if (!ticket) {
      throw Error(`No ticket found for the user ${user.id}`);
    }

    const template = 'email-verify';
    const link = createEmailRedirectionLink(
      EMAIL_TYPES.VERIFY,
      ticket,
      redirectTo
    );
    await sendEmail({
      template,
      message: {
        to: email,
        headers: {
          /** @deprecated */
          'x-ticket': {
            prepared: true,
            value: ticket,
          },
          /** @deprecated */
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
      locals: {
        link,
        displayName,
        email,
        newEmail: newEmail,
        ticket,
        redirectTo: encodeURIComponent(redirectTo),
        locale: user.locale,
        serverUrl: ENV.AUTH_SERVER_URL,
        clientUrl: ENV.AUTH_CLIENT_URL,
      },
    });
  }
};
