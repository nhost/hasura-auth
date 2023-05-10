import { User } from '@/types';
import { SessionData } from 'express-session';
import { Pool } from 'pg';
export * from './types';

import { ENV } from '../env';
import { SqlUser, UserSecurityKey } from './types';
import {
  cameliseUser,
  createUserQueryByColumn,
  createUserQueryWhere,
  getUserById,
  snakeiseUser,
} from './utils';

const pool = new Pool({
  max: ENV.AUTH_PG_POOL_MAX_CONNECTIONS,
  connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
  idleTimeoutMillis: 30000,
});

export const pgClient = {
  insertProviderRequest: async (id: string, options: SessionData) => {
    const client = await pool.connect();
    await client.query(
      `INSERT INTO "auth"."provider_requests" (id, options) VALUES($1, $2) 
        ON CONFLICT(id) DO UPDATE SET options = EXCLUDED.options;`,
      [id, options]
    );
    client.release();
  },

  deleteProviderRequest: async (id: string) => {
    const client = await pool.connect();
    await client.query(
      `DELETE FROM "auth"."provider_requests" WHERE id = $1;`,
      [id]
    );
    client.release();
  },

  providerRequest: async (id: string) => {
    const client = await pool.connect();
    const { rows } = await client.query<{ options: SessionData }>(
      `SELECT options FROM "auth"."provider_requests" WHERE id = $1;`,
      [id]
    );
    client.release();
    return rows[0];
  },

  insertRefreshToken: async (
    userId: string,
    refreshToken: string,
    expiresAt: Date
  ) => {
    const client = await pool.connect();
    await client.query(
      `INSERT INTO "auth"."refresh_tokens" (user_id, refresh_token, expires_at) VALUES($1, $2, $3);`,
      [userId, refreshToken, expiresAt]
    );
    client.release();
  },

  deleteRefreshToken: async (refreshToken: string) => {
    const client = await pool.connect();
    await client.query(
      `DELETE FROM "auth"."refresh_tokens" WHERE refresh_token = $1;`,
      [refreshToken]
    );
    client.release();
  },

  deleteUserRefreshTokens: async (userId: string) => {
    const client = await pool.connect();
    await client.query(
      `DELETE FROM "auth"."refresh_tokens" WHERE user_id = $1;`,
      [userId]
    );
    client.release();
  },

  deleteExpiredRefreshTokens: async () => {
    const client = await pool.connect();
    await client.query(
      `DELETE FROM "auth"."refresh_tokens" WHERE expires_at < NOW();`
    );
    client.release();
  },

  upsertRoles: async (roles: string[]) => {
    const client = await pool.connect();
    const { rows } = await client.query<{ role: string }>(
      `INSERT INTO "auth"."roles" (role) VALUES ${roles
        .map((role) => `('${role}')`)
        .join(', ')} ON CONFLICT DO NOTHING
        RETURNING role;`
    );
    client.release();
    return rows.map((row) => row.role);
  },

  getUserSecurityKeys: async (userId: string) => {
    const client = await pool.connect();
    const { rows } = await client.query<UserSecurityKey>(
      `SELECT id, counter, credential_id, credential_public_key, transports FROM "auth"."security_keys" WHERE user_id = $1;`,
      [userId]
    );
    client.release();
    return rows;
  },

  getUserChallenge: async (userId: string) => {
    const client = await pool.connect();
    const { rows } = await client.query<{ current_challenge: string }>(
      `SELECT current_challenge FROM "auth"."users" WHERE id = $1;`,
      [userId]
    );
    client.release();
    return rows[0];
  },

  updateUserChallenge: async (userId: string, challenge: string) => {
    const client = await pool.connect();
    await client.query(
      `UPDATE "auth"."users" SET current_challenge = $1 WHERE id = $2;`,
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
      `INSERT INTO "auth"."security_keys" (user_id, counter, credential_id, credential_public_key, nickname) VALUES($1, $2, $3, $4, $5) RETURNING id;`,
      [user_id, counter, credential_id, credential_public_key, nickname]
    );
    client.release();
    return id;
  },

  updateUserSecurityKey: async (securityKeyId: string, counter: number) => {
    const client = await pool.connect();
    await client.query(
      `UPDATE "auth"."security_keys" SET counter = $1 WHERE id = $2;`,
      [counter, securityKeyId]
    );
    client.release();
  },

  insertUserRoles: async (userId: string, roles: string[]) => {
    const client = await pool.connect();
    await client.query(
      `INSERT INTO "auth"."user_roles" (user_id, role) VALUES ${roles
        .map((role) => `('${userId}', '${role}')`)
        .join(', ')};`
    );
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
    const client = await pool.connect();
    const { rows } = await client.query(
      `INSERT INTO "auth"."user_providers" (user_id, provider_id, provider_user_id, refresh_token, access_token) VALUES($1, $2, $3, $4, $5, $6);`,
      [userId, providerId, providerUserId, refreshToken, accessToken]
    );
    client.release();
    return rows[0];
  },

  updateRefreshTokenExpiresAt: async (
    refreshToken: string,
    expires_at: Date
  ) => {
    const client = await pool.connect();
    await client.query(
      `UPDATE "auth"."refresh_tokens" rt SET expires_at = $1 FROM "auth"."users" u 
        WHERE rt.user_id = u.id
          AND rt.refresh_token = $2 
          AND rt.expires_at < NOW() 
          AND u.disabled = false ;`,
      [expires_at, refreshToken]
    );
    client.release();
  },

  getUserByRefreshToken: async (refreshToken: string) => {
    const client = await pool.connect();
    const { rows } = await client.query<{ id: string }>(
      `SELECT u.id FROM "auth"."refresh_tokens" rt, "auth"."users" u
        WHERE rt.user_id = u.id
          AND rt.refresh_token = $1
          AND u.disabled = false
          AND rt.expires_at > NOW();`,
      [refreshToken]
    );

    const user = await getUserById(client, rows[0]?.id);
    client.release();
    return cameliseUser(user);
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
      createUserQueryWhere(`u.ticket = $1 and u.ticket_expires_at > NOW()`),
      [ticket]
    );
    client.release();
    return cameliseUser(user);
  },

  getUserByChallence: async (challenge: string) => {
    const client = await pool.connect();
    const {
      rows: [user],
    } = await client.query<SqlUser>(
      createUserQueryByColumn('current_challenge'),
      [challenge]
    );
    client.release();
    return cameliseUser(user);
  },

  getUserByProvider: async (providerId: string, providerUserId: string) => {
    const client = await pool.connect();
    const { rows } = await client.query<{ user_id?: string; id: string }>(
      `SELECT user_id, id FROM "auth"."user_providers" WHERE up.provider_id = $1 AND up.provider_user_id = $2;`,
      [providerId, providerUserId]
    );
    let user: SqlUser | null = null;
    const userId = rows[0]?.user_id;
    if (userId) {
      user = await getUserById(client, userId);
    }
    client.release();
    return { id: rows[0].id, user: cameliseUser(user) };
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

  insertUser: async (user: Partial<SqlUser>) => {
    const client = await pool.connect();
    const { roles, ...rest } = user;
    const columns = Object.keys(rest);
    const values = Object.values(rest);
    const { rows } = await client.query<SqlUser>(
      `INSERT INTO "auth"."users" (${columns.join(',')}) VALUES(${[
        ...columns.keys(),
      ]
        .map((i) => `$${i + 1}`)
        .join(',')}) RETURNING *;`,
      values
    );
    // TODO roles
    console.log('TODO roles', roles);

    client.release();
    return cameliseUser(rows[0]);
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
    // TODO roles
    console.log('TODO roles', roles);
    client.release();
    return cameliseUser(rows[0]);
  },
};
