import { v4 as uuidv4 } from 'uuid';
import { Session, SignInResponse, User } from '@/types';
import { generateTicketExpiresAt } from './ticket';
import { ENV } from './env';
import { createHasuraAccessToken } from './jwt';
import { pgClient } from './postgres-client';

/**
 * Get new or update current user session
 *
 * @param userAndToken - User field fragment and current refresh token if any
 * @returns Returns new user session if no valid current refresh token is passed, otherwise update current session
 */
export const getNewOrUpdateCurrentSession = async ({
  user,
  currentRefreshToken,
}: {
  user: User;
  currentRefreshToken?: string;
}): Promise<Session> => {
  // update user's last seen
  await pgClient.updateUser({
    id: user.id,
    user: {
      // TODO performance: update through a SQL trigger
      lastSeen: new Date(),
    },
  });
  const accessToken = await createHasuraAccessToken(user);
  // TODO performance: no need to wait for the update to finish the http request. How to do this? Queue with an hasura event on the last_seen column?
  // Another option: update both the user and the refresh token in the same transaction
  // TODO performance: upsert refresh token rather than this scrappy update/insert
  const refreshToken =
    (currentRefreshToken &&
      (await pgClient.updateRefreshTokenExpiresAt(currentRefreshToken))) ||
    (await pgClient.insertRefreshToken(user.id));
  return {
    accessToken,
    accessTokenExpiresIn: ENV.AUTH_ACCESS_TOKEN_EXPIRES_IN,
    refreshToken,
    user,
  };
};

export const getSignInResponse = async ({
  user,
  checkMFA,
}: {
  user: User;
  checkMFA: boolean;
}): Promise<SignInResponse> => {
  if (checkMFA && user?.activeMfaType === 'totp') {
    // generate new ticket
    const ticket = `mfaTotp:${uuidv4()}`;
    // set ticket
    await pgClient.updateUser({
      id: user.id,
      user: {
        ticket,
        ticketExpiresAt: generateTicketExpiresAt(5 * 60),
      },
    });
    return {
      session: null,
      mfa: {
        ticket,
      },
    };
  }
  const session = await getNewOrUpdateCurrentSession({ user });
  return {
    session,
    mfa: null,
  };
};
