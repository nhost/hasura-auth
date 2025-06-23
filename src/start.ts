import 'dotenv/config';
import { logger, LOG_LEVEL } from './logger';
import { ENV } from './utils/env';
import { gqlSdk } from './utils/gql-sdk';

export const start = async () => {
  logger.info(`Log level: ${LOG_LEVEL}`);

  const { waitForHasura } = await import('@/utils');
  const { applyMigrations } = await import('./migrations');
  const { applyMetadata } = await import('./metadata');

  await waitForHasura();
  await applyMigrations();
  await applyMetadata();


  // * Insert missing default allowed roles into the database
  const { insertAuthRoles } = await gqlSdk.upsertRoles({
    roles: [
      ...new Set([
        ...ENV.AUTH_USER_DEFAULT_ALLOWED_ROLES,
        ENV.AUTH_USER_DEFAULT_ROLE,
      ]),
    ].map((role) => ({ role })),
  });
  if (insertAuthRoles?.affected_rows) {
    logger.info(
      `Inserted ${
        insertAuthRoles.affected_rows
      } roles: ${insertAuthRoles.returning.map(({ role }) => role).join(', ')}`
    );
  }
};

start();
