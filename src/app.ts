import { sendError } from '@/errors';
import { ReasonPhrases } from 'http-status-codes';
import { json } from 'body-parser';
import cors from 'cors';
import express from 'express';
import helmet from 'helmet';
import { serverErrors } from './errors';
import { httpLogger, logger, uncaughtErrorLogger } from './logger';
import { authMiddleware } from './middleware/auth';
import { addOpenApiRoute } from './openapi';
import router from './routes';
import { ENV } from './utils/env';

const app = express();

if (process.env.NODE_ENV === 'production') {
  app.set('trust proxy', 1);
}

addOpenApiRoute(app);

app.use(httpLogger);
app.use(helmet(), json(), cors());
app.use(authMiddleware);
app.use(ENV.AUTH_API_PREFIX, router);
app.use(uncaughtErrorLogger, serverErrors);

process.on('unhandledRejection', (reason) => {
  if (reason instanceof Error) {
    logger.error(`Unhandled Rejection`, {
      message: reason.message,
      stack: reason.stack,
    });
    process.exit(1);
  }

  logger.error(`Unhandled Rejection: ${reason}`);
  process.exit(1);
});

process.on('uncaughtException', (err, origin) => {
  logger.error(`Uncaught Exception`, {
    message: err.message,
    stack: err.stack,
    origin,
  });
  process.exit(1);
});

/**
 * GET /healthz
 * @summary Check if the server is up and running
 * @return 200 - Success - application/json
 * @tags General
 */
app.get('/healthz', (_req, res) => res.json(ReasonPhrases.OK));

// all other routes should throw 404 not found
app.use('*', (_req, res) => {
  return sendError(res, 'route-not-found');
});

export { app };
