import { json } from 'body-parser';
import cors from 'cors';
import express from 'express';
import helmet from 'helmet';
import { serverErrors } from './errors';
import { httpLogger, uncaughtErrorLogger } from './logger';
import { authMiddleware } from './middleware/auth';
import { addOpenApiRoute } from './openapi';
import router from './routes';

const app = express();

if (process.env.NODE_ENV === 'production') {
  app.set('trust proxy', 1);
}

addOpenApiRoute(app);

app.use(httpLogger);
app.use(helmet(), json(), cors());
app.use(authMiddleware);
app.use(router);
app.use(uncaughtErrorLogger, serverErrors);

export { app };
