import { UserQuery } from './utils/__generated__/graphql-request';

export type ClaimValueType =
  | string
  | string[]
  | number
  | number[]
  | RegExp
  | RegExp[]
  | boolean
  | boolean[]
  | null
  | undefined;

/**
 * Claims interface.
 */
export interface Claims {
  'x-hasura-user-id': string;
  'x-hasura-default-role': string;
  'x-hasura-allowed-roles': string[];
  [key: string]: ClaimValueType;
}

/**
 * PermissionVariables interface.
 */
export interface PermissionVariables {
  'user-id': string;
  'default-role': string;
  'allowed-roles': string[];
  [key: string]: ClaimValueType;
}

/**
 * Token interface.
 */
export type Token = {
  [key: string]: Claims;
} & {
  'https://hasura.io/jwt/claims': Claims;
  exp: bigint;
  iat: bigint;
  iss: string;
  sub: string;
};

// Session and user
type Metadata = Record<string, unknown>;

export type UserRegistrationOptions = {
  locale: string;
  allowedRoles: string[];
  defaultRole: string;
  displayName?: string;
  metadata: Metadata;
};

export type User = Pick<
  NonNullable<UserQuery['user']>,
  | 'id'
  | 'createdAt'
  | 'displayName'
  | 'avatarUrl'
  | 'locale'
  | 'email'
  | 'isAnonymous'
  | 'defaultRole'
  | 'metadata'
  | 'emailVerified'
  | 'phoneNumber'
  | 'phoneNumberVerified'
  | 'activeMfaType'
> & { roles: string[] };

export type Session = {
  accessToken: string;
  accessTokenExpiresIn: number;
  refreshToken: string;
  user?: User;
};

export type Mfa = {
  ticket: string | null;
};

export type SignInResponse = {
  session: Session | null;
  mfa: Mfa | null;
};

export type PasswordLessEmailBody = {
  email: string;
  options: UserRegistrationOptions & {
    redirectTo: string;
  };
};

export type PasswordLessSmsBody = {
  phoneNumber: string;
  options: UserRegistrationOptions;
};

export type OtpSmsBody = {
  phoneNumber: string;
  otp: string;
};

export type JwtSecret = {
  type: 'HS256' | 'HS238' | 'HS512' | 'RS256' | 'RS384' | 'RS512' | 'Ed25519';
  key: string;
  jwk_url?: string;
  claims_namespace?: string;
  claims_namespace_path?: string;
  claims_format?: string;
  audience?: string;
  issuer?: string;
  claims_map?: string;
  allowed_skew?: string;
  header?: string;
};

export const EMAIL_TYPES = {
  VERIFY: 'emailVerify',
  CONFIRM_CHANGE: 'emailConfirmChange',
  SIGNIN_PASSWORDLESS: 'signinPasswordless',
  PASSWORD_RESET: 'passwordReset',
} as const;
export type EmailType = typeof EMAIL_TYPES[keyof typeof EMAIL_TYPES];

// database types

type JSONValue =
  | string
  | number
  | boolean
  | { [x: string]: JSONValue }
  | Array<JSONValue>;

export type DBUser = {
  id: string;
  created_at: string;
  updated_at: string;
  last_seen: string | null;
  disabled: boolean;
  display_name: string;
  avatar_url: string;
  locale: string;
  email: string;
  phone_number: string | null;
  password_hash: string | null;
  email_verified: boolean;
  phone_number_verified: boolean;
  new_email: string | null;
  otp_method_last_used: string | null;
  otp_hash: string | null;
  otp_hash_expires_at: string | null;
  default_role: string;
  is_anonymous: boolean;
  totp_secret: string | null;
  active_mfa_type: string | null;
  ticket: string | null;
  ticket_expires_at: string | null;
  metadata: JSONValue;
};

export type DBUserRole = {
  id: string;
  created_at: string;
  user_id: string;
  role: string;
};
