import { GraphQLClient } from 'graphql-request';
import * as Dom from 'graphql-request/dist/types.dom';
import { getSdk as getSdkGraphqlNaming } from './__generated__/graphql-request-graphql-default';
import {
  getSdk as getSdkDefaultNaming,
  UsersQuery,
  UsersQueryVariables,
  InsertUserMutation,
  InsertUserMutationVariables,
  InsertRefreshTokenMutation,
  InsertRefreshTokenMutationVariables,
  AddUserSecurityKeyMutationVariables,
  AddUserSecurityKeyMutation,
  InsertUserRolesMutation,
  InsertUserRolesMutationVariables,
  UpsertRolesMutation,
  UpsertRolesMutationVariables,
  InsertUserProviderToUserMutationVariables,
  InsertUserProviderToUserMutation,
} from './__generated__/graphql-request-hasura-default';
import { ENV } from './env';
import {
  transformAuthUserProvidersInsertInput,
  transformAuthRefreshTokensInsertInput,
  transformAuthRolesInsertInput,
  transformAuthUserRolesInsertInput,
  transformAuthUserSecurityKeysInsertInput,
  transformUsersInsertInput,
  transformUsersQueryVariables,
} from './gql-type-transforms';

const isGraphqlNamingConvention =
  ENV.HASURA_GRAPHQL_NAMING_CONVENTION === 'graphql-default';

export * from './__generated__/graphql-request-hasura-default';

export const client = new GraphQLClient(ENV.HASURA_GRAPHQL_GRAPHQL_URL, {
  headers: {
    'x-hasura-admin-secret': ENV.HASURA_GRAPHQL_ADMIN_SECRET,
  },
});
const gqlSdkDefaultNaming = getSdkDefaultNaming(client);
const gqlSdkGraphqlNaming = getSdkGraphqlNaming(client);

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type AnyFunction = (...args: any[]) => any;

const choose = <Func extends AnyFunction>(
  defaultNamingFunction: Func,
  graphqlNamingFunction: Func
): ((...args: Parameters<Func>) => ReturnType<Func>) => {
  if (isGraphqlNamingConvention) {
    return graphqlNamingFunction;
  }
  return defaultNamingFunction;
};

export const gqlSdk = {
  user: choose(gqlSdkDefaultNaming.user, gqlSdkGraphqlNaming.user),
  users: (
    variables: UsersQueryVariables,
    requestHeaders?: Dom.RequestInit['headers']
  ): Promise<UsersQuery> => {
    if (isGraphqlNamingConvention) {
      return gqlSdkGraphqlNaming.users(
        transformUsersQueryVariables(variables),
        requestHeaders
      );
    }
    return gqlSdkDefaultNaming.users(variables, requestHeaders);
  },
  insertUser: (
    variables: InsertUserMutationVariables,
    requestHeaders?: Dom.RequestInit['headers']
  ): Promise<InsertUserMutation> => {
    if (isGraphqlNamingConvention) {
      return gqlSdkGraphqlNaming.insertUser(
        {
          user: transformUsersInsertInput(variables.user),
        },
        requestHeaders
      );
    }
    return gqlSdkDefaultNaming.insertUser(variables, requestHeaders);
  },
  updateUser: choose(
    gqlSdkDefaultNaming.updateUser,
    gqlSdkGraphqlNaming.updateUser
  ),
  deleteUser: choose(
    gqlSdkDefaultNaming.deleteUser,
    gqlSdkGraphqlNaming.deleteUser
  ),
  getUsersByPAT: choose(
    gqlSdkDefaultNaming.getUsersByPAT,
    gqlSdkGraphqlNaming.getUsersByPAT
  ),
  insertRefreshToken: (
    variables: InsertRefreshTokenMutationVariables,
    requestHeaders?: Dom.RequestInit['headers']
  ): Promise<InsertRefreshTokenMutation> => {
    if (isGraphqlNamingConvention) {
      return gqlSdkGraphqlNaming.insertRefreshToken(
        {
          refreshToken: transformAuthRefreshTokensInsertInput(
            variables.refreshToken
          ),
        },
        requestHeaders
      );
    }
    return gqlSdkDefaultNaming.insertRefreshToken(variables, requestHeaders);
  },
  getUsersByRefreshToken: choose(
    gqlSdkDefaultNaming.getUsersByRefreshToken,
    gqlSdkGraphqlNaming.getUsersByRefreshToken
  ),
  deleteUserRefreshTokens: choose(
    gqlSdkDefaultNaming.deleteUserRefreshTokens,
    gqlSdkGraphqlNaming.deleteUserRefreshTokens
  ),
  deleteRefreshToken: choose(
    gqlSdkDefaultNaming.deleteRefreshToken,
    gqlSdkGraphqlNaming.deleteRefreshToken
  ),
  deleteExpiredRefreshTokens: choose(
    gqlSdkDefaultNaming.deleteExpiredRefreshTokens,
    gqlSdkGraphqlNaming.deleteExpiredRefreshTokens
  ),
  getUsersByRefreshTokenAndUpdateRefreshTokenExpiresAt: choose(
    gqlSdkDefaultNaming.getUsersByRefreshTokenAndUpdateRefreshTokenExpiresAt,
    gqlSdkGraphqlNaming.getUsersByRefreshTokenAndUpdateRefreshTokenExpiresAt
  ),

  getUserChallenge: choose(
    gqlSdkDefaultNaming.getUserChallenge,
    gqlSdkGraphqlNaming.getUserChallenge
  ),
  updateUserChallenge: choose(
    gqlSdkDefaultNaming.updateUserChallenge,
    gqlSdkGraphqlNaming.updateUserChallenge
  ),

  getUserSecurityKeys: choose(
    gqlSdkDefaultNaming.getUserSecurityKeys,
    gqlSdkGraphqlNaming.getUserSecurityKeys
  ),
  addUserSecurityKey: (
    variables: AddUserSecurityKeyMutationVariables,
    requestHeaders?: Dom.RequestInit['headers']
  ): Promise<AddUserSecurityKeyMutation> => {
    if (isGraphqlNamingConvention) {
      return gqlSdkGraphqlNaming.addUserSecurityKey(
        {
          userSecurityKey: transformAuthUserSecurityKeysInsertInput(
            variables.userSecurityKey
          ),
        },
        requestHeaders
      );
    }
    return gqlSdkDefaultNaming.addUserSecurityKey(variables, requestHeaders);
  },
  updateUserSecurityKey: choose(
    gqlSdkDefaultNaming.updateUserSecurityKey,
    gqlSdkGraphqlNaming.updateUserSecurityKey
  ),

  insertUserRoles: (variables: InsertUserRolesMutationVariables, requestHeaders ?: Dom.RequestInit["headers"]): Promise<InsertUserRolesMutation> => {
    if (isGraphqlNamingConvention) {
      return gqlSdkGraphqlNaming.insertUserRoles(
        {
          userRoles: transformAuthUserRolesInsertInput(variables.userRoles),
        },
        requestHeaders
      );
    }
    return gqlSdkDefaultNaming.insertUserRoles(variables, requestHeaders);
  },
  deleteUserRolesByUserId: choose(
    gqlSdkDefaultNaming.deleteUserRolesByUserId,
    gqlSdkGraphqlNaming.deleteUserRolesByUserId
  ),
  upsertRoles: (variables: UpsertRolesMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<UpsertRolesMutation> => {
    if (isGraphqlNamingConvention) {
      return gqlSdkGraphqlNaming.upsertRoles(
        {
          roles: transformAuthRolesInsertInput(variables.roles),
        },
        requestHeaders
      );
    }
    return gqlSdkDefaultNaming.upsertRoles(variables, requestHeaders);
  },

  authUserProviders: choose(
    gqlSdkDefaultNaming.authUserProviders,
    gqlSdkGraphqlNaming.authUserProviders
  ),
  updateAuthUserprovider: choose(
    gqlSdkDefaultNaming.updateAuthUserprovider,
    gqlSdkGraphqlNaming.updateAuthUserprovider
  ),
  insertUserProviderToUser: (variables: InsertUserProviderToUserMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<InsertUserProviderToUserMutation> => {
    if (isGraphqlNamingConvention) {
      return gqlSdkGraphqlNaming.insertUserProviderToUser(
        {
          userProvider: transformAuthUserProvidersInsertInput(variables.userProvider),
        },
        requestHeaders
      );
    }
    return gqlSdkDefaultNaming.insertUserProviderToUser(variables, requestHeaders);
  },
  providerRequest: choose(
    gqlSdkDefaultNaming.providerRequest,
    gqlSdkGraphqlNaming.providerRequest
  ),
  insertProviderRequest: choose(
    gqlSdkDefaultNaming.insertProviderRequest,
    gqlSdkGraphqlNaming.insertProviderRequest
  ),
  deleteProviderRequest: choose(
    gqlSdkDefaultNaming.deleteProviderRequest,
    gqlSdkGraphqlNaming.deleteProviderRequest
  ),
};
