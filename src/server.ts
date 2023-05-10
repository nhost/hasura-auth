import { pgClient } from '@/utils';
import { ENV } from '@/utils/env';
import { logger, LOG_LEVEL } from './logger';

export const start = async () => {
  logger.info(`Log level: ${LOG_LEVEL}`);

  /**
   * Async imports allow to dynamically import the required modules,
   * and therefore only import modules that are required, when they are required.
   * It decreases the loading time on startup.
   * This distinction can make a difference if using hasura-auth to only apply migrations/metadata,
   * or to skip migrations/metadata on startup.
   */
  // if (ENV.AUTH_SKIP_INIT) {
  //   logger.info(`Skipping migrations and metadata`);
  // } else {
  const { waitForHasura } = await import('@/utils');
  const { applyMigrations } = await import('./migrations');
  const { applyMetadata } = await import('./metadata');

  // wait for hasura to be ready
  await waitForHasura();
  // apply migrations and metadata
  if (ENV.AUTH_APPLY_MIGRATIONS) {
    logger.info(`Applying migrations and metadata`);
    await applyMigrations();
    await applyMetadata();
  } else {
    logger.info(`Skipping migrations and metadata`);
  }
  // }

  // * Insert missing default allowed roles into the database
  const newRoles = await pgClient.upsertRoles([
    ...ENV.AUTH_USER_DEFAULT_ALLOWED_ROLES,
    ENV.AUTH_USER_DEFAULT_ROLE,
  ]);
  if (newRoles.length) {
    logger.info(`Inserted ${newRoles.length} roles: ${newRoles.join(', ')}`);
  }

  // if (ENV.AUTH_SKIP_SERVE) {
  //   logger.info(
  //     `AUTH_SKIP_SERVE has been set to true, therefore Hasura-Auth won't start`
  //   );
  // } else {
  await import('./env-vars-check');
  const { app } = await import('./app');

  app.listen(ENV.AUTH_PORT, () => {
    logger.info(`Running on port ${ENV.AUTH_PORT}`);
  });
  // }
};
