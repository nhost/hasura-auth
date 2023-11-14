import axios from 'axios';
import { ENV } from '@/utils/env';

export const RECAPTCHA_BASE_URL =
  'https://recaptchaenterprise.googleapis.com/v1/projects';

export const checkRecaptchaChallenge = async (recaptchaChallenge: string) => {
  if (!ENV.AUTH_RECAPTCHA_GOOGLE_API_KEY) {
    throw new Error('missing-captcha-site-key');
  }

  const payload = {
    event: {
      token: recaptchaChallenge,
      siteKey: ENV.AUTH_RECAPTCHA_SITE_KEY,
    },
  };

  const { data } = await axios.post(
    `${RECAPTCHA_BASE_URL}/${ENV.AUTH_RECAPTCHA_PROJECT_ID}/assessments?key=${ENV.AUTH_RECAPTCHA_GOOGLE_API_KEY}`,
    payload
  );

  const isValid = data.tokenProperties.valid;

  if (!isValid) {
    console.error(
      `The CreateAssessment call failed because the token was: ${data.tokenProperties.invalidReason}`
    );
  }

  return isValid;
};
