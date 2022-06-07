import { DBUser, DBUserRole, User } from '@/types';
import { gqlSdk } from '../gql-sdk';
import { postgres } from '../postgres';

export const getUserByPhoneNumber = async ({
  phoneNumber,
}: {
  phoneNumber: string;
}) => {
  const { users } = await gqlSdk.users({
    where: {
      phoneNumber: {
        _eq: phoneNumber,
      },
    },
  });

  return users[0];
};

export const getUserRoles = async ({ userId }: { userId: string }) => {
  // get user roles
  const { rows } = await postgres.runSqlParsed(
    `SELECT row_to_json(ur) FROM auth.user_roles ur WHERE user_id = %L`,
    [userId]
  );

  const userRoles = rows as DBUserRole[];

  return userRoles;
};

export const getUser = async ({
  userId,
}: {
  userId: string;
}): Promise<User> => {
  // get user
  const { rows: users } = await postgres.runSqlParsed(
    `SELECT row_to_json(u) FROM auth.users u WHERE id = %L LIMIT 1`,
    [userId]
  );

  if (users.length === 0) {
    throw new Error('Unable to get user');
  }

  const user = users[0] as DBUser;

  const userRoles = await getUserRoles({ userId });

  // construct return user
  return {
    id: user.id,
    createdAt: user.created_at,
    displayName: user.display_name,
    avatarUrl: user.avatar_url,
    locale: user.locale,
    email: user.email,
    isAnonymous: user.is_anonymous,
    defaultRole: user.default_role,
    metadata: user.metadata,
    emailVerified: user.email_verified,
    phoneNumber: user.phone_number,
    phoneNumberVerified: user.phone_number_verified,
    activeMfaType: user.active_mfa_type,
    roles: userRoles.map((role) => role.role),
  };
};

export const getUserByEmail = async (email: string): Promise<DBUser | null> => {
  const { rows } = await postgres.runSqlParsed(
    `SELECT row_to_json(u) FROM auth.users u WHERE email = %L LIMIT 1`,
    [email]
  );

  if (rows.length === 0) {
    return null;
  }

  const users = rows as DBUser[];
  const user = users[0];

  return user;
};

export const getUserByTicket = async (ticket: string) => {
  const now = new Date();

  const { users } = await gqlSdk.users({
    where: {
      _and: [
        {
          ticket: {
            _eq: ticket,
          },
        },
        {
          ticketExpiresAt: {
            _gt: now,
          },
        },
      ],
    },
  });

  if (users.length !== 1) {
    return null;
  }

  return users[0];
};

export const getUserByRefreshToken = async (
  refreshToken: string
): Promise<DBUser | null> => {
  const { rows } = await postgres.runSqlParsed(
    `SELECT row_to_json(u) FROM auth.users u JOIN auth.refresh_tokens rt ON rt.user_id = u.id WHERE rt.refresh_token = %L AND rt.expires_at > NOW() AND u.disabled = FALSE;`,
    [refreshToken]
  );

  if (rows.length === 0) {
    return null;
  }

  const users = rows as DBUser[];
  const user = users[0];

  return user;
};
