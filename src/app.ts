import express from 'express';
import helmet from 'helmet';
import { json } from 'body-parser';
import cors from 'cors';

import router from './routes';
import { serverErrors } from './errors';
import { authMiddleware } from './middleware/auth';
import { uncaughtErrorLogger, httpLogger } from './logger';

export const createApp = async () => {
  const app = express();

  if (process.env.NODE_ENV === 'production') {
    app.set('trust proxy', 1);
  }

  // TODO performance: no openapi in production - confirm and document the env var - and adapt Makefile
  if (process.env.AUTH_OPENAPI_ENABLED === 'true') {
    const { addOpenApiRoute } = await import('./openapi');
    addOpenApiRoute(app);
  }

  app.use(httpLogger);

  app.use(helmet(), json(), cors());
  app.use(authMiddleware);

  app.use(router);

  app.use(uncaughtErrorLogger, serverErrors);
  return app;
};
