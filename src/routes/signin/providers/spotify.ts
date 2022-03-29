import { Router } from 'express';
import { Strategy } from 'passport-spotify';
import { PROVIDERS } from '@config/index';
import { initProvider } from './utils';
import { sendError } from '@/errors';

export default (router: Router): void => {
  const options = PROVIDERS.spotify;

  initProvider(
    router,
    'spotify',
    Strategy,
    {
      scope: PROVIDERS.spotify?.scope,
    },
    (req, res, next) => {
      if (!PROVIDERS.spotify) {
        return sendError(res, 'disabled-endpoint');
      } else if (!options?.clientID || !options?.clientSecret) {
        throw new Error(`Missing environment variables for Spotify OAuth`);
      } else {
        return next();
      }
    }
  );
};
