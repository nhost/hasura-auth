import { v4 as uuidv4 } from 'uuid';
import { gqlSdk } from '@/utils';
import { ENV } from './env';
import crypto from 'crypto';

const hash = (value: string) =>
  crypto.createHash('sha256').update(value).digest('hex');

export const getUserByRefreshToken = async (refreshToken: string) => {
  return (
    await gqlSdk.getUsersByRefreshToken({
      refreshTokens: [refreshToken, hash(refreshToken)],
    })
  ).authRefreshTokens[0]?.user;
};

export const deleteUserRefreshTokens = async (userId: string) => {
  await gqlSdk.deleteUserRefreshTokens({ userId });
};

export const deleteRefreshToken = async (refreshToken: string) => {
  // * delete both refresh token and its hash value
  await gqlSdk.deleteRefreshTokens({
    refreshTokens: [refreshToken, hash(refreshToken)],
  });
};

const newRefreshExpiry = () => {
  const date = new Date();

  // cant return this becuase this will return a unix timestamp directly
  date.setSeconds(date.getSeconds() + ENV.AUTH_REFRESH_TOKEN_EXPIRES_IN);

  // instead we must return the js date object
  return date;
};

export const updateRefreshTokenExpiry = async (refreshToken: string) => {
  await gqlSdk.getUsersByRefreshTokenAndUpdateRefreshTokenExpiresAt({
    refreshTokens: [refreshToken, hash(refreshToken)],
    expiresAt: new Date(newRefreshExpiry()),
  });

  return refreshToken;
};

export const getNewRefreshToken = async (
  userId: string,
  refreshToken = uuidv4()
) => {
  await gqlSdk.insertRefreshToken({
    refreshToken: {
      userId,
      refreshToken: hash(refreshToken),
      expiresAt: new Date(newRefreshExpiry()),
    },
  });

  return refreshToken;
};
//
