import { logger } from '@/logger';
import { Source, overrideMetadata } from '@/utils';

/**
 * This is the metadata source for the `auth` schema. It contains all the
 * necessary tables and relationships for the `auth` schema to work.
 */
export const authMetadataSource: Source = {
  name: 'default',
  kind: 'postgres',
  tables: [
    {
      table: {
        name: 'provider_requests',
        schema: 'auth',
      },
      configuration: {
        column_config: {
          id: {
            custom_name: 'id',
          },
          options: {
            custom_name: 'options',
          },
        },
        custom_column_names: {
          id: 'id',
          options: 'options',
        },
        custom_name: 'authProviderRequests',
        custom_root_fields: {
          delete: 'deleteAuthProviderRequests',
          delete_by_pk: 'deleteAuthProviderRequest',
          insert: 'insertAuthProviderRequests',
          insert_one: 'insertAuthProviderRequest',
          select: 'authProviderRequests',
          select_aggregate: 'authProviderRequestsAggregate',
          select_by_pk: 'authProviderRequest',
          update: 'updateAuthProviderRequests',
          update_by_pk: 'updateAuthProviderRequest',
        },
      },
    },
    {
      table: {
        name: 'providers',
        schema: 'auth',
      },
      configuration: {
        column_config: {
          id: {
            custom_name: 'id',
          },
        },
        custom_column_names: {
          id: 'id',
        },
        custom_name: 'authProviders',
        custom_root_fields: {
          delete: 'deleteAuthProviders',
          delete_by_pk: 'deleteAuthProvider',
          insert: 'insertAuthProviders',
          insert_one: 'insertAuthProvider',
          select: 'authProviders',
          select_aggregate: 'authProvidersAggregate',
          select_by_pk: 'authProvider',
          update: 'updateAuthProviders',
          update_by_pk: 'updateAuthProvider',
        },
      },
      array_relationships: [
        {
          name: 'userProviders',
          using: {
            foreign_key_constraint_on: {
              column: 'provider_id',
              table: {
                name: 'user_providers',
                schema: 'auth',
              },
            },
          },
        },
      ],
    },
    {
      table: { name: 'refresh_token_types', schema: 'auth' },
      is_enum: true,
      configuration: {
        custom_name: 'authRefreshTokenTypes',
        custom_root_fields: {
          select: 'authRefreshTokenTypes',
          select_by_pk: 'authRefreshTokenType',
          select_aggregate: 'authRefreshTokenTypesAggregate',
          insert: 'insertAuthRefreshTokenTypes',
          insert_one: 'insertAuthRefreshTokenType',
          update: 'updateAuthRefreshTokenTypes',
          update_by_pk: 'updateAuthRefreshTokenType',
          delete: 'deleteAuthRefreshTokenTypes',
          delete_by_pk: 'deleteAuthRefreshTokenType',
        },
      },
      array_relationships: [
        {
          name: 'refreshTokens',
          using: {
            foreign_key_constraint_on: {
              column: 'type',
              table: {
                name: 'refresh_tokens',
                schema: 'auth',
              },
            },
          },
        },
      ],
    },
    {
      table: {
        name: 'refresh_tokens',
        schema: 'auth',
      },
      configuration: {
        column_config: {
          created_at: {
            custom_name: 'createdAt',
          },
          expires_at: {
            custom_name: 'expiresAt',
          },
          refresh_token_hash: {
            custom_name: 'refreshTokenHash',
          },
          user_id: {
            custom_name: 'userId',
          },
        },
        custom_column_names: {
          created_at: 'createdAt',
          expires_at: 'expiresAt',
          refresh_token_hash: 'refreshTokenHash',
          user_id: 'userId',
        },
        custom_name: 'authRefreshTokens',
        custom_root_fields: {
          delete: 'deleteAuthRefreshTokens',
          delete_by_pk: 'deleteAuthRefreshToken',
          insert: 'insertAuthRefreshTokens',
          insert_one: 'insertAuthRefreshToken',
          select: 'authRefreshTokens',
          select_aggregate: 'authRefreshTokensAggregate',
          select_by_pk: 'authRefreshToken',
          update: 'updateAuthRefreshTokens',
          update_by_pk: 'updateAuthRefreshToken',
        },
      },
      object_relationships: [
        {
          name: 'refreshTokenType',
          using: {
            foreign_key_constraint_on: 'type',
          },
        },
        {
          name: 'user',
          using: {
            foreign_key_constraint_on: 'user_id',
          },
        },
      ],
      select_permissions: [
        {
          role: 'user',
          permission: {
            columns: [
              'id',
              'created_at',
              'expires_at',
              'metadata',
              'type',
              'user_id',
            ],
            filter: {
              user_id: {
                _eq: 'X-Hasura-User-Id',
              },
            },
          },
        },
      ],
      delete_permissions: [
        {
          role: 'user',
          permission: {
            filter: {
              _and: [
                {
                  user_id: {
                    _eq: 'X-Hasura-User-Id',
                  },
                },
                {
                  type: {
                    _eq: 'pat',
                  },
                },
              ],
            },
          },
        },
      ],
    },
    {
      table: {
        name: 'roles',
        schema: 'auth',
      },
      configuration: {
        column_config: {
          role: {
            custom_name: 'role',
          },
        },
        custom_column_names: {
          role: 'role',
        },
        custom_name: 'authRoles',
        custom_root_fields: {
          delete: 'deleteAuthRoles',
          delete_by_pk: 'deleteAuthRole',
          insert: 'insertAuthRoles',
          insert_one: 'insertAuthRole',
          select: 'authRoles',
          select_aggregate: 'authRolesAggregate',
          select_by_pk: 'authRole',
          update: 'updateAuthRoles',
          update_by_pk: 'updateAuthRole',
        },
      },
      array_relationships: [
        {
          name: 'userRoles',
          using: {
            foreign_key_constraint_on: {
              column: 'role',
              table: {
                name: 'user_roles',
                schema: 'auth',
              },
            },
          },
        },
        {
          name: 'usersByDefaultRole',
          using: {
            foreign_key_constraint_on: {
              column: 'default_role',
              table: {
                name: 'users',
                schema: 'auth',
              },
            },
          },
        },
      ],
    },
    {
      table: {
        name: 'user_providers',
        schema: 'auth',
      },
      configuration: {
        column_config: {
          access_token: {
            custom_name: 'accessToken',
          },
          created_at: {
            custom_name: 'createdAt',
          },
          id: {
            custom_name: 'id',
          },
          provider_id: {
            custom_name: 'providerId',
          },
          provider_user_id: {
            custom_name: 'providerUserId',
          },
          updated_at: {
            custom_name: 'updatedAt',
          },
          user_id: {
            custom_name: 'userId',
          },
        },
        custom_column_names: {
          access_token: 'accessToken',
          created_at: 'createdAt',
          id: 'id',
          provider_id: 'providerId',
          provider_user_id: 'providerUserId',
          updated_at: 'updatedAt',
          user_id: 'userId',
          refresh_token: 'refreshToken',
        },
        custom_name: 'authUserProviders',
        custom_root_fields: {
          delete: 'deleteAuthUserProviders',
          delete_by_pk: 'deleteAuthUserProvider',
          insert: 'insertAuthUserProviders',
          insert_one: 'insertAuthUserProvider',
          select: 'authUserProviders',
          select_aggregate: 'authUserProvidersAggregate',
          select_by_pk: 'authUserProvider',
          update: 'updateAuthUserProviders',
          update_by_pk: 'updateAuthUserProvider',
        },
      },
      object_relationships: [
        {
          name: 'provider',
          using: {
            foreign_key_constraint_on: 'provider_id',
          },
        },
        {
          name: 'user',
          using: {
            foreign_key_constraint_on: 'user_id',
          },
        },
      ],
    },
    {
      table: {
        name: 'user_roles',
        schema: 'auth',
      },
      configuration: {
        column_config: {
          created_at: {
            custom_name: 'createdAt',
          },
          id: {
            custom_name: 'id',
          },
          role: {
            custom_name: 'role',
          },
          user_id: {
            custom_name: 'userId',
          },
        },
        custom_column_names: {
          created_at: 'createdAt',
          id: 'id',
          role: 'role',
          user_id: 'userId',
        },
        custom_name: 'authUserRoles',
        custom_root_fields: {
          delete: 'deleteAuthUserRoles',
          delete_by_pk: 'deleteAuthUserRole',
          insert: 'insertAuthUserRoles',
          insert_one: 'insertAuthUserRole',
          select: 'authUserRoles',
          select_aggregate: 'authUserRolesAggregate',
          select_by_pk: 'authUserRole',
          update: 'updateAuthUserRoles',
          update_by_pk: 'updateAuthUserRole',
        },
      },
      object_relationships: [
        {
          name: 'roleByRole',
          using: {
            foreign_key_constraint_on: 'role',
          },
        },
        {
          name: 'user',
          using: {
            foreign_key_constraint_on: 'user_id',
          },
        },
      ],
    },
    {
      table: {
        name: 'user_security_keys',
        schema: 'auth',
      },
      configuration: {
        column_config: {
          credential_id: {
            custom_name: 'credentialId',
          },
          credential_public_key: {
            custom_name: 'credentialPublicKey',
          },
          id: {
            custom_name: 'id',
          },
          user_id: {
            custom_name: 'userId',
          },
        },
        custom_column_names: {
          credential_id: 'credentialId',
          credential_public_key: 'credentialPublicKey',
          id: 'id',
          user_id: 'userId',
        },
        custom_name: 'authUserSecurityKeys',
        custom_root_fields: {
          delete: 'deleteAuthUserSecurityKeys',
          delete_by_pk: 'deleteAuthUserSecurityKey',
          insert: 'insertAuthUserSecurityKeys',
          insert_one: 'insertAuthUserSecurityKey',
          select: 'authUserSecurityKeys',
          select_aggregate: 'authUserSecurityKeysAggregate',
          select_by_pk: 'authUserSecurityKey',
          update: 'updateAuthUserSecurityKeys',
          update_by_pk: 'updateAuthUserSecurityKey',
        },
      },
      object_relationships: [
        {
          name: 'user',
          using: {
            foreign_key_constraint_on: 'user_id',
          },
        },
      ],
    },
    {
      table: {
        name: 'users',
        schema: 'auth',
      },
      configuration: {
        column_config: {
          active_mfa_type: {
            custom_name: 'activeMfaType',
          },
          avatar_url: {
            custom_name: 'avatarUrl',
          },
          created_at: {
            custom_name: 'createdAt',
          },
          default_role: {
            custom_name: 'defaultRole',
          },
          disabled: {
            custom_name: 'disabled',
          },
          display_name: {
            custom_name: 'displayName',
          },
          email: {
            custom_name: 'email',
          },
          email_verified: {
            custom_name: 'emailVerified',
          },
          id: {
            custom_name: 'id',
          },
          is_anonymous: {
            custom_name: 'isAnonymous',
          },
          last_seen: {
            custom_name: 'lastSeen',
          },
          locale: {
            custom_name: 'locale',
          },
          new_email: {
            custom_name: 'newEmail',
          },
          otp_hash: {
            custom_name: 'otpHash',
          },
          otp_hash_expires_at: {
            custom_name: 'otpHashExpiresAt',
          },
          otp_method_last_used: {
            custom_name: 'otpMethodLastUsed',
          },
          password_hash: {
            custom_name: 'passwordHash',
          },
          phone_number: {
            custom_name: 'phoneNumber',
          },
          phone_number_verified: {
            custom_name: 'phoneNumberVerified',
          },
          ticket: {
            custom_name: 'ticket',
          },
          ticket_expires_at: {
            custom_name: 'ticketExpiresAt',
          },
          totp_secret: {
            custom_name: 'totpSecret',
          },
          updated_at: {
            custom_name: 'updatedAt',
          },
          webauthn_current_challenge: {
            custom_name: 'currentChallenge',
          },
        },
        custom_column_names: {
          active_mfa_type: 'activeMfaType',
          avatar_url: 'avatarUrl',
          created_at: 'createdAt',
          default_role: 'defaultRole',
          disabled: 'disabled',
          display_name: 'displayName',
          email: 'email',
          email_verified: 'emailVerified',
          id: 'id',
          is_anonymous: 'isAnonymous',
          last_seen: 'lastSeen',
          locale: 'locale',
          new_email: 'newEmail',
          otp_hash: 'otpHash',
          otp_hash_expires_at: 'otpHashExpiresAt',
          otp_method_last_used: 'otpMethodLastUsed',
          password_hash: 'passwordHash',
          phone_number: 'phoneNumber',
          phone_number_verified: 'phoneNumberVerified',
          ticket: 'ticket',
          ticket_expires_at: 'ticketExpiresAt',
          totp_secret: 'totpSecret',
          updated_at: 'updatedAt',
          webauthn_current_challenge: 'currentChallenge',
        },
        custom_name: 'users',
        custom_root_fields: {
          delete: 'deleteUsers',
          delete_by_pk: 'deleteUser',
          insert: 'insertUsers',
          insert_one: 'insertUser',
          select: 'users',
          select_aggregate: 'usersAggregate',
          select_by_pk: 'user',
          update: 'updateUsers',
          update_by_pk: 'updateUser',
        },
      },
      object_relationships: [
        {
          name: 'defaultRoleByRole',
          using: {
            foreign_key_constraint_on: 'default_role',
          },
        },
      ],
      array_relationships: [
        {
          name: 'refreshTokens',
          using: {
            foreign_key_constraint_on: {
              column: 'user_id',
              table: {
                name: 'refresh_tokens',
                schema: 'auth',
              },
            },
          },
        },
        {
          name: 'roles',
          using: {
            foreign_key_constraint_on: {
              column: 'user_id',
              table: {
                name: 'user_roles',
                schema: 'auth',
              },
            },
          },
        },
        {
          name: 'securityKeys',
          using: {
            foreign_key_constraint_on: {
              column: 'user_id',
              table: {
                name: 'user_security_keys',
                schema: 'auth',
              },
            },
          },
        },
        {
          name: 'userProviders',
          using: {
            foreign_key_constraint_on: {
              column: 'user_id',
              table: {
                name: 'user_providers',
                schema: 'auth',
              },
            },
          },
        },
      ],
    },
  ],
  configuration: {
    connection_info: {
      database_url: {
        from_env: 'HASURA_GRAPHQL_DATABASE_URL',
      },
      isolation_level: 'read-committed',
      pool_settings: {
        connection_lifetime: 600,
        idle_timeout: 180,
        max_connections: 50,
        retries: 1,
      },
      use_prepared_statements: true,
    },
  },
};

export const applyMetadata = async (): Promise<void> => {
  logger.info('Applying metadata...');
  try {
    await overrideMetadata(authMetadataSource);
    logger.info('Metadata applied');
  } catch (e) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const error = e as any;
    const message = error.response?.data || error.message;
    logger.warn('Impossible to apply metadata', message);
  }
};
