require('tsconfig-paths/register');
import { Client } from 'pg';

import { ENV } from '../src/utils/env';

export default async (): Promise<void> => {
  const client = new Client({
    connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
  });
  try {
    await client.connect();
    await client.query(
      `INSERT INTO auth.roles (role) VALUES ('editor'), ('super-admin') ON CONFLICT DO NOTHING;;`
    );
  } finally {
    await client.end();
  }
};
