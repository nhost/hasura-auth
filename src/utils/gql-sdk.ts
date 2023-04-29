import { GraphQLClient } from 'graphql-request';
import * as Dom from 'graphql-request/dist/types.dom';
import { getSdk as getSdkGraphqlNaming } from './__generated__/graphql-request-graphql';
import {
  getSdk as getSdkDefaultNaming,
  UsersQuery,
  UsersQueryVariables,
} from './__generated__/graphql-request-default';
import { ENV } from './env';

const isGraphqlNamingConvention =
  ENV.HASURA_GRAPHQL_NAMING_CONVENTION === 'graphql-default';

export * from './__generated__/graphql-request-default';

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
        // Users_Bool_Exp actually matches up with UsersBoolExp, but the the type system
        // doesn't know that because of the named types of relationships like "roles"
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        { where: variables.where as any },
        requestHeaders
      );
    }
    return gqlSdkDefaultNaming.users(variables, requestHeaders);
  },
  insertUser: choose(
    gqlSdkDefaultNaming.insertUser,
    gqlSdkGraphqlNaming.insertUser
  ),
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

  insertRefreshToken: choose(
    gqlSdkDefaultNaming.insertRefreshToken,
    gqlSdkGraphqlNaming.insertRefreshToken
  ),
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
  addUserSecurityKey: choose(
    gqlSdkDefaultNaming.addUserSecurityKey,
    gqlSdkGraphqlNaming.addUserSecurityKey
  ),
  updateUserSecurityKey: choose(
    gqlSdkDefaultNaming.updateUserSecurityKey,
    gqlSdkGraphqlNaming.updateUserSecurityKey
  ),

  insertUserRoles: choose(
    gqlSdkDefaultNaming.insertUserRoles,
    gqlSdkGraphqlNaming.insertUserRoles
  ),
  deleteUserRolesByUserId: choose(
    gqlSdkDefaultNaming.deleteUserRolesByUserId,
    gqlSdkGraphqlNaming.deleteUserRolesByUserId
  ),
  upsertRoles: choose(
    gqlSdkDefaultNaming.upsertRoles,
    gqlSdkGraphqlNaming.upsertRoles
  ),

  authUserProviders: choose(
    gqlSdkDefaultNaming.authUserProviders,
    gqlSdkGraphqlNaming.authUserProviders
  ),
  updateAuthUserprovider: choose(
    gqlSdkDefaultNaming.updateAuthUserprovider,
    gqlSdkGraphqlNaming.updateAuthUserprovider
  ),
  insertUserProviderToUser: choose(
    gqlSdkDefaultNaming.insertUserProviderToUser,
    gqlSdkGraphqlNaming.insertUserProviderToUser
  ),
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
