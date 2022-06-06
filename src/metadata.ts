import axios from 'axios';
import { logger } from './logger';
import { ENV } from './utils/env';

/**
 * https://hasura.io/docs/latest/graphql/core/api-reference/schema-metadata-api/relationship.html
 * Here we are using the schema-metadata-api to track the relationships between auth tables
 **/

interface Table {
  name: string;
  schema: string;
}

interface TableArgs {
  source?: string;
  table: Table;
}

type TableConfig = {
  custom_name?: string;
  identifier?: string;
  custom_root_fields?: {
    select?: string;
    select_by_pk?: string;
    select_aggregate?: string;
    insert?: string;
    insert_one?: string;
    update?: string;
    update_by_pk?: string;
    delete?: string;
    delete_by_pk?: string;
  };
  custom_column_names?: {
    [key: string]: string;
  };
};

type TrackTableArgs = TableArgs & {
  configuration?: TableConfig;
};
type UntrackTableArgs = TableArgs & {
  cascade?: boolean;
};

type TableCustomisationArgs = TableArgs & {
  configuration?: TableConfig;
};

type DropRelationshipArgs = TableArgs & {
  relationship: string;
};

type CreateRelationshipArgs = TableArgs & {
  name: string;
  using: {
    foreign_key_constraint_on:
      | {
          table: Table;
          columns: string[];
        }
      | string[];
  };
};

export const runMetadataRequest = async (args: { type: string; args: {} }) => {
  await axios.post(
    ENV.HASURA_GRAPHQL_GRAPHQL_URL.replace('/v1/graphql', '/v1/metadata'),
    args,
    {
      headers: {
        'x-hasura-admin-secret': ENV.HASURA_GRAPHQL_ADMIN_SECRET,
      },
    }
  );
};

// https://hasura.io/docs/latest/graphql/core/api-reference/schema-metadata-api/table-view.html#track-table-v2
export const trackTable = async (args: TrackTableArgs) => {
  try {
    await runMetadataRequest({
      type: 'pg_track_table',
      args,
    });
  } catch (error: any) {
    if (error.response.data.code !== 'already-tracked') {
      logger.error(error);
      throw new Error(`Error tracking table ${args.table.name}`);
    } else {
      logger.debug(`Table ${args.table.name} already tracked`);
    }
  }
};

// https://hasura.io/docs/latest/graphql/core/api-reference/schema-metadata-api/table-view.html#untrack-table
export const untrackTable = async (args: UntrackTableArgs) => {
  try {
    await runMetadataRequest({
      type: 'pg_untrack_table',
      args,
    });
  } catch (error: any) {
    logger.error(error);
    throw new Error(`Error untracking table ${args.table.name}`);
  }
};

export const reloadMetadata = async (
  args: {
    reload_remote_schemas?: boolean;
    reload_sources?: boolean;
    recreate_event_triggers?: boolean;
  } = {}
) => {
  await runMetadataRequest({
    type: 'reload_metadata',
    args,
  });
};
export const setTableCustomization = async (args: TableCustomisationArgs) => {
  logger.info(`Set table customization for ${args.table.name}`);

  try {
    await runMetadataRequest({
      type: 'pg_set_table_customization',
      args,
    });
  } catch (error: any) {
    logger.error(error);
    throw new Error('error setting customization for table ' + args.table.name);
  }
};

export const createObjectRelationship = async (
  args: CreateRelationshipArgs
) => {
  logger.info(`Set object relationship ${args.name} for ${args.table.name}`);
  try {
    await runMetadataRequest({
      type: 'pg_create_object_relationship',
      args,
    });
  } catch (error: any) {
    if (error.response.data.code !== 'already-exists') {
      throw new Error(
        `Error creating object relationship for table ${args.table.name}`
      );
    } else {
      logger.debug(
        `Object relationship ${args.name} on table ${args.table.name} is already created`
      );
    }
  }
};

export const createArrayRelationship = async (args: CreateRelationshipArgs) => {
  logger.info(`Create array relationship ${args.name} for ${args.table.name}`);
  try {
    await runMetadataRequest({
      type: 'pg_create_array_relationship',
      args,
    });
  } catch (error: any) {
    if (error.response.data.code !== 'already-exists') {
      throw new Error(
        `Error creating array relationship for table ${args.table.name}`
      );
    }
    logger.debug(
      `Array relationship ${args.name} on table ${args.table.name} is already created`
    );
  }
};

export const dropRelationship = async (args: DropRelationshipArgs) => {
  logger.info(`Drop relationship ${args.relationship} for ${args.table.name}`);
  try {
    await runMetadataRequest({
      type: 'pg_drop_relationship',
      args,
    });
  } catch (error: any) {
    if (error.response.data.code !== 'already-exists') {
      throw new Error(
        `Error dropping relationship for table ${args.table.name}`
      );
    } else {
      logger.debug(
        `Object relationship ${args.relationship} on table ${args.table.name} is already created`
      );
    }
  }
};

export const applyMetadata = async (): Promise<void> => {
  logger.info('Applying metadata...');

  logger.debug('Reloading metadata...');
  await reloadMetadata();
  logger.debug('Metadata reloaded');

  try {
    await trackTable({
      source: 'default',
      table: {
        schema: 'auth',
        name: 'provider_requests',
      },
    });
    await trackTable({
      source: 'default',
      table: {
        schema: 'auth',
        name: 'refresh_tokens',
      },
    });
    await trackTable({
      source: 'default',
      table: {
        schema: 'auth',
        name: 'roles',
      },
    });
    await trackTable({
      source: 'default',
      table: {
        schema: 'auth',
        name: 'user_providers',
      },
    });
    await trackTable({
      source: 'default',
      table: {
        schema: 'auth',
        name: 'user_roles',
      },
    });
    await trackTable({
      source: 'default',
      table: {
        schema: 'auth',
        name: 'users',
      },
    });
    await trackTable({
      source: 'default',
      table: {
        schema: 'auth',
        name: 'providers',
      },
    });
  } finally {
    logger.debug('Reloading metadata...');
    await reloadMetadata();
    logger.debug('Metadata reloaded');
  }
  logger.info('Metadata applied');
};
