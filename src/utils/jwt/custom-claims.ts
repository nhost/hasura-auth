import { jsonToGraphQLQuery, VariableType } from 'json-to-graphql-query';
import { get, set } from 'dot-prop';
import jsonata from 'jsonata';

import { ENV } from '../env';
import { client } from '../gql-sdk';
import { logger } from '@/logger';

/**
 * List of known JSON fields, in a dot-prop notation
 */
const JSON_FIELDS = ['metadata'];

/**
 * Convert array to Postgres array
 * @param arr js array to be converted to Postgres array
 */
function toPgArray(arr: string[]): string {
  const m = arr.map((e) => `"${e}"`).join(',');
  return `{${m}}`;
}

export const escapeValueToPg = (value: unknown): string => {
  // ? Why escaping values? See:
  // * https://hasura.io/docs/latest/graphql/core/auth/authorization/roles-variables.html#format-of-session-variables
  // * https://github.com/hasura/graphql-engine/issues/1902
  if (typeof value === 'string') {
    return value;
  } else if (Array.isArray(value)) {
    return toPgArray(value);
  } else {
    return JSON.stringify(value ?? null);
  }
};

/**
 * Generate a GraphQL request from the values of all the keys of the given jwtFields
 * All values are set together to prepare one single query, to avoid calling the
 * server as many times as there are keys in jwtFields
 * @param jwtFields
 * @returns
 */
const createCustomFieldQuery = (jwtFields: Record<string, string>): string => {
  // * Build the transitional Js object that will be used by json-to-graphql-query
  // ! Advanced JSONata expressions won't work, only basic 'dot' paths e.g. author.books.title
  const fields = Object.values(jwtFields).reduce<Record<string, unknown>>(
    (aggr, path) => {
      set(aggr, path, true);
      return aggr;
    },
    {}
  );

  // * Do not extend JSON field sub-paths into the GraphQL query
  JSON_FIELDS.forEach((field) => {
    if (get(fields, field)) {
      set(fields, field, true);
    }
  });

  // * Prepare the query so it will accept userId as a variable
  const query = {
    query: {
      __variables: {
        userId: 'uuid!',
      },
      user: {
        __args: {
          id: new VariableType('userId'),
        },
        ...fields,
      },
    },
  };
  // * Generate the GraphQL from the above object with json-to-graphql-query
  return jsonToGraphQLQuery(query);
};

/**
 * Generate `x-hasura-` custom claims from a user id
 * @param userId
 * @returns
 */
export const generateCustomClaims = async (userId: string) => {
  // * Don't error if no custom claim is configured, but return an empty object
  if (Object.keys(ENV.AUTH_JWT_CUSTOM_CLAIMS).length === 0) return {};
  // * Generate the GraphQL request from the configuration
  const request = createCustomFieldQuery(ENV.AUTH_JWT_CUSTOM_CLAIMS);
  try {
    // * Fetch user data that is required for all custom expression to be evaluated
    const {
      data: { user },
    } = await client.rawRequest(request, {
      userId,
    });

    // * Aggregate each 'name' key of the object with its value evaluated in JSONata
    return Object.entries(ENV.AUTH_JWT_CUSTOM_CLAIMS).reduce<
      Record<string, unknown>
    >((aggr, [name, path]) => {
      try {
        // * Parse the path into a JSONata AST
        const expression = jsonata(path);
        // * Evaluate the JSONata expression from the fetched user data
        const value = expression.evaluate(user, expression);
        // * Convert value into a PostgreSQL format that can be used by the Hasura permissions system
        // * see {@link escapeValueToPg}
        const jsonataValue = escapeValueToPg(value);
        // * Add the result to the aggregated object, prefixed with `x-hasura-`
        aggr[`x-hasura-${name}`] = jsonataValue;
      } catch {
        // * Don't raise errors if JSONata fails, and log a warning
        logger.warn(`Invalid JSONata expression`, { user, path });
        // * Do not add the claim to the aggregated object
        /** @see {@link https://github.com/nhost/hasura-auth/issues/95} */
        // aggr[`x-hasura-${name}`] = null;
      }
      return aggr;
    }, {});
  } catch {
    // * If the query fail, don't raise an error, but log a warning and return an empty object
    logger.warn(`Invalid custom JWT GraphQL Query`, { request });
    return {};
  }
};
