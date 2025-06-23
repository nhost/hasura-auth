import 'dotenv/config';
import { logger, LOG_LEVEL } from './logger';

export const start = async () => {
  logger.info(`Log level: ${LOG_LEVEL}`);

  const { applyMigrations } = await import('./migrations');

  await applyMigrations();
};

start();
