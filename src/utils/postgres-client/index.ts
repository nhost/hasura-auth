import { v4 as uuidv4 } from 'uuid';
import { User } from '@/types';
import { SessionData } from 'express-session';
import { Pool } from 'pg';
export * from './types';

import { ENV } from '../env';
import { hashRefreshToken, newRefreshExpiry } from '../refresh-token';
import { SqlUser, UserSecurityKey } from './types';
import {
  cameliseUser,
  createUserQueryByColumn,
  createUserQueryWhere,
  getUserById,
  insertUserRoles,
  snakeiseUser,
} from './utils';
import nr from "newrelic";

const pool = new Pool({
  max: ENV.AUTH_PG_POOL_MAX_CONNECTIONS,
  connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
  idleTimeoutMillis: 30000,
});

export const pgClient = {
  async query<T>(text: string, params?: any[]) {
    const client = await nr.startSegment('pg-pool:connect', true, async () => await pool.connect());
    const result = nr.startSegment('pg-pool:query', true, async () => await client.query<T>(text, params));
    await nr.startSegment('pg-pool:release', true, async () => client.release());
    return result;
  },

  insertProviderRequest: async (id: string, options: SessionData) => {
    await pgClient.query(
      `INSERT INTO "auth"."provider_requests" (id, options) VALUES($1, $2) 
        ON CONFLICT(id) DO UPDATE SET options = EXCLUDED.options;`,
      [id, options]
    );
  },

  deleteProviderRequest: async (id: string) => {
    await pgClient.query(
      `DELETE FROM "auth"."provider_requests" WHERE id = $1;`,
      [id]
    );
  },

  providerRequest: async (id: string) => {
    const { rows } = await pgClient.query<{ options: SessionData }>(
      `SELECT options FROM "auth"."provider_requests" WHERE id = $1;`,
      [id]
    );
    return rows[0];
  },

  insertRefreshToken: async (
    userId: string,
    refreshToken: string = uuidv4()
  ) => {
    await pgClient.query(
      `INSERT INTO "auth"."refresh_tokens" (user_id, refresh_token, expires_at) VALUES($1, $2, $3);`,
      [userId, refreshToken, new Date(newRefreshExpiry())]
    );
    return refreshToken;
  },

  deleteRefreshToken: async (refreshToken: string) => {
    await pgClient.query(
      `DELETE FROM "auth"."refresh_tokens" WHERE refresh_token_hash = $1;`,
      [hashRefreshToken(refreshToken)]
    );
  },

  deleteUserRefreshTokens: async (userId: string) => {
    await pgClient.query(
      `DELETE FROM "auth"."refresh_tokens" WHERE user_id = $1;`,
      [userId]
    );
  },

  deleteExpiredRefreshTokens: async () => {
    await pgClient.query(
      `DELETE FROM "auth"."refresh_tokens" WHERE expires_at < NOW();`
    );
  },

  upsertRoles: async (roles: string[]) => {
    const { rows } = await pgClient.query<{ role: string }>(
      `INSERT INTO "auth"."roles" (role) VALUES ${roles
        .map((_, i) => `($${i + 1})`)
        .join(', ')} ON CONFLICT DO NOTHING
        RETURNING role;`,
      roles
    );
    return rows.map((row) => row.role);
  },

  getUserSecurityKeys: async (userId: string) => {
    const client = await pool.connect();
    const { rows } = await client.query<UserSecurityKey>(
      `SELECT id, counter, credential_id, credential_public_key, transports FROM "auth"."user_security_keys" WHERE user_id = $1;`,
      [userId]
    );
    client.release();
    return rows;
  },

  getUserChallenge: async (userId: string) => {
    const client = await pool.connect();
    const { rows } = await client.query<{ webauthn_current_challenge: string }>(
      `SELECT webauthn_current_challenge FROM "auth"."users" WHERE id = $1;`,
      [userId]
    );
    client.release();
    return rows[0];
  },

  updateUserChallenge: async (userId: string, challenge: string) => {
    const client = await pool.connect();
    await client.query(
      `UPDATE "auth"."users" SET webauthn_current_challenge = $1 WHERE id = $2;`,
      [challenge, userId]
    );
    client.release();
  },

  addUserSecurityKey: async ({
    user_id,
    counter,
    credential_id,
    credential_public_key,
    nickname,
  }: Pick<
    UserSecurityKey,
    | 'user_id'
    | 'counter'
    | 'credential_id'
    | 'credential_public_key'
    | 'nickname'
  >) => {
    const client = await pool.connect();
    const {
      rows: [{ id }],
    } = await client.query<{ id: string }>(
      `INSERT INTO "auth"."user_security_keys" (user_id, counter, credential_id, credential_public_key, nickname) VALUES($1, $2, $3, $4, $5) RETURNING id;`,
      [user_id, counter, credential_id, credential_public_key, nickname]
    );
    client.release();
    return id;
  },

  updateUserSecurityKey: async (securityKeyId: string, counter: number) => {
    const client = await pool.connect();
    await client.query(
      `UPDATE "auth"."user_security_keys" SET counter = $1 WHERE id = $2;`,
      [counter, securityKeyId]
    );
    client.release();
  },

  insertUserRoles: async (userId: string, roles: string[]) => {
    const client = await pool.connect();
    await insertUserRoles(client, userId, roles);
    client.release();
  },

  deleteUserRolesByUserId: async (userId: string) => {
    const client = await pool.connect();
    await client.query(`DELETE FROM "auth"."user_roles" WHERE user_id = $1;`, [
      userId,
    ]);
    client.release();
  },

  deleteUser: async (userId: string) => {
    const client = await pool.connect();
    await client.query(`DELETE FROM "auth"."user_roles" WHERE user_id = $1;`, [
      userId,
    ]);
    await client.query(`DELETE FROM "auth"."users" WHERE id = $1;`, [userId]);
    client.release();
  },

  insertUserProviderToUser: async ({
    userId,
    providerId,
    providerUserId,
    refreshToken,
    accessToken,
  }: {
    userId: string;
    providerId: string;
    providerUserId: string;
    refreshToken?: string;
    accessToken?: string;
  }) => {
    const { rows } = await pgClient.query(
      `INSERT INTO "auth"."user_providers" (user_id, provider_id, provider_user_id, refresh_token, access_token) VALUES($1, $2, $3, $4, $5) RETURNING *;`,
      [userId, providerId, providerUserId, refreshToken, accessToken]
    );
    return rows[0];
  },

  updateRefreshTokenExpiresAt: async (refreshToken: string) => {
    await pgClient.query(
      `UPDATE "auth"."refresh_tokens" rt SET expires_at = $1 FROM "auth"."users" u 
        WHERE rt.user_id = u.id
          AND rt.refresh_token_hash = $2 
          AND rt.expires_at < NOW() 
          AND u.disabled = false ;`,
      [new Date(newRefreshExpiry()), hashRefreshToken(refreshToken)]
    );
  },

  getUserByRefreshToken: async (refreshToken: string) => {
    const { rows } = await pgClient.query<{ id: string }>(
      `SELECT u.id 
        FROM auth.refresh_tokens AS rt
        JOIN auth.users AS u ON rt.user_id = u.id
        WHERE rt.refresh_token_hash = $1
          AND u.disabled = false
          AND rt.expires_at > NOW();`,
      [hashRefreshToken(refreshToken)]
    );

    return  await pgClient.getUserById(rows[0]?.id);
  },

  getUserById: async (userId: string) => {
    const client = await pool.connect();
    const user = await getUserById(client, userId);
    client.release();
    return cameliseUser(user);
  },

  getUserByEmail: async (email: string) => {
    const client = await pool.connect();
    const {
      rows: [user],
    } = await client.query<SqlUser>(createUserQueryByColumn('email'), [email]);
    client.release();
    return cameliseUser(user);
  },

  getUserByPhoneNumber: async (phoneNumber: string) => {
    const client = await pool.connect();
    const {
      rows: [user],
    } = await client.query<SqlUser>(createUserQueryByColumn('phone_number'), [
      phoneNumber,
    ]);
    client.release();
    return cameliseUser(user);
  },

  getUserByPhoneNumberAndOtp: async (phoneNumber: string) => {
    const client = await pool.connect();
    const {
      rows: [user],
    } = await client.query<SqlUser>(
      createUserQueryWhere(
        `u.phone_number = $1 AND u.otp_method_last_used = 'sms' AND u.otp_hash_expires_at > NOW()`
      ),
      [phoneNumber]
    );
    client.release();
    return cameliseUser(user);
  },

  getUserByTicket: async (ticket: string) => {
    const client = await pool.connect();
    const {
      rows: [user],
    } = await client.query<SqlUser>(
      createUserQueryWhere(`u.ticket = $1 AND u.ticket_expires_at > NOW()`),
      [ticket]
    );
    client.release();
    return cameliseUser(user);
  },

  getUserByChallenge: async (challenge: string) => {
    const client = await pool.connect();
    const {
      rows: [user],
    } = await client.query<SqlUser>(
      createUserQueryByColumn('webauthn_current_challenge'),
      [challenge]
    );
    client.release();
    return cameliseUser(user);
  },

  getUserByProvider: async (providerId: string, providerUserId: string) => {
    const client = await pool.connect();
    const { rows } = await client.query<{ user_id?: string; id: string }>(
      `SELECT user_id, id FROM "auth"."user_providers" WHERE provider_id = $1 AND provider_user_id = $2;`,
      [providerId, providerUserId]
    );
    let user: SqlUser | null = null;
    const userId = rows[0]?.user_id;
    if (userId) {
      user = await getUserById(client, userId);
    }
    client.release();
    return { id: rows[0]?.id, user: cameliseUser(user) };
  },

  updateAuthUserprovider: async (
    id: string,
    {
      accessToken,
      refreshToken,
    }: { accessToken?: string; refreshToken?: string }
  ) => {
    const client = await pool.connect();
    await client.query(
      `UPDATE "auth"."user_providers" SET access_token = $1, refresh_token = $2 WHERE id = $3;`,
      [accessToken, refreshToken, id]
    );
    client.release();
  },

  insertUser: async (user: Partial<User>) => {
    const client = await pool.connect();
    const transformedUser = snakeiseUser(user) as Partial<SqlUser>;
    const { roles, ...rest } = transformedUser;
    const columns = Object.keys(rest);
    const values = Object.values(rest);
    const {
      rows: [insertedUser],
    } = await client.query<SqlUser>(
      `INSERT INTO "auth"."users" (${columns.join(',')}) VALUES(${[
        ...columns.keys(),
      ]
        .map((i) => `$${i + 1}`)
        .join(',')}) RETURNING *;`,
      values
    );
    if (roles) {
      await insertUserRoles(client, insertedUser.id, roles);
    }
    client.release();
    return cameliseUser(insertedUser);
  },

  updateUser: async ({ id, user }: { id: string; user: Partial<User> }) => {
    const transformedUser = snakeiseUser(user) as Partial<SqlUser>;
    const { roles, ...rest } = transformedUser;
    const client = await pool.connect();
    const columns = Object.keys(rest);
    const values = Object.values(rest);
    const { rows } = await client.query(
      `UPDATE "auth"."users" SET ${columns
        .map((key, i) => `${key} = $${i + 1}`)
        .join(',')} WHERE id = $${columns.length + 1} RETURNING *;`,
      [...values, id]
    );
    if (roles) {
      await insertUserRoles(client, id, roles);
    }
    client.release();
    return cameliseUser(rows[0]);
  },
};
