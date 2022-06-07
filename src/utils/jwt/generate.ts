import { JWT } from 'jose';
import { ClaimValueType, DBUser, JwtSecret } from '@/types';
import { ENV } from '../env';
import { getUserRoles } from '../user';

// const RSA_TYPES = ["RS256", "RS384", "RS512"];
const ALLOWED_JWT_TYPES = ['HS256', 'HS384', 'HS512'];

const jwt = JSON.parse(ENV.HASURA_GRAPHQL_JWT_SECRET) as JwtSecret;

if (!ALLOWED_JWT_TYPES.includes(jwt.type)) {
  throw new Error(`Invalid JWT type: ${jwt.type}`);
}

if (!jwt.key) {
  throw new Error('Empty JWT key');
}

/**
 * * Signs a payload with the existing JWT configuration
 */
export const sign = ({ payload, user }: { payload: object; user: DBUser }) => {
  const jwt = JSON.parse(ENV.HASURA_GRAPHQL_JWT_SECRET) as JwtSecret;

  return JWT.sign(payload, jwt.key, {
    algorithm: jwt.type,
    expiresIn: `${ENV.AUTH_ACCESS_TOKEN_EXPIRES_IN}s`,
    subject: user.id,
    issuer: jwt.issuer ? jwt.issuer : 'hasura-auth',
  });
};

/**
 * Create an object that contains all the permission variables of the user,
 * i.e. user-id, allowed-roles, default-role and the kebab-cased columns
 * of the public.tables columns defined in JWT_CUSTOM_FIELDS
 * @param jwt if true, add a 'x-hasura-' prefix to the property names, and stringifies the values (required by Hasura)
 */
const generateHasuraClaims = async (
  user: DBUser
): Promise<{
  [key: string]: ClaimValueType;
}> => {
  const userRoles = getUserRoles({ userId: user.id });

  const allowedRoles = (await userRoles).map((role) => role.role);

  // add user's default role to allowed roles
  if (!allowedRoles.includes(user.default_role)) {
    allowedRoles.push(user.default_role);
  }

  // const customClaims = await generateCustomClaims(user.id);
  return {
    // ...customClaims,
    [`x-hasura-allowed-roles`]: allowedRoles,
    [`x-hasura-default-role`]: user.default_role,
    [`x-hasura-user-id`]: user.id,
    [`x-hasura-user-is-anonymous`]: user.is_anonymous.toString(),
  };
};
/**
 * Create JWT ENV.
 */
export const createHasuraAccessToken = async (
  user: DBUser
): Promise<string> => {
  const jwt = JSON.parse(ENV.HASURA_GRAPHQL_JWT_SECRET) as JwtSecret;

  const jwtNameSpace = jwt.claims_namespace
    ? jwt.claims_namespace
    : 'https://hasura.io/jwt/claims';

  return sign({
    payload: {
      [jwtNameSpace]: await generateHasuraClaims(user),
    },
    user,
  });
};
