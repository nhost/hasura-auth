import cors from 'cors';
import { errors } from './errors';
import express from 'express';
import helmet from 'helmet';
import { json } from 'body-parser';
import morgan from 'morgan';
import morganBody from 'morgan-body';
import { limiter } from './limiter';
import router from './routes';
import passport from 'passport';
import { authMiddleware } from './middleware/auth';
import logger from './logger';

const app = express();

if (process.env.NODE_ENV === 'production') {
  app.set('trust proxy', 1);
  app.use(limiter);
}

// app.use(morgan('combined'))
// app.use(
//   morgan(
//     ':remote-addr - :remote-user [:date[clf]] ":method :url HTTP/:http-version" :status :res[content-length] :body'
//   )
// );
// app.use((req, res, next) => {
//   req.logger = logger;
//   return next();
// });

app.use(helmet());
app.use(json());
app.use(cors({ credentials: true, origin: true }));

morganBody(app);

app.use(authMiddleware);

app.use(passport.initialize());

app.use(router);
app.use(errors);

export { app };
