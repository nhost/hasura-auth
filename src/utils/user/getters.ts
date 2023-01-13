import { SessionUser } from '@/types';
import { pgClient } from '../postgres-client';

export const getSessionUser = async (params: {
  userId: string;
}): Promise<SessionUser> => {
  const {
    id,
    createdAt,
    roles,
    displayName,
    avatarUrl,
    locale,
    email,
    isAnonymous,
    defaultRole,
    metadata,
    emailVerified,
    phoneNumber,
    phoneNumberVerified,
    activeMfaType,
  } = await pgClient.getUserById(params.userId);

  return {
    id,
    createdAt,
    roles,
    displayName,
    avatarUrl,
    locale,
    email,
    isAnonymous,
    defaultRole,
    metadata,
    emailVerified,
    phoneNumber,
    phoneNumberVerified,
    activeMfaType,
  };
};

export const getUserByEmail = (email: string) => pgClient.getUserByEmail(email);

export const getUserByTicket = (ticket: string) =>
  pgClient.getUserByTicket(ticket);
