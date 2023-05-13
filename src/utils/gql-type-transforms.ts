import {
  AuthRefreshTokens_Insert_Input,
  AuthUserSecurityKeys_Insert_Input,
  AuthRoles_Insert_Input,
  AuthUserRoles_Insert_Input,
  AuthUserProviders_Insert_Input,
  Users_Constraint,
  Users_Insert_Input,
  Users_Obj_Rel_Insert_Input,
  Users_Update_Column,
  UsersQueryVariables,
} from './__generated__/graphql-request-hasura-default';
import {
  AuthRefreshTokensInsertInput,
  AuthRolesInsertInput,
  AuthUserProvidersInsertInput,
  AuthUserRolesInsertInput,
  AuthUserSecurityKeysInsertInput,
  InputMaybe,
  UsersBoolExp,
  UsersConstraint,
  UsersInsertInput,
  UsersObjRelInsertInput,
  UsersOnConflict,
  UsersQueryVariables as UsersQueryVariablesGraphql,
  UsersUpdateColumn,
} from './__generated__/graphql-request-graphql-default';

type ValueOf<T> = T[keyof T];

function transformEnum<
  R extends string | number,
  FromEnumType extends { [key: string]: R },
  ToEnumType extends { [key: string]: R }
>(
  fromEnum: FromEnumType,
  toEnum: ToEnumType,
  enumValue: ValueOf<FromEnumType>
): ValueOf<ToEnumType> {
  const keys = Object.keys(fromEnum).filter((x) => fromEnum[x] == enumValue);
  const key = keys.length > 0 ? keys[0] : '';
  return toEnum[key as keyof typeof toEnum];
}

export const transformUsersQueryVariables = (
  variables: UsersQueryVariables
): UsersQueryVariablesGraphql => {
  // Users_Bool_Exp actually matches up with UsersBoolExp, but the the type system
  // doesn't know that because of the named types of relationships like "roles"
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  return { ...variables, where: variables.where as any };
};

export const transformUsersInsertInput = (
  user: Users_Insert_Input
): UsersInsertInput => {
  return {
    ...user,
    defaultRoleByRole: null, // TODO: transform
    refreshTokens: null, // TODO: transform
    roles: null, // TODO: transform
    securityKeys: null, // TODO: transform
    userProviders: null, // TODO: transform
  };
};

const transformUsersObjRelInsertInput = (
  userRelation: InputMaybe<Users_Obj_Rel_Insert_Input> | undefined
): InputMaybe<UsersObjRelInsertInput> => {
  if (userRelation?.data) {
    const data = transformUsersInsertInput(userRelation.data);
    if (userRelation.on_conflict) {
      const onConflict: UsersOnConflict = {
        ...userRelation.on_conflict,
        constraint: transformEnum(
          Users_Constraint,
          UsersConstraint,
          userRelation.on_conflict.constraint
        ),
        update_columns: userRelation.on_conflict.update_columns?.map((x) =>
          transformEnum(Users_Update_Column, UsersUpdateColumn, x)
        ),
        where: userRelation.on_conflict.where as UsersBoolExp,
      };
      return {
        data,
        onConflict,
      };
    }
    return {
      data,
    };
  }
  return null;
};

export const transformAuthRefreshTokensInsertInput = (
  refreshToken: AuthRefreshTokens_Insert_Input
): AuthRefreshTokensInsertInput => {
  return {
    ...refreshToken,
    type: null, // TODO: transform
    refreshTokenType: null, // TODO: transform
    user: transformUsersObjRelInsertInput(refreshToken.user),
  };
};

export const transformAuthUserSecurityKeysInsertInput = (
  userSecurityKey: AuthUserSecurityKeys_Insert_Input
): AuthUserSecurityKeysInsertInput => {
  return {
    ...userSecurityKey,
    user: transformUsersObjRelInsertInput(userSecurityKey.user),
  };
};

const transformAuthUserRolesInsertInputOne = (
  userRole: AuthUserRoles_Insert_Input
): AuthUserRolesInsertInput => {
  return {
    ...userRole,
    roleByRole: null, // TODO: transform
    user: transformUsersObjRelInsertInput(userRole.user),
  };
};

export const transformAuthUserRolesInsertInput = (
  userRoles: AuthUserRoles_Insert_Input | AuthUserRoles_Insert_Input[]
): AuthUserRolesInsertInput | AuthUserRolesInsertInput[] => {
  if (Array.isArray(userRoles)) {
    return userRoles.map(transformAuthUserRolesInsertInputOne);
  }
  return transformAuthUserRolesInsertInputOne(userRoles);
};

const transformAuthRolesInsertInputOne = (
  role: AuthRoles_Insert_Input
): AuthRolesInsertInput => {
  return {
    ...role,
  };
};

export const transformAuthRolesInsertInput = (
  roles: AuthRoles_Insert_Input | AuthRoles_Insert_Input[]
): AuthRolesInsertInput | AuthRolesInsertInput[] => {
  if (Array.isArray(roles)) {
    return roles.map(transformAuthRolesInsertInputOne);
  }
  return transformAuthRolesInsertInputOne(roles);
};

export const transformAuthUserProvidersInsertInput = (
  authUserProviders: AuthUserProviders_Insert_Input
): AuthUserProvidersInsertInput => {
  return {
    ...authUserProviders,
    user: transformUsersObjRelInsertInput(authUserProviders.user),
    provider: null, // TODO: transform
  };
};
