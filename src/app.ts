import express from 'express';
import helmet from 'helmet';
import { json } from 'body-parser';
import cors from 'cors';

import router from './routes';
import { serverErrors } from './errors';
import { authMiddleware } from './middleware/auth';
import { uncaughtErrorLogger, httpLogger, logger } from './logger';
import { ENV } from "@/utils";
import { addOpenApiRoute } from "@/openapi";

const app = express();

if (process.env.NODE_ENV === 'production') {
  app.set('trust proxy', 1);
}

if (ENV.AUTH_OPENAPI_SPEC) {
  logger.info('OpenAPI spec is enabled');
  addOpenApiRoute(app);
}
app.use(httpLogger);

app.use(helmet(), json(), cors({
  maxAge: 86400
}));

app.use(authMiddleware);

app.use(router);

app.use(uncaughtErrorLogger, serverErrors);
export { app };
