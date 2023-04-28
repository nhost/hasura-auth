require('dotenv').config();

const namingConvention =
  process.env.HASURA_GRAPHQL_NAMING_CONVENTION || 'default';

module.exports = {
  overwrite: true,
  schema: [
    {
      [process.env.HASURA_GRAPHQL_GRAPHQL_URL]: {
        headers: {
          'x-hasura-admin-secret': process.env.HASURA_GRAPHQL_ADMIN_SECRET,
        },
      },
    },
  ],

  documents: `src/gql/naming-convention-${namingConvention}/**/*.graphql`,
  generates: {
    [`src/utils/__generated__/graphql-request-${namingConvention}.ts`]: {
      plugins: [
        {
          add: {
            content: [
              '/* eslint-disable @typescript-eslint/no-explicit-any, @typescript-eslint/no-unused-vars */',
            ],
          },
        },
        'typescript',
        'typescript-operations',
        'typescript-graphql-request',
      ],
      config: {
        skipTypename: true,
      },
    },
  },
};
