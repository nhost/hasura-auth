/* eslint-disable @typescript-eslint/no-explicit-any, @typescript-eslint/no-unused-vars */
import { GraphQLClient } from 'graphql-request';
import * as Dom from 'graphql-request/dist/types.dom';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  citext: any;
  date: any;
  jsonb: any;
  lquery: any;
  ltree: any;
  ltxtquery: any;
  timestamptz: any;
  uuid: any;
};

/** Boolean expression to compare columns of type "Boolean". All fields are combined with logical 'AND'. */
export type Boolean_Comparison_Exp = {
  _eq?: InputMaybe<Scalars['Boolean']>;
  _gt?: InputMaybe<Scalars['Boolean']>;
  _gte?: InputMaybe<Scalars['Boolean']>;
  _in?: InputMaybe<Array<Scalars['Boolean']>>;
  _is_null?: InputMaybe<Scalars['Boolean']>;
  _lt?: InputMaybe<Scalars['Boolean']>;
  _lte?: InputMaybe<Scalars['Boolean']>;
  _neq?: InputMaybe<Scalars['Boolean']>;
  _nin?: InputMaybe<Array<Scalars['Boolean']>>;
};

/** Boolean expression to compare columns of type "String". All fields are combined with logical 'AND'. */
export type String_Comparison_Exp = {
  _eq?: InputMaybe<Scalars['String']>;
  _gt?: InputMaybe<Scalars['String']>;
  _gte?: InputMaybe<Scalars['String']>;
  /** does the column match the given case-insensitive pattern */
  _ilike?: InputMaybe<Scalars['String']>;
  _in?: InputMaybe<Array<Scalars['String']>>;
  /** does the column match the given POSIX regular expression, case insensitive */
  _iregex?: InputMaybe<Scalars['String']>;
  _is_null?: InputMaybe<Scalars['Boolean']>;
  /** does the column match the given pattern */
  _like?: InputMaybe<Scalars['String']>;
  _lt?: InputMaybe<Scalars['String']>;
  _lte?: InputMaybe<Scalars['String']>;
  _neq?: InputMaybe<Scalars['String']>;
  /** does the column NOT match the given case-insensitive pattern */
  _nilike?: InputMaybe<Scalars['String']>;
  _nin?: InputMaybe<Array<Scalars['String']>>;
  /** does the column NOT match the given POSIX regular expression, case insensitive */
  _niregex?: InputMaybe<Scalars['String']>;
  /** does the column NOT match the given pattern */
  _nlike?: InputMaybe<Scalars['String']>;
  /** does the column NOT match the given POSIX regular expression, case sensitive */
  _nregex?: InputMaybe<Scalars['String']>;
  /** does the column NOT match the given SQL regular expression */
  _nsimilar?: InputMaybe<Scalars['String']>;
  /** does the column match the given POSIX regular expression, case sensitive */
  _regex?: InputMaybe<Scalars['String']>;
  /** does the column match the given SQL regular expression */
  _similar?: InputMaybe<Scalars['String']>;
};

/** columns and relationships of "auth.provider_requests" */
export type AuthProviderRequests = {
  __typename?: 'authProviderRequests';
  id: Scalars['uuid'];
  options?: Maybe<Scalars['jsonb']>;
};


/** columns and relationships of "auth.provider_requests" */
export type AuthProviderRequestsOptionsArgs = {
  path?: InputMaybe<Scalars['String']>;
};

/** aggregated selection of "auth.provider_requests" */
export type AuthProviderRequests_Aggregate = {
  __typename?: 'authProviderRequests_aggregate';
  aggregate?: Maybe<AuthProviderRequests_Aggregate_Fields>;
  nodes: Array<AuthProviderRequests>;
};

/** aggregate fields of "auth.provider_requests" */
export type AuthProviderRequests_Aggregate_Fields = {
  __typename?: 'authProviderRequests_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<AuthProviderRequests_Max_Fields>;
  min?: Maybe<AuthProviderRequests_Min_Fields>;
};


/** aggregate fields of "auth.provider_requests" */
export type AuthProviderRequests_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<AuthProviderRequests_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** append existing jsonb value of filtered columns with new jsonb value */
export type AuthProviderRequests_Append_Input = {
  options?: InputMaybe<Scalars['jsonb']>;
};

/** Boolean expression to filter rows from the table "auth.provider_requests". All fields are combined with a logical 'AND'. */
export type AuthProviderRequests_Bool_Exp = {
  _and?: InputMaybe<Array<AuthProviderRequests_Bool_Exp>>;
  _not?: InputMaybe<AuthProviderRequests_Bool_Exp>;
  _or?: InputMaybe<Array<AuthProviderRequests_Bool_Exp>>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  options?: InputMaybe<Jsonb_Comparison_Exp>;
};

/** unique or primary key constraints on table "auth.provider_requests" */
export enum AuthProviderRequests_Constraint {
  /** unique or primary key constraint */
  ProviderRequestsPkey = 'provider_requests_pkey'
}

/** delete the field or element with specified path (for JSON arrays, negative integers count from the end) */
export type AuthProviderRequests_Delete_At_Path_Input = {
  options?: InputMaybe<Array<Scalars['String']>>;
};

/** delete the array element with specified index (negative integers count from the end). throws an error if top level container is not an array */
export type AuthProviderRequests_Delete_Elem_Input = {
  options?: InputMaybe<Scalars['Int']>;
};

/** delete key/value pair or string element. key/value pairs are matched based on their key value */
export type AuthProviderRequests_Delete_Key_Input = {
  options?: InputMaybe<Scalars['String']>;
};

/** input type for inserting data into table "auth.provider_requests" */
export type AuthProviderRequests_Insert_Input = {
  id?: InputMaybe<Scalars['uuid']>;
  options?: InputMaybe<Scalars['jsonb']>;
};

/** aggregate max on columns */
export type AuthProviderRequests_Max_Fields = {
  __typename?: 'authProviderRequests_max_fields';
  id?: Maybe<Scalars['uuid']>;
};

/** aggregate min on columns */
export type AuthProviderRequests_Min_Fields = {
  __typename?: 'authProviderRequests_min_fields';
  id?: Maybe<Scalars['uuid']>;
};

/** response of any mutation on the table "auth.provider_requests" */
export type AuthProviderRequests_Mutation_Response = {
  __typename?: 'authProviderRequests_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<AuthProviderRequests>;
};

/** on conflict condition type for table "auth.provider_requests" */
export type AuthProviderRequests_On_Conflict = {
  constraint: AuthProviderRequests_Constraint;
  update_columns?: Array<AuthProviderRequests_Update_Column>;
  where?: InputMaybe<AuthProviderRequests_Bool_Exp>;
};

/** Ordering options when selecting data from "auth.provider_requests". */
export type AuthProviderRequests_Order_By = {
  id?: InputMaybe<Order_By>;
  options?: InputMaybe<Order_By>;
};

/** primary key columns input for table: authProviderRequests */
export type AuthProviderRequests_Pk_Columns_Input = {
  id: Scalars['uuid'];
};

/** prepend existing jsonb value of filtered columns with new jsonb value */
export type AuthProviderRequests_Prepend_Input = {
  options?: InputMaybe<Scalars['jsonb']>;
};

/** select columns of table "auth.provider_requests" */
export enum AuthProviderRequests_Select_Column {
  /** column name */
  Id = 'id',
  /** column name */
  Options = 'options'
}

/** input type for updating data in table "auth.provider_requests" */
export type AuthProviderRequests_Set_Input = {
  id?: InputMaybe<Scalars['uuid']>;
  options?: InputMaybe<Scalars['jsonb']>;
};

/** update columns of table "auth.provider_requests" */
export enum AuthProviderRequests_Update_Column {
  /** column name */
  Id = 'id',
  /** column name */
  Options = 'options'
}

/** columns and relationships of "auth.providers" */
export type AuthProviders = {
  __typename?: 'authProviders';
  id: Scalars['String'];
  /** An array relationship */
  userProviders: Array<AuthUserProviders>;
  /** An aggregate relationship */
  userProviders_aggregate: AuthUserProviders_Aggregate;
};


/** columns and relationships of "auth.providers" */
export type AuthProvidersUserProvidersArgs = {
  distinct_on?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserProviders_Order_By>>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};


/** columns and relationships of "auth.providers" */
export type AuthProvidersUserProviders_AggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserProviders_Order_By>>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};

/** aggregated selection of "auth.providers" */
export type AuthProviders_Aggregate = {
  __typename?: 'authProviders_aggregate';
  aggregate?: Maybe<AuthProviders_Aggregate_Fields>;
  nodes: Array<AuthProviders>;
};

/** aggregate fields of "auth.providers" */
export type AuthProviders_Aggregate_Fields = {
  __typename?: 'authProviders_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<AuthProviders_Max_Fields>;
  min?: Maybe<AuthProviders_Min_Fields>;
};


/** aggregate fields of "auth.providers" */
export type AuthProviders_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<AuthProviders_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** Boolean expression to filter rows from the table "auth.providers". All fields are combined with a logical 'AND'. */
export type AuthProviders_Bool_Exp = {
  _and?: InputMaybe<Array<AuthProviders_Bool_Exp>>;
  _not?: InputMaybe<AuthProviders_Bool_Exp>;
  _or?: InputMaybe<Array<AuthProviders_Bool_Exp>>;
  id?: InputMaybe<String_Comparison_Exp>;
  userProviders?: InputMaybe<AuthUserProviders_Bool_Exp>;
};

/** unique or primary key constraints on table "auth.providers" */
export enum AuthProviders_Constraint {
  /** unique or primary key constraint */
  ProvidersPkey = 'providers_pkey'
}

/** input type for inserting data into table "auth.providers" */
export type AuthProviders_Insert_Input = {
  id?: InputMaybe<Scalars['String']>;
  userProviders?: InputMaybe<AuthUserProviders_Arr_Rel_Insert_Input>;
};

/** aggregate max on columns */
export type AuthProviders_Max_Fields = {
  __typename?: 'authProviders_max_fields';
  id?: Maybe<Scalars['String']>;
};

/** aggregate min on columns */
export type AuthProviders_Min_Fields = {
  __typename?: 'authProviders_min_fields';
  id?: Maybe<Scalars['String']>;
};

/** response of any mutation on the table "auth.providers" */
export type AuthProviders_Mutation_Response = {
  __typename?: 'authProviders_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<AuthProviders>;
};

/** input type for inserting object relation for remote table "auth.providers" */
export type AuthProviders_Obj_Rel_Insert_Input = {
  data: AuthProviders_Insert_Input;
  /** on conflict condition */
  on_conflict?: InputMaybe<AuthProviders_On_Conflict>;
};

/** on conflict condition type for table "auth.providers" */
export type AuthProviders_On_Conflict = {
  constraint: AuthProviders_Constraint;
  update_columns?: Array<AuthProviders_Update_Column>;
  where?: InputMaybe<AuthProviders_Bool_Exp>;
};

/** Ordering options when selecting data from "auth.providers". */
export type AuthProviders_Order_By = {
  id?: InputMaybe<Order_By>;
  userProviders_aggregate?: InputMaybe<AuthUserProviders_Aggregate_Order_By>;
};

/** primary key columns input for table: authProviders */
export type AuthProviders_Pk_Columns_Input = {
  id: Scalars['String'];
};

/** select columns of table "auth.providers" */
export enum AuthProviders_Select_Column {
  /** column name */
  Id = 'id'
}

/** input type for updating data in table "auth.providers" */
export type AuthProviders_Set_Input = {
  id?: InputMaybe<Scalars['String']>;
};

/** update columns of table "auth.providers" */
export enum AuthProviders_Update_Column {
  /** column name */
  Id = 'id'
}

/** columns and relationships of "auth.refresh_tokens" */
export type AuthRefreshTokens = {
  __typename?: 'authRefreshTokens';
  createdAt: Scalars['timestamptz'];
  expiresAt: Scalars['timestamptz'];
  refreshToken: Scalars['uuid'];
  /** An object relationship */
  user: Users;
  userId: Scalars['uuid'];
};

/** aggregated selection of "auth.refresh_tokens" */
export type AuthRefreshTokens_Aggregate = {
  __typename?: 'authRefreshTokens_aggregate';
  aggregate?: Maybe<AuthRefreshTokens_Aggregate_Fields>;
  nodes: Array<AuthRefreshTokens>;
};

/** aggregate fields of "auth.refresh_tokens" */
export type AuthRefreshTokens_Aggregate_Fields = {
  __typename?: 'authRefreshTokens_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<AuthRefreshTokens_Max_Fields>;
  min?: Maybe<AuthRefreshTokens_Min_Fields>;
};


/** aggregate fields of "auth.refresh_tokens" */
export type AuthRefreshTokens_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<AuthRefreshTokens_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "auth.refresh_tokens" */
export type AuthRefreshTokens_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<AuthRefreshTokens_Max_Order_By>;
  min?: InputMaybe<AuthRefreshTokens_Min_Order_By>;
};

/** input type for inserting array relation for remote table "auth.refresh_tokens" */
export type AuthRefreshTokens_Arr_Rel_Insert_Input = {
  data: Array<AuthRefreshTokens_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<AuthRefreshTokens_On_Conflict>;
};

/** Boolean expression to filter rows from the table "auth.refresh_tokens". All fields are combined with a logical 'AND'. */
export type AuthRefreshTokens_Bool_Exp = {
  _and?: InputMaybe<Array<AuthRefreshTokens_Bool_Exp>>;
  _not?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
  _or?: InputMaybe<Array<AuthRefreshTokens_Bool_Exp>>;
  createdAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  expiresAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  refreshToken?: InputMaybe<Uuid_Comparison_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  userId?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "auth.refresh_tokens" */
export enum AuthRefreshTokens_Constraint {
  /** unique or primary key constraint */
  RefreshTokensPkey = 'refresh_tokens_pkey'
}

/** input type for inserting data into table "auth.refresh_tokens" */
export type AuthRefreshTokens_Insert_Input = {
  createdAt?: InputMaybe<Scalars['timestamptz']>;
  expiresAt?: InputMaybe<Scalars['timestamptz']>;
  refreshToken?: InputMaybe<Scalars['uuid']>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  userId?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type AuthRefreshTokens_Max_Fields = {
  __typename?: 'authRefreshTokens_max_fields';
  createdAt?: Maybe<Scalars['timestamptz']>;
  expiresAt?: Maybe<Scalars['timestamptz']>;
  refreshToken?: Maybe<Scalars['uuid']>;
  userId?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "auth.refresh_tokens" */
export type AuthRefreshTokens_Max_Order_By = {
  createdAt?: InputMaybe<Order_By>;
  expiresAt?: InputMaybe<Order_By>;
  refreshToken?: InputMaybe<Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type AuthRefreshTokens_Min_Fields = {
  __typename?: 'authRefreshTokens_min_fields';
  createdAt?: Maybe<Scalars['timestamptz']>;
  expiresAt?: Maybe<Scalars['timestamptz']>;
  refreshToken?: Maybe<Scalars['uuid']>;
  userId?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "auth.refresh_tokens" */
export type AuthRefreshTokens_Min_Order_By = {
  createdAt?: InputMaybe<Order_By>;
  expiresAt?: InputMaybe<Order_By>;
  refreshToken?: InputMaybe<Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "auth.refresh_tokens" */
export type AuthRefreshTokens_Mutation_Response = {
  __typename?: 'authRefreshTokens_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<AuthRefreshTokens>;
};

/** on conflict condition type for table "auth.refresh_tokens" */
export type AuthRefreshTokens_On_Conflict = {
  constraint: AuthRefreshTokens_Constraint;
  update_columns?: Array<AuthRefreshTokens_Update_Column>;
  where?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
};

/** Ordering options when selecting data from "auth.refresh_tokens". */
export type AuthRefreshTokens_Order_By = {
  createdAt?: InputMaybe<Order_By>;
  expiresAt?: InputMaybe<Order_By>;
  refreshToken?: InputMaybe<Order_By>;
  user?: InputMaybe<Users_Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** primary key columns input for table: authRefreshTokens */
export type AuthRefreshTokens_Pk_Columns_Input = {
  refreshToken: Scalars['uuid'];
};

/** select columns of table "auth.refresh_tokens" */
export enum AuthRefreshTokens_Select_Column {
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  ExpiresAt = 'expiresAt',
  /** column name */
  RefreshToken = 'refreshToken',
  /** column name */
  UserId = 'userId'
}

/** input type for updating data in table "auth.refresh_tokens" */
export type AuthRefreshTokens_Set_Input = {
  createdAt?: InputMaybe<Scalars['timestamptz']>;
  expiresAt?: InputMaybe<Scalars['timestamptz']>;
  refreshToken?: InputMaybe<Scalars['uuid']>;
  userId?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "auth.refresh_tokens" */
export enum AuthRefreshTokens_Update_Column {
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  ExpiresAt = 'expiresAt',
  /** column name */
  RefreshToken = 'refreshToken',
  /** column name */
  UserId = 'userId'
}

/** columns and relationships of "auth.roles" */
export type AuthRoles = {
  __typename?: 'authRoles';
  role: Scalars['String'];
  /** An array relationship */
  userRoles: Array<AuthUserRoles>;
  /** An aggregate relationship */
  userRoles_aggregate: AuthUserRoles_Aggregate;
  /** An array relationship */
  usersByDefaultRole: Array<Users>;
  /** An aggregate relationship */
  usersByDefaultRole_aggregate: Users_Aggregate;
};


/** columns and relationships of "auth.roles" */
export type AuthRolesUserRolesArgs = {
  distinct_on?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserRoles_Order_By>>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};


/** columns and relationships of "auth.roles" */
export type AuthRolesUserRoles_AggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserRoles_Order_By>>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};


/** columns and relationships of "auth.roles" */
export type AuthRolesUsersByDefaultRoleArgs = {
  distinct_on?: InputMaybe<Array<Users_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Users_Order_By>>;
  where?: InputMaybe<Users_Bool_Exp>;
};


/** columns and relationships of "auth.roles" */
export type AuthRolesUsersByDefaultRole_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Users_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Users_Order_By>>;
  where?: InputMaybe<Users_Bool_Exp>;
};

/** aggregated selection of "auth.roles" */
export type AuthRoles_Aggregate = {
  __typename?: 'authRoles_aggregate';
  aggregate?: Maybe<AuthRoles_Aggregate_Fields>;
  nodes: Array<AuthRoles>;
};

/** aggregate fields of "auth.roles" */
export type AuthRoles_Aggregate_Fields = {
  __typename?: 'authRoles_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<AuthRoles_Max_Fields>;
  min?: Maybe<AuthRoles_Min_Fields>;
};


/** aggregate fields of "auth.roles" */
export type AuthRoles_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<AuthRoles_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** Boolean expression to filter rows from the table "auth.roles". All fields are combined with a logical 'AND'. */
export type AuthRoles_Bool_Exp = {
  _and?: InputMaybe<Array<AuthRoles_Bool_Exp>>;
  _not?: InputMaybe<AuthRoles_Bool_Exp>;
  _or?: InputMaybe<Array<AuthRoles_Bool_Exp>>;
  role?: InputMaybe<String_Comparison_Exp>;
  userRoles?: InputMaybe<AuthUserRoles_Bool_Exp>;
  usersByDefaultRole?: InputMaybe<Users_Bool_Exp>;
};

/** unique or primary key constraints on table "auth.roles" */
export enum AuthRoles_Constraint {
  /** unique or primary key constraint */
  RolesPkey = 'roles_pkey'
}

/** input type for inserting data into table "auth.roles" */
export type AuthRoles_Insert_Input = {
  role?: InputMaybe<Scalars['String']>;
  userRoles?: InputMaybe<AuthUserRoles_Arr_Rel_Insert_Input>;
  usersByDefaultRole?: InputMaybe<Users_Arr_Rel_Insert_Input>;
};

/** aggregate max on columns */
export type AuthRoles_Max_Fields = {
  __typename?: 'authRoles_max_fields';
  role?: Maybe<Scalars['String']>;
};

/** aggregate min on columns */
export type AuthRoles_Min_Fields = {
  __typename?: 'authRoles_min_fields';
  role?: Maybe<Scalars['String']>;
};

/** response of any mutation on the table "auth.roles" */
export type AuthRoles_Mutation_Response = {
  __typename?: 'authRoles_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<AuthRoles>;
};

/** input type for inserting object relation for remote table "auth.roles" */
export type AuthRoles_Obj_Rel_Insert_Input = {
  data: AuthRoles_Insert_Input;
  /** on conflict condition */
  on_conflict?: InputMaybe<AuthRoles_On_Conflict>;
};

/** on conflict condition type for table "auth.roles" */
export type AuthRoles_On_Conflict = {
  constraint: AuthRoles_Constraint;
  update_columns?: Array<AuthRoles_Update_Column>;
  where?: InputMaybe<AuthRoles_Bool_Exp>;
};

/** Ordering options when selecting data from "auth.roles". */
export type AuthRoles_Order_By = {
  role?: InputMaybe<Order_By>;
  userRoles_aggregate?: InputMaybe<AuthUserRoles_Aggregate_Order_By>;
  usersByDefaultRole_aggregate?: InputMaybe<Users_Aggregate_Order_By>;
};

/** primary key columns input for table: authRoles */
export type AuthRoles_Pk_Columns_Input = {
  role: Scalars['String'];
};

/** select columns of table "auth.roles" */
export enum AuthRoles_Select_Column {
  /** column name */
  Role = 'role'
}

/** input type for updating data in table "auth.roles" */
export type AuthRoles_Set_Input = {
  role?: InputMaybe<Scalars['String']>;
};

/** update columns of table "auth.roles" */
export enum AuthRoles_Update_Column {
  /** column name */
  Role = 'role'
}

/** columns and relationships of "auth.user_providers" */
export type AuthUserProviders = {
  __typename?: 'authUserProviders';
  accessToken: Scalars['String'];
  createdAt: Scalars['timestamptz'];
  id: Scalars['uuid'];
  /** An object relationship */
  provider: AuthProviders;
  providerId: Scalars['String'];
  providerUserId: Scalars['String'];
  refreshToken?: Maybe<Scalars['String']>;
  updatedAt: Scalars['timestamptz'];
  /** An object relationship */
  user: Users;
  userId: Scalars['uuid'];
};

/** aggregated selection of "auth.user_providers" */
export type AuthUserProviders_Aggregate = {
  __typename?: 'authUserProviders_aggregate';
  aggregate?: Maybe<AuthUserProviders_Aggregate_Fields>;
  nodes: Array<AuthUserProviders>;
};

/** aggregate fields of "auth.user_providers" */
export type AuthUserProviders_Aggregate_Fields = {
  __typename?: 'authUserProviders_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<AuthUserProviders_Max_Fields>;
  min?: Maybe<AuthUserProviders_Min_Fields>;
};


/** aggregate fields of "auth.user_providers" */
export type AuthUserProviders_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "auth.user_providers" */
export type AuthUserProviders_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<AuthUserProviders_Max_Order_By>;
  min?: InputMaybe<AuthUserProviders_Min_Order_By>;
};

/** input type for inserting array relation for remote table "auth.user_providers" */
export type AuthUserProviders_Arr_Rel_Insert_Input = {
  data: Array<AuthUserProviders_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<AuthUserProviders_On_Conflict>;
};

/** Boolean expression to filter rows from the table "auth.user_providers". All fields are combined with a logical 'AND'. */
export type AuthUserProviders_Bool_Exp = {
  _and?: InputMaybe<Array<AuthUserProviders_Bool_Exp>>;
  _not?: InputMaybe<AuthUserProviders_Bool_Exp>;
  _or?: InputMaybe<Array<AuthUserProviders_Bool_Exp>>;
  accessToken?: InputMaybe<String_Comparison_Exp>;
  createdAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  provider?: InputMaybe<AuthProviders_Bool_Exp>;
  providerId?: InputMaybe<String_Comparison_Exp>;
  providerUserId?: InputMaybe<String_Comparison_Exp>;
  refreshToken?: InputMaybe<String_Comparison_Exp>;
  updatedAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  userId?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "auth.user_providers" */
export enum AuthUserProviders_Constraint {
  /** unique or primary key constraint */
  UserProvidersPkey = 'user_providers_pkey',
  /** unique or primary key constraint */
  UserProvidersProviderIdProviderUserIdKey = 'user_providers_provider_id_provider_user_id_key',
  /** unique or primary key constraint */
  UserProvidersUserIdProviderIdKey = 'user_providers_user_id_provider_id_key'
}

/** input type for inserting data into table "auth.user_providers" */
export type AuthUserProviders_Insert_Input = {
  accessToken?: InputMaybe<Scalars['String']>;
  createdAt?: InputMaybe<Scalars['timestamptz']>;
  id?: InputMaybe<Scalars['uuid']>;
  provider?: InputMaybe<AuthProviders_Obj_Rel_Insert_Input>;
  providerId?: InputMaybe<Scalars['String']>;
  providerUserId?: InputMaybe<Scalars['String']>;
  refreshToken?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['timestamptz']>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  userId?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type AuthUserProviders_Max_Fields = {
  __typename?: 'authUserProviders_max_fields';
  accessToken?: Maybe<Scalars['String']>;
  createdAt?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  providerId?: Maybe<Scalars['String']>;
  providerUserId?: Maybe<Scalars['String']>;
  refreshToken?: Maybe<Scalars['String']>;
  updatedAt?: Maybe<Scalars['timestamptz']>;
  userId?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "auth.user_providers" */
export type AuthUserProviders_Max_Order_By = {
  accessToken?: InputMaybe<Order_By>;
  createdAt?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  providerId?: InputMaybe<Order_By>;
  providerUserId?: InputMaybe<Order_By>;
  refreshToken?: InputMaybe<Order_By>;
  updatedAt?: InputMaybe<Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type AuthUserProviders_Min_Fields = {
  __typename?: 'authUserProviders_min_fields';
  accessToken?: Maybe<Scalars['String']>;
  createdAt?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  providerId?: Maybe<Scalars['String']>;
  providerUserId?: Maybe<Scalars['String']>;
  refreshToken?: Maybe<Scalars['String']>;
  updatedAt?: Maybe<Scalars['timestamptz']>;
  userId?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "auth.user_providers" */
export type AuthUserProviders_Min_Order_By = {
  accessToken?: InputMaybe<Order_By>;
  createdAt?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  providerId?: InputMaybe<Order_By>;
  providerUserId?: InputMaybe<Order_By>;
  refreshToken?: InputMaybe<Order_By>;
  updatedAt?: InputMaybe<Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "auth.user_providers" */
export type AuthUserProviders_Mutation_Response = {
  __typename?: 'authUserProviders_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<AuthUserProviders>;
};

/** on conflict condition type for table "auth.user_providers" */
export type AuthUserProviders_On_Conflict = {
  constraint: AuthUserProviders_Constraint;
  update_columns?: Array<AuthUserProviders_Update_Column>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};

/** Ordering options when selecting data from "auth.user_providers". */
export type AuthUserProviders_Order_By = {
  accessToken?: InputMaybe<Order_By>;
  createdAt?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  provider?: InputMaybe<AuthProviders_Order_By>;
  providerId?: InputMaybe<Order_By>;
  providerUserId?: InputMaybe<Order_By>;
  refreshToken?: InputMaybe<Order_By>;
  updatedAt?: InputMaybe<Order_By>;
  user?: InputMaybe<Users_Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** primary key columns input for table: authUserProviders */
export type AuthUserProviders_Pk_Columns_Input = {
  id: Scalars['uuid'];
};

/** select columns of table "auth.user_providers" */
export enum AuthUserProviders_Select_Column {
  /** column name */
  AccessToken = 'accessToken',
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  Id = 'id',
  /** column name */
  ProviderId = 'providerId',
  /** column name */
  ProviderUserId = 'providerUserId',
  /** column name */
  RefreshToken = 'refreshToken',
  /** column name */
  UpdatedAt = 'updatedAt',
  /** column name */
  UserId = 'userId'
}

/** input type for updating data in table "auth.user_providers" */
export type AuthUserProviders_Set_Input = {
  accessToken?: InputMaybe<Scalars['String']>;
  createdAt?: InputMaybe<Scalars['timestamptz']>;
  id?: InputMaybe<Scalars['uuid']>;
  providerId?: InputMaybe<Scalars['String']>;
  providerUserId?: InputMaybe<Scalars['String']>;
  refreshToken?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['timestamptz']>;
  userId?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "auth.user_providers" */
export enum AuthUserProviders_Update_Column {
  /** column name */
  AccessToken = 'accessToken',
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  Id = 'id',
  /** column name */
  ProviderId = 'providerId',
  /** column name */
  ProviderUserId = 'providerUserId',
  /** column name */
  RefreshToken = 'refreshToken',
  /** column name */
  UpdatedAt = 'updatedAt',
  /** column name */
  UserId = 'userId'
}

/** columns and relationships of "auth.user_roles" */
export type AuthUserRoles = {
  __typename?: 'authUserRoles';
  createdAt: Scalars['timestamptz'];
  id: Scalars['uuid'];
  role: Scalars['String'];
  /** An object relationship */
  roleByRole: AuthRoles;
  /** An object relationship */
  user: Users;
  userId: Scalars['uuid'];
};

/** aggregated selection of "auth.user_roles" */
export type AuthUserRoles_Aggregate = {
  __typename?: 'authUserRoles_aggregate';
  aggregate?: Maybe<AuthUserRoles_Aggregate_Fields>;
  nodes: Array<AuthUserRoles>;
};

/** aggregate fields of "auth.user_roles" */
export type AuthUserRoles_Aggregate_Fields = {
  __typename?: 'authUserRoles_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<AuthUserRoles_Max_Fields>;
  min?: Maybe<AuthUserRoles_Min_Fields>;
};


/** aggregate fields of "auth.user_roles" */
export type AuthUserRoles_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "auth.user_roles" */
export type AuthUserRoles_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<AuthUserRoles_Max_Order_By>;
  min?: InputMaybe<AuthUserRoles_Min_Order_By>;
};

/** input type for inserting array relation for remote table "auth.user_roles" */
export type AuthUserRoles_Arr_Rel_Insert_Input = {
  data: Array<AuthUserRoles_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<AuthUserRoles_On_Conflict>;
};

/** Boolean expression to filter rows from the table "auth.user_roles". All fields are combined with a logical 'AND'. */
export type AuthUserRoles_Bool_Exp = {
  _and?: InputMaybe<Array<AuthUserRoles_Bool_Exp>>;
  _not?: InputMaybe<AuthUserRoles_Bool_Exp>;
  _or?: InputMaybe<Array<AuthUserRoles_Bool_Exp>>;
  createdAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  role?: InputMaybe<String_Comparison_Exp>;
  roleByRole?: InputMaybe<AuthRoles_Bool_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  userId?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "auth.user_roles" */
export enum AuthUserRoles_Constraint {
  /** unique or primary key constraint */
  UserRolesPkey = 'user_roles_pkey',
  /** unique or primary key constraint */
  UserRolesUserIdRoleKey = 'user_roles_user_id_role_key'
}

/** input type for inserting data into table "auth.user_roles" */
export type AuthUserRoles_Insert_Input = {
  createdAt?: InputMaybe<Scalars['timestamptz']>;
  id?: InputMaybe<Scalars['uuid']>;
  role?: InputMaybe<Scalars['String']>;
  roleByRole?: InputMaybe<AuthRoles_Obj_Rel_Insert_Input>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  userId?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type AuthUserRoles_Max_Fields = {
  __typename?: 'authUserRoles_max_fields';
  createdAt?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  role?: Maybe<Scalars['String']>;
  userId?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "auth.user_roles" */
export type AuthUserRoles_Max_Order_By = {
  createdAt?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  role?: InputMaybe<Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type AuthUserRoles_Min_Fields = {
  __typename?: 'authUserRoles_min_fields';
  createdAt?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  role?: Maybe<Scalars['String']>;
  userId?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "auth.user_roles" */
export type AuthUserRoles_Min_Order_By = {
  createdAt?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  role?: InputMaybe<Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "auth.user_roles" */
export type AuthUserRoles_Mutation_Response = {
  __typename?: 'authUserRoles_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<AuthUserRoles>;
};

/** on conflict condition type for table "auth.user_roles" */
export type AuthUserRoles_On_Conflict = {
  constraint: AuthUserRoles_Constraint;
  update_columns?: Array<AuthUserRoles_Update_Column>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};

/** Ordering options when selecting data from "auth.user_roles". */
export type AuthUserRoles_Order_By = {
  createdAt?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  role?: InputMaybe<Order_By>;
  roleByRole?: InputMaybe<AuthRoles_Order_By>;
  user?: InputMaybe<Users_Order_By>;
  userId?: InputMaybe<Order_By>;
};

/** primary key columns input for table: authUserRoles */
export type AuthUserRoles_Pk_Columns_Input = {
  id: Scalars['uuid'];
};

/** select columns of table "auth.user_roles" */
export enum AuthUserRoles_Select_Column {
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  Id = 'id',
  /** column name */
  Role = 'role',
  /** column name */
  UserId = 'userId'
}

/** input type for updating data in table "auth.user_roles" */
export type AuthUserRoles_Set_Input = {
  createdAt?: InputMaybe<Scalars['timestamptz']>;
  id?: InputMaybe<Scalars['uuid']>;
  role?: InputMaybe<Scalars['String']>;
  userId?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "auth.user_roles" */
export enum AuthUserRoles_Update_Column {
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  Id = 'id',
  /** column name */
  Role = 'role',
  /** column name */
  UserId = 'userId'
}

/** columns and relationships of "auth.user_private" */
export type Auth_User_Private = {
  __typename?: 'auth_user_private';
  birthdate?: Maybe<Scalars['date']>;
  email?: Maybe<Scalars['citext']>;
  gender?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['uuid']>;
  is_username_set?: Maybe<Scalars['Boolean']>;
  phone_number?: Maybe<Scalars['String']>;
};

/** aggregated selection of "auth.user_private" */
export type Auth_User_Private_Aggregate = {
  __typename?: 'auth_user_private_aggregate';
  aggregate?: Maybe<Auth_User_Private_Aggregate_Fields>;
  nodes: Array<Auth_User_Private>;
};

/** aggregate fields of "auth.user_private" */
export type Auth_User_Private_Aggregate_Fields = {
  __typename?: 'auth_user_private_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Auth_User_Private_Max_Fields>;
  min?: Maybe<Auth_User_Private_Min_Fields>;
};


/** aggregate fields of "auth.user_private" */
export type Auth_User_Private_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Auth_User_Private_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** Boolean expression to filter rows from the table "auth.user_private". All fields are combined with a logical 'AND'. */
export type Auth_User_Private_Bool_Exp = {
  _and?: InputMaybe<Array<Auth_User_Private_Bool_Exp>>;
  _not?: InputMaybe<Auth_User_Private_Bool_Exp>;
  _or?: InputMaybe<Array<Auth_User_Private_Bool_Exp>>;
  birthdate?: InputMaybe<Date_Comparison_Exp>;
  email?: InputMaybe<Citext_Comparison_Exp>;
  gender?: InputMaybe<String_Comparison_Exp>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  is_username_set?: InputMaybe<Boolean_Comparison_Exp>;
  phone_number?: InputMaybe<String_Comparison_Exp>;
};

/** input type for inserting data into table "auth.user_private" */
export type Auth_User_Private_Insert_Input = {
  birthdate?: InputMaybe<Scalars['date']>;
  email?: InputMaybe<Scalars['citext']>;
  gender?: InputMaybe<Scalars['String']>;
  id?: InputMaybe<Scalars['uuid']>;
  is_username_set?: InputMaybe<Scalars['Boolean']>;
  phone_number?: InputMaybe<Scalars['String']>;
};

/** aggregate max on columns */
export type Auth_User_Private_Max_Fields = {
  __typename?: 'auth_user_private_max_fields';
  birthdate?: Maybe<Scalars['date']>;
  email?: Maybe<Scalars['citext']>;
  gender?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['uuid']>;
  phone_number?: Maybe<Scalars['String']>;
};

/** aggregate min on columns */
export type Auth_User_Private_Min_Fields = {
  __typename?: 'auth_user_private_min_fields';
  birthdate?: Maybe<Scalars['date']>;
  email?: Maybe<Scalars['citext']>;
  gender?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['uuid']>;
  phone_number?: Maybe<Scalars['String']>;
};

/** response of any mutation on the table "auth.user_private" */
export type Auth_User_Private_Mutation_Response = {
  __typename?: 'auth_user_private_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Auth_User_Private>;
};

/** input type for inserting object relation for remote table "auth.user_private" */
export type Auth_User_Private_Obj_Rel_Insert_Input = {
  data: Auth_User_Private_Insert_Input;
};

/** Ordering options when selecting data from "auth.user_private". */
export type Auth_User_Private_Order_By = {
  birthdate?: InputMaybe<Order_By>;
  email?: InputMaybe<Order_By>;
  gender?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  is_username_set?: InputMaybe<Order_By>;
  phone_number?: InputMaybe<Order_By>;
};

/** select columns of table "auth.user_private" */
export enum Auth_User_Private_Select_Column {
  /** column name */
  Birthdate = 'birthdate',
  /** column name */
  Email = 'email',
  /** column name */
  Gender = 'gender',
  /** column name */
  Id = 'id',
  /** column name */
  IsUsernameSet = 'is_username_set',
  /** column name */
  PhoneNumber = 'phone_number'
}

/** input type for updating data in table "auth.user_private" */
export type Auth_User_Private_Set_Input = {
  birthdate?: InputMaybe<Scalars['date']>;
  email?: InputMaybe<Scalars['citext']>;
  gender?: InputMaybe<Scalars['String']>;
  id?: InputMaybe<Scalars['uuid']>;
  is_username_set?: InputMaybe<Scalars['Boolean']>;
  phone_number?: InputMaybe<Scalars['String']>;
};

/** columns and relationships of "block" */
export type Block = {
  __typename?: 'block';
  /** An array relationship */
  block_children: Array<Block>;
  /** An aggregate relationship */
  block_children_aggregate: Block_Aggregate;
  /** An array relationship */
  block_followers: Array<Block_Follower>;
  /** An aggregate relationship */
  block_followers_aggregate: Block_Follower_Aggregate;
  /** An array relationship */
  block_likes: Array<Block_Like>;
  /** An aggregate relationship */
  block_likes_aggregate: Block_Like_Aggregate;
  /** An object relationship */
  block_parent?: Maybe<Block>;
  /** Block parent id for hierarchy. If null, the block will be at the root */
  block_parent_id?: Maybe<Scalars['uuid']>;
  /** An array relationship */
  block_referees: Array<Block_Reference>;
  /** An aggregate relationship */
  block_referees_aggregate: Block_Reference_Aggregate;
  /** An array relationship */
  block_referrers: Array<Block_Reference>;
  /** An aggregate relationship */
  block_referrers_aggregate: Block_Reference_Aggregate;
  /** An array relationship */
  block_reposts: Array<Block_Repost>;
  /** An aggregate relationship */
  block_reposts_aggregate: Block_Repost_Aggregate;
  /** Slate children. If null, empty text */
  children?: Maybe<Scalars['jsonb']>;
  created_at: Scalars['timestamptz'];
  /** Slate properties */
  data?: Maybe<Scalars['jsonb']>;
  /** An object relationship */
  fork_parent?: Maybe<Block>;
  /** Forked from */
  fork_parent_id?: Maybe<Scalars['uuid']>;
  /** An array relationship */
  forks: Array<Block>;
  /** An aggregate relationship */
  forks_aggregate: Block_Aggregate;
  id: Scalars['uuid'];
  is_posted?: Maybe<Scalars['Boolean']>;
  /** Lexicographical order of the block inside the parent block */
  order?: Maybe<Scalars['String']>;
  path?: Maybe<Scalars['ltree']>;
  /** Full text without marks */
  text?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
  /** Type of the block. If null, it's a paragraph */
  type?: Maybe<Scalars['String']>;
  updated_at: Scalars['timestamptz'];
  /** An object relationship */
  user: Users;
  /** Author id */
  user_id: Scalars['uuid'];
};


/** columns and relationships of "block" */
export type BlockBlock_ChildrenArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_Children_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_FollowersArgs = {
  distinct_on?: InputMaybe<Array<Block_Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Follower_Order_By>>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_Followers_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Follower_Order_By>>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_LikesArgs = {
  distinct_on?: InputMaybe<Array<Block_Like_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Like_Order_By>>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_Likes_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Like_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Like_Order_By>>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_RefereesArgs = {
  distinct_on?: InputMaybe<Array<Block_Reference_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Reference_Order_By>>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_Referees_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Reference_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Reference_Order_By>>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_ReferrersArgs = {
  distinct_on?: InputMaybe<Array<Block_Reference_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Reference_Order_By>>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_Referrers_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Reference_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Reference_Order_By>>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_RepostsArgs = {
  distinct_on?: InputMaybe<Array<Block_Repost_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Repost_Order_By>>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockBlock_Reposts_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Repost_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Repost_Order_By>>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockChildrenArgs = {
  path?: InputMaybe<Scalars['String']>;
};


/** columns and relationships of "block" */
export type BlockDataArgs = {
  path?: InputMaybe<Scalars['String']>;
};


/** columns and relationships of "block" */
export type BlockForksArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


/** columns and relationships of "block" */
export type BlockForks_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};

/** aggregated selection of "block" */
export type Block_Aggregate = {
  __typename?: 'block_aggregate';
  aggregate?: Maybe<Block_Aggregate_Fields>;
  nodes: Array<Block>;
};

/** aggregate fields of "block" */
export type Block_Aggregate_Fields = {
  __typename?: 'block_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Block_Max_Fields>;
  min?: Maybe<Block_Min_Fields>;
};


/** aggregate fields of "block" */
export type Block_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Block_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "block" */
export type Block_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<Block_Max_Order_By>;
  min?: InputMaybe<Block_Min_Order_By>;
};

/** append existing jsonb value of filtered columns with new jsonb value */
export type Block_Append_Input = {
  /** Slate children. If null, empty text */
  children?: InputMaybe<Scalars['jsonb']>;
  /** Slate properties */
  data?: InputMaybe<Scalars['jsonb']>;
};

/** input type for inserting array relation for remote table "block" */
export type Block_Arr_Rel_Insert_Input = {
  data: Array<Block_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<Block_On_Conflict>;
};

/** Boolean expression to filter rows from the table "block". All fields are combined with a logical 'AND'. */
export type Block_Bool_Exp = {
  _and?: InputMaybe<Array<Block_Bool_Exp>>;
  _not?: InputMaybe<Block_Bool_Exp>;
  _or?: InputMaybe<Array<Block_Bool_Exp>>;
  block_children?: InputMaybe<Block_Bool_Exp>;
  block_followers?: InputMaybe<Block_Follower_Bool_Exp>;
  block_likes?: InputMaybe<Block_Like_Bool_Exp>;
  block_parent?: InputMaybe<Block_Bool_Exp>;
  block_parent_id?: InputMaybe<Uuid_Comparison_Exp>;
  block_referees?: InputMaybe<Block_Reference_Bool_Exp>;
  block_referrers?: InputMaybe<Block_Reference_Bool_Exp>;
  block_reposts?: InputMaybe<Block_Repost_Bool_Exp>;
  children?: InputMaybe<Jsonb_Comparison_Exp>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  data?: InputMaybe<Jsonb_Comparison_Exp>;
  fork_parent?: InputMaybe<Block_Bool_Exp>;
  fork_parent_id?: InputMaybe<Uuid_Comparison_Exp>;
  forks?: InputMaybe<Block_Bool_Exp>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  is_posted?: InputMaybe<Boolean_Comparison_Exp>;
  order?: InputMaybe<String_Comparison_Exp>;
  path?: InputMaybe<Ltree_Comparison_Exp>;
  text?: InputMaybe<String_Comparison_Exp>;
  title?: InputMaybe<String_Comparison_Exp>;
  type?: InputMaybe<String_Comparison_Exp>;
  updated_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  user_id?: InputMaybe<Uuid_Comparison_Exp>;
};

/** columns and relationships of "block_comment" */
export type Block_Comment = {
  __typename?: 'block_comment';
  block_id: Scalars['uuid'];
  created_at: Scalars['timestamptz'];
  id: Scalars['uuid'];
  text: Scalars['String'];
  /** An object relationship */
  user: Users;
  user_id: Scalars['uuid'];
};

/** aggregated selection of "block_comment" */
export type Block_Comment_Aggregate = {
  __typename?: 'block_comment_aggregate';
  aggregate?: Maybe<Block_Comment_Aggregate_Fields>;
  nodes: Array<Block_Comment>;
};

/** aggregate fields of "block_comment" */
export type Block_Comment_Aggregate_Fields = {
  __typename?: 'block_comment_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Block_Comment_Max_Fields>;
  min?: Maybe<Block_Comment_Min_Fields>;
};


/** aggregate fields of "block_comment" */
export type Block_Comment_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Block_Comment_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "block_comment" */
export type Block_Comment_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<Block_Comment_Max_Order_By>;
  min?: InputMaybe<Block_Comment_Min_Order_By>;
};

/** input type for inserting array relation for remote table "block_comment" */
export type Block_Comment_Arr_Rel_Insert_Input = {
  data: Array<Block_Comment_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<Block_Comment_On_Conflict>;
};

/** Boolean expression to filter rows from the table "block_comment". All fields are combined with a logical 'AND'. */
export type Block_Comment_Bool_Exp = {
  _and?: InputMaybe<Array<Block_Comment_Bool_Exp>>;
  _not?: InputMaybe<Block_Comment_Bool_Exp>;
  _or?: InputMaybe<Array<Block_Comment_Bool_Exp>>;
  block_id?: InputMaybe<Uuid_Comparison_Exp>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  text?: InputMaybe<String_Comparison_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  user_id?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "block_comment" */
export enum Block_Comment_Constraint {
  /** unique or primary key constraint */
  BlockCommentPkey = 'block_comment_pkey'
}

/** input type for inserting data into table "block_comment" */
export type Block_Comment_Insert_Input = {
  block_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  id?: InputMaybe<Scalars['uuid']>;
  text?: InputMaybe<Scalars['String']>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type Block_Comment_Max_Fields = {
  __typename?: 'block_comment_max_fields';
  block_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  text?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "block_comment" */
export type Block_Comment_Max_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  text?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type Block_Comment_Min_Fields = {
  __typename?: 'block_comment_min_fields';
  block_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  text?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "block_comment" */
export type Block_Comment_Min_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  text?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "block_comment" */
export type Block_Comment_Mutation_Response = {
  __typename?: 'block_comment_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Block_Comment>;
};

/** on conflict condition type for table "block_comment" */
export type Block_Comment_On_Conflict = {
  constraint: Block_Comment_Constraint;
  update_columns?: Array<Block_Comment_Update_Column>;
  where?: InputMaybe<Block_Comment_Bool_Exp>;
};

/** Ordering options when selecting data from "block_comment". */
export type Block_Comment_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  text?: InputMaybe<Order_By>;
  user?: InputMaybe<Users_Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** primary key columns input for table: block_comment */
export type Block_Comment_Pk_Columns_Input = {
  id: Scalars['uuid'];
};

/** select columns of table "block_comment" */
export enum Block_Comment_Select_Column {
  /** column name */
  BlockId = 'block_id',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  Id = 'id',
  /** column name */
  Text = 'text',
  /** column name */
  UserId = 'user_id'
}

/** input type for updating data in table "block_comment" */
export type Block_Comment_Set_Input = {
  block_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  id?: InputMaybe<Scalars['uuid']>;
  text?: InputMaybe<Scalars['String']>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "block_comment" */
export enum Block_Comment_Update_Column {
  /** column name */
  BlockId = 'block_id',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  Id = 'id',
  /** column name */
  Text = 'text',
  /** column name */
  UserId = 'user_id'
}

/** unique or primary key constraints on table "block" */
export enum Block_Constraint {
  /** unique or primary key constraint */
  BlockPkey = 'block_pkey'
}

/** delete the field or element with specified path (for JSON arrays, negative integers count from the end) */
export type Block_Delete_At_Path_Input = {
  /** Slate children. If null, empty text */
  children?: InputMaybe<Array<Scalars['String']>>;
  /** Slate properties */
  data?: InputMaybe<Array<Scalars['String']>>;
};

/** delete the array element with specified index (negative integers count from the end). throws an error if top level container is not an array */
export type Block_Delete_Elem_Input = {
  /** Slate children. If null, empty text */
  children?: InputMaybe<Scalars['Int']>;
  /** Slate properties */
  data?: InputMaybe<Scalars['Int']>;
};

/** delete key/value pair or string element. key/value pairs are matched based on their key value */
export type Block_Delete_Key_Input = {
  /** Slate children. If null, empty text */
  children?: InputMaybe<Scalars['String']>;
  /** Slate properties */
  data?: InputMaybe<Scalars['String']>;
};

/** columns and relationships of "block_follower" */
export type Block_Follower = {
  __typename?: 'block_follower';
  /** An object relationship */
  block: Block;
  block_id: Scalars['uuid'];
  created_at: Scalars['timestamptz'];
  /** An object relationship */
  user: Users;
  user_id: Scalars['uuid'];
};

/** aggregated selection of "block_follower" */
export type Block_Follower_Aggregate = {
  __typename?: 'block_follower_aggregate';
  aggregate?: Maybe<Block_Follower_Aggregate_Fields>;
  nodes: Array<Block_Follower>;
};

/** aggregate fields of "block_follower" */
export type Block_Follower_Aggregate_Fields = {
  __typename?: 'block_follower_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Block_Follower_Max_Fields>;
  min?: Maybe<Block_Follower_Min_Fields>;
};


/** aggregate fields of "block_follower" */
export type Block_Follower_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Block_Follower_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "block_follower" */
export type Block_Follower_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<Block_Follower_Max_Order_By>;
  min?: InputMaybe<Block_Follower_Min_Order_By>;
};

/** input type for inserting array relation for remote table "block_follower" */
export type Block_Follower_Arr_Rel_Insert_Input = {
  data: Array<Block_Follower_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<Block_Follower_On_Conflict>;
};

/** Boolean expression to filter rows from the table "block_follower". All fields are combined with a logical 'AND'. */
export type Block_Follower_Bool_Exp = {
  _and?: InputMaybe<Array<Block_Follower_Bool_Exp>>;
  _not?: InputMaybe<Block_Follower_Bool_Exp>;
  _or?: InputMaybe<Array<Block_Follower_Bool_Exp>>;
  block?: InputMaybe<Block_Bool_Exp>;
  block_id?: InputMaybe<Uuid_Comparison_Exp>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  user_id?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "block_follower" */
export enum Block_Follower_Constraint {
  /** unique or primary key constraint */
  BlockFollowerPkey = 'block_follower_pkey'
}

/** input type for inserting data into table "block_follower" */
export type Block_Follower_Insert_Input = {
  block?: InputMaybe<Block_Obj_Rel_Insert_Input>;
  block_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type Block_Follower_Max_Fields = {
  __typename?: 'block_follower_max_fields';
  block_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "block_follower" */
export type Block_Follower_Max_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type Block_Follower_Min_Fields = {
  __typename?: 'block_follower_min_fields';
  block_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "block_follower" */
export type Block_Follower_Min_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "block_follower" */
export type Block_Follower_Mutation_Response = {
  __typename?: 'block_follower_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Block_Follower>;
};

/** on conflict condition type for table "block_follower" */
export type Block_Follower_On_Conflict = {
  constraint: Block_Follower_Constraint;
  update_columns?: Array<Block_Follower_Update_Column>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};

/** Ordering options when selecting data from "block_follower". */
export type Block_Follower_Order_By = {
  block?: InputMaybe<Block_Order_By>;
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user?: InputMaybe<Users_Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** primary key columns input for table: block_follower */
export type Block_Follower_Pk_Columns_Input = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};

/** select columns of table "block_follower" */
export enum Block_Follower_Select_Column {
  /** column name */
  BlockId = 'block_id',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  UserId = 'user_id'
}

/** input type for updating data in table "block_follower" */
export type Block_Follower_Set_Input = {
  block_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "block_follower" */
export enum Block_Follower_Update_Column {
  /** column name */
  BlockId = 'block_id',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  UserId = 'user_id'
}

/** input type for inserting data into table "block" */
export type Block_Insert_Input = {
  block_children?: InputMaybe<Block_Arr_Rel_Insert_Input>;
  block_followers?: InputMaybe<Block_Follower_Arr_Rel_Insert_Input>;
  block_likes?: InputMaybe<Block_Like_Arr_Rel_Insert_Input>;
  block_parent?: InputMaybe<Block_Obj_Rel_Insert_Input>;
  /** Block parent id for hierarchy. If null, the block will be at the root */
  block_parent_id?: InputMaybe<Scalars['uuid']>;
  block_referees?: InputMaybe<Block_Reference_Arr_Rel_Insert_Input>;
  block_referrers?: InputMaybe<Block_Reference_Arr_Rel_Insert_Input>;
  block_reposts?: InputMaybe<Block_Repost_Arr_Rel_Insert_Input>;
  /** Slate children. If null, empty text */
  children?: InputMaybe<Scalars['jsonb']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  /** Slate properties */
  data?: InputMaybe<Scalars['jsonb']>;
  fork_parent?: InputMaybe<Block_Obj_Rel_Insert_Input>;
  /** Forked from */
  fork_parent_id?: InputMaybe<Scalars['uuid']>;
  forks?: InputMaybe<Block_Arr_Rel_Insert_Input>;
  id?: InputMaybe<Scalars['uuid']>;
  is_posted?: InputMaybe<Scalars['Boolean']>;
  /** Lexicographical order of the block inside the parent block */
  order?: InputMaybe<Scalars['String']>;
  path?: InputMaybe<Scalars['ltree']>;
  /** Full text without marks */
  text?: InputMaybe<Scalars['String']>;
  title?: InputMaybe<Scalars['String']>;
  /** Type of the block. If null, it's a paragraph */
  type?: InputMaybe<Scalars['String']>;
  updated_at?: InputMaybe<Scalars['timestamptz']>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  /** Author id */
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** columns and relationships of "block_like" */
export type Block_Like = {
  __typename?: 'block_like';
  /** An object relationship */
  block: Block;
  block_id: Scalars['uuid'];
  created_at: Scalars['timestamptz'];
  /** An object relationship */
  user: Users;
  user_id: Scalars['uuid'];
};

/** aggregated selection of "block_like" */
export type Block_Like_Aggregate = {
  __typename?: 'block_like_aggregate';
  aggregate?: Maybe<Block_Like_Aggregate_Fields>;
  nodes: Array<Block_Like>;
};

/** aggregate fields of "block_like" */
export type Block_Like_Aggregate_Fields = {
  __typename?: 'block_like_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Block_Like_Max_Fields>;
  min?: Maybe<Block_Like_Min_Fields>;
};


/** aggregate fields of "block_like" */
export type Block_Like_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Block_Like_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "block_like" */
export type Block_Like_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<Block_Like_Max_Order_By>;
  min?: InputMaybe<Block_Like_Min_Order_By>;
};

/** input type for inserting array relation for remote table "block_like" */
export type Block_Like_Arr_Rel_Insert_Input = {
  data: Array<Block_Like_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<Block_Like_On_Conflict>;
};

/** Boolean expression to filter rows from the table "block_like". All fields are combined with a logical 'AND'. */
export type Block_Like_Bool_Exp = {
  _and?: InputMaybe<Array<Block_Like_Bool_Exp>>;
  _not?: InputMaybe<Block_Like_Bool_Exp>;
  _or?: InputMaybe<Array<Block_Like_Bool_Exp>>;
  block?: InputMaybe<Block_Bool_Exp>;
  block_id?: InputMaybe<Uuid_Comparison_Exp>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  user_id?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "block_like" */
export enum Block_Like_Constraint {
  /** unique or primary key constraint */
  BlockLikePkey = 'block_like_pkey'
}

/** input type for inserting data into table "block_like" */
export type Block_Like_Insert_Input = {
  block?: InputMaybe<Block_Obj_Rel_Insert_Input>;
  block_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type Block_Like_Max_Fields = {
  __typename?: 'block_like_max_fields';
  block_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "block_like" */
export type Block_Like_Max_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type Block_Like_Min_Fields = {
  __typename?: 'block_like_min_fields';
  block_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "block_like" */
export type Block_Like_Min_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "block_like" */
export type Block_Like_Mutation_Response = {
  __typename?: 'block_like_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Block_Like>;
};

/** on conflict condition type for table "block_like" */
export type Block_Like_On_Conflict = {
  constraint: Block_Like_Constraint;
  update_columns?: Array<Block_Like_Update_Column>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};

/** Ordering options when selecting data from "block_like". */
export type Block_Like_Order_By = {
  block?: InputMaybe<Block_Order_By>;
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user?: InputMaybe<Users_Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** primary key columns input for table: block_like */
export type Block_Like_Pk_Columns_Input = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};

/** select columns of table "block_like" */
export enum Block_Like_Select_Column {
  /** column name */
  BlockId = 'block_id',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  UserId = 'user_id'
}

/** input type for updating data in table "block_like" */
export type Block_Like_Set_Input = {
  block_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "block_like" */
export enum Block_Like_Update_Column {
  /** column name */
  BlockId = 'block_id',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  UserId = 'user_id'
}

/** aggregate max on columns */
export type Block_Max_Fields = {
  __typename?: 'block_max_fields';
  /** Block parent id for hierarchy. If null, the block will be at the root */
  block_parent_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  /** Forked from */
  fork_parent_id?: Maybe<Scalars['uuid']>;
  id?: Maybe<Scalars['uuid']>;
  /** Lexicographical order of the block inside the parent block */
  order?: Maybe<Scalars['String']>;
  path?: Maybe<Scalars['ltree']>;
  /** Full text without marks */
  text?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
  /** Type of the block. If null, it's a paragraph */
  type?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamptz']>;
  /** Author id */
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "block" */
export type Block_Max_Order_By = {
  /** Block parent id for hierarchy. If null, the block will be at the root */
  block_parent_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  /** Forked from */
  fork_parent_id?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  /** Lexicographical order of the block inside the parent block */
  order?: InputMaybe<Order_By>;
  path?: InputMaybe<Order_By>;
  /** Full text without marks */
  text?: InputMaybe<Order_By>;
  title?: InputMaybe<Order_By>;
  /** Type of the block. If null, it's a paragraph */
  type?: InputMaybe<Order_By>;
  updated_at?: InputMaybe<Order_By>;
  /** Author id */
  user_id?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type Block_Min_Fields = {
  __typename?: 'block_min_fields';
  /** Block parent id for hierarchy. If null, the block will be at the root */
  block_parent_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  /** Forked from */
  fork_parent_id?: Maybe<Scalars['uuid']>;
  id?: Maybe<Scalars['uuid']>;
  /** Lexicographical order of the block inside the parent block */
  order?: Maybe<Scalars['String']>;
  path?: Maybe<Scalars['ltree']>;
  /** Full text without marks */
  text?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
  /** Type of the block. If null, it's a paragraph */
  type?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamptz']>;
  /** Author id */
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "block" */
export type Block_Min_Order_By = {
  /** Block parent id for hierarchy. If null, the block will be at the root */
  block_parent_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  /** Forked from */
  fork_parent_id?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  /** Lexicographical order of the block inside the parent block */
  order?: InputMaybe<Order_By>;
  path?: InputMaybe<Order_By>;
  /** Full text without marks */
  text?: InputMaybe<Order_By>;
  title?: InputMaybe<Order_By>;
  /** Type of the block. If null, it's a paragraph */
  type?: InputMaybe<Order_By>;
  updated_at?: InputMaybe<Order_By>;
  /** Author id */
  user_id?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "block" */
export type Block_Mutation_Response = {
  __typename?: 'block_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Block>;
};

/** input type for inserting object relation for remote table "block" */
export type Block_Obj_Rel_Insert_Input = {
  data: Block_Insert_Input;
  /** on conflict condition */
  on_conflict?: InputMaybe<Block_On_Conflict>;
};

/** on conflict condition type for table "block" */
export type Block_On_Conflict = {
  constraint: Block_Constraint;
  update_columns?: Array<Block_Update_Column>;
  where?: InputMaybe<Block_Bool_Exp>;
};

/** Ordering options when selecting data from "block". */
export type Block_Order_By = {
  block_children_aggregate?: InputMaybe<Block_Aggregate_Order_By>;
  block_followers_aggregate?: InputMaybe<Block_Follower_Aggregate_Order_By>;
  block_likes_aggregate?: InputMaybe<Block_Like_Aggregate_Order_By>;
  block_parent?: InputMaybe<Block_Order_By>;
  block_parent_id?: InputMaybe<Order_By>;
  block_referees_aggregate?: InputMaybe<Block_Reference_Aggregate_Order_By>;
  block_referrers_aggregate?: InputMaybe<Block_Reference_Aggregate_Order_By>;
  block_reposts_aggregate?: InputMaybe<Block_Repost_Aggregate_Order_By>;
  children?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  data?: InputMaybe<Order_By>;
  fork_parent?: InputMaybe<Block_Order_By>;
  fork_parent_id?: InputMaybe<Order_By>;
  forks_aggregate?: InputMaybe<Block_Aggregate_Order_By>;
  id?: InputMaybe<Order_By>;
  is_posted?: InputMaybe<Order_By>;
  order?: InputMaybe<Order_By>;
  path?: InputMaybe<Order_By>;
  text?: InputMaybe<Order_By>;
  title?: InputMaybe<Order_By>;
  type?: InputMaybe<Order_By>;
  updated_at?: InputMaybe<Order_By>;
  user?: InputMaybe<Users_Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** primary key columns input for table: block */
export type Block_Pk_Columns_Input = {
  id: Scalars['uuid'];
};

/** prepend existing jsonb value of filtered columns with new jsonb value */
export type Block_Prepend_Input = {
  /** Slate children. If null, empty text */
  children?: InputMaybe<Scalars['jsonb']>;
  /** Slate properties */
  data?: InputMaybe<Scalars['jsonb']>;
};

/** columns and relationships of "block_reference" */
export type Block_Reference = {
  __typename?: 'block_reference';
  /** An object relationship */
  block_referee: Block;
  block_referee_id: Scalars['uuid'];
  /** An object relationship */
  block_referrer: Block;
  block_referrer_id: Scalars['uuid'];
  created_at: Scalars['timestamptz'];
};

/** aggregated selection of "block_reference" */
export type Block_Reference_Aggregate = {
  __typename?: 'block_reference_aggregate';
  aggregate?: Maybe<Block_Reference_Aggregate_Fields>;
  nodes: Array<Block_Reference>;
};

/** aggregate fields of "block_reference" */
export type Block_Reference_Aggregate_Fields = {
  __typename?: 'block_reference_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Block_Reference_Max_Fields>;
  min?: Maybe<Block_Reference_Min_Fields>;
};


/** aggregate fields of "block_reference" */
export type Block_Reference_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Block_Reference_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "block_reference" */
export type Block_Reference_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<Block_Reference_Max_Order_By>;
  min?: InputMaybe<Block_Reference_Min_Order_By>;
};

/** input type for inserting array relation for remote table "block_reference" */
export type Block_Reference_Arr_Rel_Insert_Input = {
  data: Array<Block_Reference_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<Block_Reference_On_Conflict>;
};

/** Boolean expression to filter rows from the table "block_reference". All fields are combined with a logical 'AND'. */
export type Block_Reference_Bool_Exp = {
  _and?: InputMaybe<Array<Block_Reference_Bool_Exp>>;
  _not?: InputMaybe<Block_Reference_Bool_Exp>;
  _or?: InputMaybe<Array<Block_Reference_Bool_Exp>>;
  block_referee?: InputMaybe<Block_Bool_Exp>;
  block_referee_id?: InputMaybe<Uuid_Comparison_Exp>;
  block_referrer?: InputMaybe<Block_Bool_Exp>;
  block_referrer_id?: InputMaybe<Uuid_Comparison_Exp>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
};

/** unique or primary key constraints on table "block_reference" */
export enum Block_Reference_Constraint {
  /** unique or primary key constraint */
  BlockReferencePkey = 'block_reference_pkey'
}

/** input type for inserting data into table "block_reference" */
export type Block_Reference_Insert_Input = {
  block_referee?: InputMaybe<Block_Obj_Rel_Insert_Input>;
  block_referee_id?: InputMaybe<Scalars['uuid']>;
  block_referrer?: InputMaybe<Block_Obj_Rel_Insert_Input>;
  block_referrer_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
};

/** aggregate max on columns */
export type Block_Reference_Max_Fields = {
  __typename?: 'block_reference_max_fields';
  block_referee_id?: Maybe<Scalars['uuid']>;
  block_referrer_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
};

/** order by max() on columns of table "block_reference" */
export type Block_Reference_Max_Order_By = {
  block_referee_id?: InputMaybe<Order_By>;
  block_referrer_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type Block_Reference_Min_Fields = {
  __typename?: 'block_reference_min_fields';
  block_referee_id?: Maybe<Scalars['uuid']>;
  block_referrer_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
};

/** order by min() on columns of table "block_reference" */
export type Block_Reference_Min_Order_By = {
  block_referee_id?: InputMaybe<Order_By>;
  block_referrer_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "block_reference" */
export type Block_Reference_Mutation_Response = {
  __typename?: 'block_reference_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Block_Reference>;
};

/** on conflict condition type for table "block_reference" */
export type Block_Reference_On_Conflict = {
  constraint: Block_Reference_Constraint;
  update_columns?: Array<Block_Reference_Update_Column>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};

/** Ordering options when selecting data from "block_reference". */
export type Block_Reference_Order_By = {
  block_referee?: InputMaybe<Block_Order_By>;
  block_referee_id?: InputMaybe<Order_By>;
  block_referrer?: InputMaybe<Block_Order_By>;
  block_referrer_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
};

/** primary key columns input for table: block_reference */
export type Block_Reference_Pk_Columns_Input = {
  block_referee_id: Scalars['uuid'];
  block_referrer_id: Scalars['uuid'];
};

/** select columns of table "block_reference" */
export enum Block_Reference_Select_Column {
  /** column name */
  BlockRefereeId = 'block_referee_id',
  /** column name */
  BlockReferrerId = 'block_referrer_id',
  /** column name */
  CreatedAt = 'created_at'
}

/** input type for updating data in table "block_reference" */
export type Block_Reference_Set_Input = {
  block_referee_id?: InputMaybe<Scalars['uuid']>;
  block_referrer_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
};

/** update columns of table "block_reference" */
export enum Block_Reference_Update_Column {
  /** column name */
  BlockRefereeId = 'block_referee_id',
  /** column name */
  BlockReferrerId = 'block_referrer_id',
  /** column name */
  CreatedAt = 'created_at'
}

/** columns and relationships of "block_repost" */
export type Block_Repost = {
  __typename?: 'block_repost';
  /** An object relationship */
  block: Block;
  block_id: Scalars['uuid'];
  created_at: Scalars['timestamptz'];
  /** An object relationship */
  user: Users;
  user_id: Scalars['uuid'];
};

/** aggregated selection of "block_repost" */
export type Block_Repost_Aggregate = {
  __typename?: 'block_repost_aggregate';
  aggregate?: Maybe<Block_Repost_Aggregate_Fields>;
  nodes: Array<Block_Repost>;
};

/** aggregate fields of "block_repost" */
export type Block_Repost_Aggregate_Fields = {
  __typename?: 'block_repost_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Block_Repost_Max_Fields>;
  min?: Maybe<Block_Repost_Min_Fields>;
};


/** aggregate fields of "block_repost" */
export type Block_Repost_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Block_Repost_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "block_repost" */
export type Block_Repost_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<Block_Repost_Max_Order_By>;
  min?: InputMaybe<Block_Repost_Min_Order_By>;
};

/** input type for inserting array relation for remote table "block_repost" */
export type Block_Repost_Arr_Rel_Insert_Input = {
  data: Array<Block_Repost_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<Block_Repost_On_Conflict>;
};

/** Boolean expression to filter rows from the table "block_repost". All fields are combined with a logical 'AND'. */
export type Block_Repost_Bool_Exp = {
  _and?: InputMaybe<Array<Block_Repost_Bool_Exp>>;
  _not?: InputMaybe<Block_Repost_Bool_Exp>;
  _or?: InputMaybe<Array<Block_Repost_Bool_Exp>>;
  block?: InputMaybe<Block_Bool_Exp>;
  block_id?: InputMaybe<Uuid_Comparison_Exp>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  user_id?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "block_repost" */
export enum Block_Repost_Constraint {
  /** unique or primary key constraint */
  BlockPostPkey = 'block_post_pkey'
}

/** input type for inserting data into table "block_repost" */
export type Block_Repost_Insert_Input = {
  block?: InputMaybe<Block_Obj_Rel_Insert_Input>;
  block_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type Block_Repost_Max_Fields = {
  __typename?: 'block_repost_max_fields';
  block_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "block_repost" */
export type Block_Repost_Max_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type Block_Repost_Min_Fields = {
  __typename?: 'block_repost_min_fields';
  block_id?: Maybe<Scalars['uuid']>;
  created_at?: Maybe<Scalars['timestamptz']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "block_repost" */
export type Block_Repost_Min_Order_By = {
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "block_repost" */
export type Block_Repost_Mutation_Response = {
  __typename?: 'block_repost_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Block_Repost>;
};

/** on conflict condition type for table "block_repost" */
export type Block_Repost_On_Conflict = {
  constraint: Block_Repost_Constraint;
  update_columns?: Array<Block_Repost_Update_Column>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};

/** Ordering options when selecting data from "block_repost". */
export type Block_Repost_Order_By = {
  block?: InputMaybe<Block_Order_By>;
  block_id?: InputMaybe<Order_By>;
  created_at?: InputMaybe<Order_By>;
  user?: InputMaybe<Users_Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** primary key columns input for table: block_repost */
export type Block_Repost_Pk_Columns_Input = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};

/** select columns of table "block_repost" */
export enum Block_Repost_Select_Column {
  /** column name */
  BlockId = 'block_id',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  UserId = 'user_id'
}

/** input type for updating data in table "block_repost" */
export type Block_Repost_Set_Input = {
  block_id?: InputMaybe<Scalars['uuid']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "block_repost" */
export enum Block_Repost_Update_Column {
  /** column name */
  BlockId = 'block_id',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  UserId = 'user_id'
}

/** select columns of table "block" */
export enum Block_Select_Column {
  /** column name */
  BlockParentId = 'block_parent_id',
  /** column name */
  Children = 'children',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  Data = 'data',
  /** column name */
  ForkParentId = 'fork_parent_id',
  /** column name */
  Id = 'id',
  /** column name */
  IsPosted = 'is_posted',
  /** column name */
  Order = 'order',
  /** column name */
  Path = 'path',
  /** column name */
  Text = 'text',
  /** column name */
  Title = 'title',
  /** column name */
  Type = 'type',
  /** column name */
  UpdatedAt = 'updated_at',
  /** column name */
  UserId = 'user_id'
}

/** input type for updating data in table "block" */
export type Block_Set_Input = {
  /** Block parent id for hierarchy. If null, the block will be at the root */
  block_parent_id?: InputMaybe<Scalars['uuid']>;
  /** Slate children. If null, empty text */
  children?: InputMaybe<Scalars['jsonb']>;
  created_at?: InputMaybe<Scalars['timestamptz']>;
  /** Slate properties */
  data?: InputMaybe<Scalars['jsonb']>;
  /** Forked from */
  fork_parent_id?: InputMaybe<Scalars['uuid']>;
  id?: InputMaybe<Scalars['uuid']>;
  is_posted?: InputMaybe<Scalars['Boolean']>;
  /** Lexicographical order of the block inside the parent block */
  order?: InputMaybe<Scalars['String']>;
  path?: InputMaybe<Scalars['ltree']>;
  /** Full text without marks */
  text?: InputMaybe<Scalars['String']>;
  title?: InputMaybe<Scalars['String']>;
  /** Type of the block. If null, it's a paragraph */
  type?: InputMaybe<Scalars['String']>;
  updated_at?: InputMaybe<Scalars['timestamptz']>;
  /** Author id */
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "block" */
export enum Block_Update_Column {
  /** column name */
  BlockParentId = 'block_parent_id',
  /** column name */
  Children = 'children',
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  Data = 'data',
  /** column name */
  ForkParentId = 'fork_parent_id',
  /** column name */
  Id = 'id',
  /** column name */
  IsPosted = 'is_posted',
  /** column name */
  Order = 'order',
  /** column name */
  Path = 'path',
  /** column name */
  Text = 'text',
  /** column name */
  Title = 'title',
  /** column name */
  Type = 'type',
  /** column name */
  UpdatedAt = 'updated_at',
  /** column name */
  UserId = 'user_id'
}

/** Boolean expression to compare columns of type "citext". All fields are combined with logical 'AND'. */
export type Citext_Comparison_Exp = {
  _eq?: InputMaybe<Scalars['citext']>;
  _gt?: InputMaybe<Scalars['citext']>;
  _gte?: InputMaybe<Scalars['citext']>;
  /** does the column match the given case-insensitive pattern */
  _ilike?: InputMaybe<Scalars['citext']>;
  _in?: InputMaybe<Array<Scalars['citext']>>;
  /** does the column match the given POSIX regular expression, case insensitive */
  _iregex?: InputMaybe<Scalars['citext']>;
  _is_null?: InputMaybe<Scalars['Boolean']>;
  /** does the column match the given pattern */
  _like?: InputMaybe<Scalars['citext']>;
  _lt?: InputMaybe<Scalars['citext']>;
  _lte?: InputMaybe<Scalars['citext']>;
  _neq?: InputMaybe<Scalars['citext']>;
  /** does the column NOT match the given case-insensitive pattern */
  _nilike?: InputMaybe<Scalars['citext']>;
  _nin?: InputMaybe<Array<Scalars['citext']>>;
  /** does the column NOT match the given POSIX regular expression, case insensitive */
  _niregex?: InputMaybe<Scalars['citext']>;
  /** does the column NOT match the given pattern */
  _nlike?: InputMaybe<Scalars['citext']>;
  /** does the column NOT match the given POSIX regular expression, case sensitive */
  _nregex?: InputMaybe<Scalars['citext']>;
  /** does the column NOT match the given SQL regular expression */
  _nsimilar?: InputMaybe<Scalars['citext']>;
  /** does the column match the given POSIX regular expression, case sensitive */
  _regex?: InputMaybe<Scalars['citext']>;
  /** does the column match the given SQL regular expression */
  _similar?: InputMaybe<Scalars['citext']>;
};

/** Boolean expression to compare columns of type "date". All fields are combined with logical 'AND'. */
export type Date_Comparison_Exp = {
  _eq?: InputMaybe<Scalars['date']>;
  _gt?: InputMaybe<Scalars['date']>;
  _gte?: InputMaybe<Scalars['date']>;
  _in?: InputMaybe<Array<Scalars['date']>>;
  _is_null?: InputMaybe<Scalars['Boolean']>;
  _lt?: InputMaybe<Scalars['date']>;
  _lte?: InputMaybe<Scalars['date']>;
  _neq?: InputMaybe<Scalars['date']>;
  _nin?: InputMaybe<Array<Scalars['date']>>;
};

/** columns and relationships of "follower" */
export type Follower = {
  __typename?: 'follower';
  created_at?: Maybe<Scalars['timestamptz']>;
  /** An object relationship */
  follower: Users;
  follower_id: Scalars['uuid'];
  /** An object relationship */
  following: Users;
  following_id: Scalars['uuid'];
};

/** aggregated selection of "follower" */
export type Follower_Aggregate = {
  __typename?: 'follower_aggregate';
  aggregate?: Maybe<Follower_Aggregate_Fields>;
  nodes: Array<Follower>;
};

/** aggregate fields of "follower" */
export type Follower_Aggregate_Fields = {
  __typename?: 'follower_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Follower_Max_Fields>;
  min?: Maybe<Follower_Min_Fields>;
};


/** aggregate fields of "follower" */
export type Follower_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Follower_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "follower" */
export type Follower_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<Follower_Max_Order_By>;
  min?: InputMaybe<Follower_Min_Order_By>;
};

/** input type for inserting array relation for remote table "follower" */
export type Follower_Arr_Rel_Insert_Input = {
  data: Array<Follower_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<Follower_On_Conflict>;
};

/** Boolean expression to filter rows from the table "follower". All fields are combined with a logical 'AND'. */
export type Follower_Bool_Exp = {
  _and?: InputMaybe<Array<Follower_Bool_Exp>>;
  _not?: InputMaybe<Follower_Bool_Exp>;
  _or?: InputMaybe<Array<Follower_Bool_Exp>>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  follower?: InputMaybe<Users_Bool_Exp>;
  follower_id?: InputMaybe<Uuid_Comparison_Exp>;
  following?: InputMaybe<Users_Bool_Exp>;
  following_id?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "follower" */
export enum Follower_Constraint {
  /** unique or primary key constraint */
  FollowerPkey = 'follower_pkey'
}

/** input type for inserting data into table "follower" */
export type Follower_Insert_Input = {
  created_at?: InputMaybe<Scalars['timestamptz']>;
  follower?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  follower_id?: InputMaybe<Scalars['uuid']>;
  following?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  following_id?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type Follower_Max_Fields = {
  __typename?: 'follower_max_fields';
  created_at?: Maybe<Scalars['timestamptz']>;
  follower_id?: Maybe<Scalars['uuid']>;
  following_id?: Maybe<Scalars['uuid']>;
};

/** order by max() on columns of table "follower" */
export type Follower_Max_Order_By = {
  created_at?: InputMaybe<Order_By>;
  follower_id?: InputMaybe<Order_By>;
  following_id?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type Follower_Min_Fields = {
  __typename?: 'follower_min_fields';
  created_at?: Maybe<Scalars['timestamptz']>;
  follower_id?: Maybe<Scalars['uuid']>;
  following_id?: Maybe<Scalars['uuid']>;
};

/** order by min() on columns of table "follower" */
export type Follower_Min_Order_By = {
  created_at?: InputMaybe<Order_By>;
  follower_id?: InputMaybe<Order_By>;
  following_id?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "follower" */
export type Follower_Mutation_Response = {
  __typename?: 'follower_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Follower>;
};

/** on conflict condition type for table "follower" */
export type Follower_On_Conflict = {
  constraint: Follower_Constraint;
  update_columns?: Array<Follower_Update_Column>;
  where?: InputMaybe<Follower_Bool_Exp>;
};

/** Ordering options when selecting data from "follower". */
export type Follower_Order_By = {
  created_at?: InputMaybe<Order_By>;
  follower?: InputMaybe<Users_Order_By>;
  follower_id?: InputMaybe<Order_By>;
  following?: InputMaybe<Users_Order_By>;
  following_id?: InputMaybe<Order_By>;
};

/** primary key columns input for table: follower */
export type Follower_Pk_Columns_Input = {
  follower_id: Scalars['uuid'];
  following_id: Scalars['uuid'];
};

/** select columns of table "follower" */
export enum Follower_Select_Column {
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  FollowerId = 'follower_id',
  /** column name */
  FollowingId = 'following_id'
}

/** input type for updating data in table "follower" */
export type Follower_Set_Input = {
  created_at?: InputMaybe<Scalars['timestamptz']>;
  follower_id?: InputMaybe<Scalars['uuid']>;
  following_id?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "follower" */
export enum Follower_Update_Column {
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  FollowerId = 'follower_id',
  /** column name */
  FollowingId = 'following_id'
}

/** Boolean expression to compare columns of type "jsonb". All fields are combined with logical 'AND'. */
export type Jsonb_Comparison_Exp = {
  /** is the column contained in the given json value */
  _contained_in?: InputMaybe<Scalars['jsonb']>;
  /** does the column contain the given json value at the top level */
  _contains?: InputMaybe<Scalars['jsonb']>;
  _eq?: InputMaybe<Scalars['jsonb']>;
  _gt?: InputMaybe<Scalars['jsonb']>;
  _gte?: InputMaybe<Scalars['jsonb']>;
  /** does the string exist as a top-level key in the column */
  _has_key?: InputMaybe<Scalars['String']>;
  /** do all of these strings exist as top-level keys in the column */
  _has_keys_all?: InputMaybe<Array<Scalars['String']>>;
  /** do any of these strings exist as top-level keys in the column */
  _has_keys_any?: InputMaybe<Array<Scalars['String']>>;
  _in?: InputMaybe<Array<Scalars['jsonb']>>;
  _is_null?: InputMaybe<Scalars['Boolean']>;
  _lt?: InputMaybe<Scalars['jsonb']>;
  _lte?: InputMaybe<Scalars['jsonb']>;
  _neq?: InputMaybe<Scalars['jsonb']>;
  _nin?: InputMaybe<Array<Scalars['jsonb']>>;
};

/** Boolean expression to compare columns of type "ltree". All fields are combined with logical 'AND'. */
export type Ltree_Comparison_Exp = {
  /** is the left argument an ancestor of right (or equal)? */
  _ancestor?: InputMaybe<Scalars['ltree']>;
  /** does array contain an ancestor of `ltree`? */
  _ancestor_any?: InputMaybe<Array<Scalars['ltree']>>;
  /** is the left argument a descendant of right (or equal)? */
  _descendant?: InputMaybe<Scalars['ltree']>;
  /** does array contain a descendant of `ltree`? */
  _descendant_any?: InputMaybe<Array<Scalars['ltree']>>;
  _eq?: InputMaybe<Scalars['ltree']>;
  _gt?: InputMaybe<Scalars['ltree']>;
  _gte?: InputMaybe<Scalars['ltree']>;
  _in?: InputMaybe<Array<Scalars['ltree']>>;
  _is_null?: InputMaybe<Scalars['Boolean']>;
  _lt?: InputMaybe<Scalars['ltree']>;
  _lte?: InputMaybe<Scalars['ltree']>;
  /** does `ltree` match `lquery`? */
  _matches?: InputMaybe<Scalars['lquery']>;
  /** does `ltree` match any `lquery` in array? */
  _matches_any?: InputMaybe<Array<Scalars['String']>>;
  /** does `ltree` match `ltxtquery`? */
  _matches_fulltext?: InputMaybe<Scalars['ltxtquery']>;
  _neq?: InputMaybe<Scalars['ltree']>;
  _nin?: InputMaybe<Array<Scalars['ltree']>>;
};

/** mutation root */
export type Mutation_Root = {
  __typename?: 'mutation_root';
  /** delete single row from the table: "auth.providers" */
  deleteAuthProvider?: Maybe<AuthProviders>;
  /** delete single row from the table: "auth.provider_requests" */
  deleteAuthProviderRequest?: Maybe<AuthProviderRequests>;
  /** delete data from the table: "auth.provider_requests" */
  deleteAuthProviderRequests?: Maybe<AuthProviderRequests_Mutation_Response>;
  /** delete data from the table: "auth.providers" */
  deleteAuthProviders?: Maybe<AuthProviders_Mutation_Response>;
  /** delete single row from the table: "auth.refresh_tokens" */
  deleteAuthRefreshToken?: Maybe<AuthRefreshTokens>;
  /** delete data from the table: "auth.refresh_tokens" */
  deleteAuthRefreshTokens?: Maybe<AuthRefreshTokens_Mutation_Response>;
  /** delete single row from the table: "auth.roles" */
  deleteAuthRole?: Maybe<AuthRoles>;
  /** delete data from the table: "auth.roles" */
  deleteAuthRoles?: Maybe<AuthRoles_Mutation_Response>;
  /** delete single row from the table: "auth.user_providers" */
  deleteAuthUserProvider?: Maybe<AuthUserProviders>;
  /** delete data from the table: "auth.user_providers" */
  deleteAuthUserProviders?: Maybe<AuthUserProviders_Mutation_Response>;
  /** delete single row from the table: "auth.user_roles" */
  deleteAuthUserRole?: Maybe<AuthUserRoles>;
  /** delete data from the table: "auth.user_roles" */
  deleteAuthUserRoles?: Maybe<AuthUserRoles_Mutation_Response>;
  /** delete data from the table: "auth.user_private" */
  delete_auth_user_private?: Maybe<Auth_User_Private_Mutation_Response>;
  /** delete data from the table: "block" */
  delete_block?: Maybe<Block_Mutation_Response>;
  /** delete single row from the table: "block" */
  delete_block_by_pk?: Maybe<Block>;
  /** delete data from the table: "block_comment" */
  delete_block_comment?: Maybe<Block_Comment_Mutation_Response>;
  /** delete single row from the table: "block_comment" */
  delete_block_comment_by_pk?: Maybe<Block_Comment>;
  /** delete data from the table: "block_follower" */
  delete_block_follower?: Maybe<Block_Follower_Mutation_Response>;
  /** delete single row from the table: "block_follower" */
  delete_block_follower_by_pk?: Maybe<Block_Follower>;
  /** delete data from the table: "block_like" */
  delete_block_like?: Maybe<Block_Like_Mutation_Response>;
  /** delete single row from the table: "block_like" */
  delete_block_like_by_pk?: Maybe<Block_Like>;
  /** delete data from the table: "block_reference" */
  delete_block_reference?: Maybe<Block_Reference_Mutation_Response>;
  /** delete single row from the table: "block_reference" */
  delete_block_reference_by_pk?: Maybe<Block_Reference>;
  /** delete data from the table: "block_repost" */
  delete_block_repost?: Maybe<Block_Repost_Mutation_Response>;
  /** delete single row from the table: "block_repost" */
  delete_block_repost_by_pk?: Maybe<Block_Repost>;
  /** delete data from the table: "follower" */
  delete_follower?: Maybe<Follower_Mutation_Response>;
  /** delete single row from the table: "follower" */
  delete_follower_by_pk?: Maybe<Follower>;
  /** delete data from the table: "auth.users" */
  delete_users?: Maybe<Users_Mutation_Response>;
  /** delete single row from the table: "auth.users" */
  delete_users_by_pk?: Maybe<Users>;
  /** delete data from the table: "workspace" */
  delete_workspace?: Maybe<Workspace_Mutation_Response>;
  /** delete single row from the table: "workspace" */
  delete_workspace_by_pk?: Maybe<Workspace>;
  /** insert a single row into the table: "auth.providers" */
  insertAuthProvider?: Maybe<AuthProviders>;
  /** insert a single row into the table: "auth.provider_requests" */
  insertAuthProviderRequest?: Maybe<AuthProviderRequests>;
  /** insert data into the table: "auth.provider_requests" */
  insertAuthProviderRequests?: Maybe<AuthProviderRequests_Mutation_Response>;
  /** insert data into the table: "auth.providers" */
  insertAuthProviders?: Maybe<AuthProviders_Mutation_Response>;
  /** insert a single row into the table: "auth.refresh_tokens" */
  insertAuthRefreshToken?: Maybe<AuthRefreshTokens>;
  /** insert data into the table: "auth.refresh_tokens" */
  insertAuthRefreshTokens?: Maybe<AuthRefreshTokens_Mutation_Response>;
  /** insert a single row into the table: "auth.roles" */
  insertAuthRole?: Maybe<AuthRoles>;
  /** insert data into the table: "auth.roles" */
  insertAuthRoles?: Maybe<AuthRoles_Mutation_Response>;
  /** insert a single row into the table: "auth.user_providers" */
  insertAuthUserProvider?: Maybe<AuthUserProviders>;
  /** insert data into the table: "auth.user_providers" */
  insertAuthUserProviders?: Maybe<AuthUserProviders_Mutation_Response>;
  /** insert a single row into the table: "auth.user_roles" */
  insertAuthUserRole?: Maybe<AuthUserRoles>;
  /** insert data into the table: "auth.user_roles" */
  insertAuthUserRoles?: Maybe<AuthUserRoles_Mutation_Response>;
  /** insert data into the table: "auth.user_private" */
  insert_auth_user_private?: Maybe<Auth_User_Private_Mutation_Response>;
  /** insert a single row into the table: "auth.user_private" */
  insert_auth_user_private_one?: Maybe<Auth_User_Private>;
  /** insert data into the table: "block" */
  insert_block?: Maybe<Block_Mutation_Response>;
  /** insert data into the table: "block_comment" */
  insert_block_comment?: Maybe<Block_Comment_Mutation_Response>;
  /** insert a single row into the table: "block_comment" */
  insert_block_comment_one?: Maybe<Block_Comment>;
  /** insert data into the table: "block_follower" */
  insert_block_follower?: Maybe<Block_Follower_Mutation_Response>;
  /** insert a single row into the table: "block_follower" */
  insert_block_follower_one?: Maybe<Block_Follower>;
  /** insert data into the table: "block_like" */
  insert_block_like?: Maybe<Block_Like_Mutation_Response>;
  /** insert a single row into the table: "block_like" */
  insert_block_like_one?: Maybe<Block_Like>;
  /** insert a single row into the table: "block" */
  insert_block_one?: Maybe<Block>;
  /** insert data into the table: "block_reference" */
  insert_block_reference?: Maybe<Block_Reference_Mutation_Response>;
  /** insert a single row into the table: "block_reference" */
  insert_block_reference_one?: Maybe<Block_Reference>;
  /** insert data into the table: "block_repost" */
  insert_block_repost?: Maybe<Block_Repost_Mutation_Response>;
  /** insert a single row into the table: "block_repost" */
  insert_block_repost_one?: Maybe<Block_Repost>;
  /** insert data into the table: "follower" */
  insert_follower?: Maybe<Follower_Mutation_Response>;
  /** insert a single row into the table: "follower" */
  insert_follower_one?: Maybe<Follower>;
  /** insert data into the table: "auth.users" */
  insert_users?: Maybe<Users_Mutation_Response>;
  /** insert a single row into the table: "auth.users" */
  insert_users_one?: Maybe<Users>;
  /** insert data into the table: "workspace" */
  insert_workspace?: Maybe<Workspace_Mutation_Response>;
  /** insert a single row into the table: "workspace" */
  insert_workspace_one?: Maybe<Workspace>;
  save_tags?: Maybe<Save_Tags_Output>;
  /** update single row of the table: "auth.providers" */
  updateAuthProvider?: Maybe<AuthProviders>;
  /** update single row of the table: "auth.provider_requests" */
  updateAuthProviderRequest?: Maybe<AuthProviderRequests>;
  /** update data of the table: "auth.provider_requests" */
  updateAuthProviderRequests?: Maybe<AuthProviderRequests_Mutation_Response>;
  /** update data of the table: "auth.providers" */
  updateAuthProviders?: Maybe<AuthProviders_Mutation_Response>;
  /** update single row of the table: "auth.refresh_tokens" */
  updateAuthRefreshToken?: Maybe<AuthRefreshTokens>;
  /** update data of the table: "auth.refresh_tokens" */
  updateAuthRefreshTokens?: Maybe<AuthRefreshTokens_Mutation_Response>;
  /** update single row of the table: "auth.roles" */
  updateAuthRole?: Maybe<AuthRoles>;
  /** update data of the table: "auth.roles" */
  updateAuthRoles?: Maybe<AuthRoles_Mutation_Response>;
  /** update single row of the table: "auth.user_providers" */
  updateAuthUserProvider?: Maybe<AuthUserProviders>;
  /** update data of the table: "auth.user_providers" */
  updateAuthUserProviders?: Maybe<AuthUserProviders_Mutation_Response>;
  /** update single row of the table: "auth.user_roles" */
  updateAuthUserRole?: Maybe<AuthUserRoles>;
  /** update data of the table: "auth.user_roles" */
  updateAuthUserRoles?: Maybe<AuthUserRoles_Mutation_Response>;
  /** update data of the table: "auth.user_private" */
  update_auth_user_private?: Maybe<Auth_User_Private_Mutation_Response>;
  /** update data of the table: "block" */
  update_block?: Maybe<Block_Mutation_Response>;
  /** update single row of the table: "block" */
  update_block_by_pk?: Maybe<Block>;
  /** update data of the table: "block_comment" */
  update_block_comment?: Maybe<Block_Comment_Mutation_Response>;
  /** update single row of the table: "block_comment" */
  update_block_comment_by_pk?: Maybe<Block_Comment>;
  /** update data of the table: "block_follower" */
  update_block_follower?: Maybe<Block_Follower_Mutation_Response>;
  /** update single row of the table: "block_follower" */
  update_block_follower_by_pk?: Maybe<Block_Follower>;
  /** update data of the table: "block_like" */
  update_block_like?: Maybe<Block_Like_Mutation_Response>;
  /** update single row of the table: "block_like" */
  update_block_like_by_pk?: Maybe<Block_Like>;
  /** update data of the table: "block_reference" */
  update_block_reference?: Maybe<Block_Reference_Mutation_Response>;
  /** update single row of the table: "block_reference" */
  update_block_reference_by_pk?: Maybe<Block_Reference>;
  /** update data of the table: "block_repost" */
  update_block_repost?: Maybe<Block_Repost_Mutation_Response>;
  /** update single row of the table: "block_repost" */
  update_block_repost_by_pk?: Maybe<Block_Repost>;
  /** update data of the table: "follower" */
  update_follower?: Maybe<Follower_Mutation_Response>;
  /** update single row of the table: "follower" */
  update_follower_by_pk?: Maybe<Follower>;
  /** update data of the table: "auth.users" */
  update_users?: Maybe<Users_Mutation_Response>;
  /** update single row of the table: "auth.users" */
  update_users_by_pk?: Maybe<Users>;
  /** update data of the table: "workspace" */
  update_workspace?: Maybe<Workspace_Mutation_Response>;
  /** update single row of the table: "workspace" */
  update_workspace_by_pk?: Maybe<Workspace>;
};


/** mutation root */
export type Mutation_RootDeleteAuthProviderArgs = {
  id: Scalars['String'];
};


/** mutation root */
export type Mutation_RootDeleteAuthProviderRequestArgs = {
  id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDeleteAuthProviderRequestsArgs = {
  where: AuthProviderRequests_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDeleteAuthProvidersArgs = {
  where: AuthProviders_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDeleteAuthRefreshTokenArgs = {
  refreshToken: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDeleteAuthRefreshTokensArgs = {
  where: AuthRefreshTokens_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDeleteAuthRoleArgs = {
  role: Scalars['String'];
};


/** mutation root */
export type Mutation_RootDeleteAuthRolesArgs = {
  where: AuthRoles_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDeleteAuthUserProviderArgs = {
  id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDeleteAuthUserProvidersArgs = {
  where: AuthUserProviders_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDeleteAuthUserRoleArgs = {
  id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDeleteAuthUserRolesArgs = {
  where: AuthUserRoles_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Auth_User_PrivateArgs = {
  where: Auth_User_Private_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_BlockArgs = {
  where: Block_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Block_By_PkArgs = {
  id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDelete_Block_CommentArgs = {
  where: Block_Comment_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Block_Comment_By_PkArgs = {
  id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDelete_Block_FollowerArgs = {
  where: Block_Follower_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Block_Follower_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDelete_Block_LikeArgs = {
  where: Block_Like_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Block_Like_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDelete_Block_ReferenceArgs = {
  where: Block_Reference_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Block_Reference_By_PkArgs = {
  block_referee_id: Scalars['uuid'];
  block_referrer_id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDelete_Block_RepostArgs = {
  where: Block_Repost_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Block_Repost_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDelete_FollowerArgs = {
  where: Follower_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Follower_By_PkArgs = {
  follower_id: Scalars['uuid'];
  following_id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDelete_UsersArgs = {
  where: Users_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Users_By_PkArgs = {
  id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootDelete_WorkspaceArgs = {
  where: Workspace_Bool_Exp;
};


/** mutation root */
export type Mutation_RootDelete_Workspace_By_PkArgs = {
  id: Scalars['uuid'];
};


/** mutation root */
export type Mutation_RootInsertAuthProviderArgs = {
  object: AuthProviders_Insert_Input;
  on_conflict?: InputMaybe<AuthProviders_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthProviderRequestArgs = {
  object: AuthProviderRequests_Insert_Input;
  on_conflict?: InputMaybe<AuthProviderRequests_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthProviderRequestsArgs = {
  objects: Array<AuthProviderRequests_Insert_Input>;
  on_conflict?: InputMaybe<AuthProviderRequests_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthProvidersArgs = {
  objects: Array<AuthProviders_Insert_Input>;
  on_conflict?: InputMaybe<AuthProviders_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthRefreshTokenArgs = {
  object: AuthRefreshTokens_Insert_Input;
  on_conflict?: InputMaybe<AuthRefreshTokens_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthRefreshTokensArgs = {
  objects: Array<AuthRefreshTokens_Insert_Input>;
  on_conflict?: InputMaybe<AuthRefreshTokens_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthRoleArgs = {
  object: AuthRoles_Insert_Input;
  on_conflict?: InputMaybe<AuthRoles_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthRolesArgs = {
  objects: Array<AuthRoles_Insert_Input>;
  on_conflict?: InputMaybe<AuthRoles_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthUserProviderArgs = {
  object: AuthUserProviders_Insert_Input;
  on_conflict?: InputMaybe<AuthUserProviders_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthUserProvidersArgs = {
  objects: Array<AuthUserProviders_Insert_Input>;
  on_conflict?: InputMaybe<AuthUserProviders_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthUserRoleArgs = {
  object: AuthUserRoles_Insert_Input;
  on_conflict?: InputMaybe<AuthUserRoles_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsertAuthUserRolesArgs = {
  objects: Array<AuthUserRoles_Insert_Input>;
  on_conflict?: InputMaybe<AuthUserRoles_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Auth_User_PrivateArgs = {
  objects: Array<Auth_User_Private_Insert_Input>;
};


/** mutation root */
export type Mutation_RootInsert_Auth_User_Private_OneArgs = {
  object: Auth_User_Private_Insert_Input;
};


/** mutation root */
export type Mutation_RootInsert_BlockArgs = {
  objects: Array<Block_Insert_Input>;
  on_conflict?: InputMaybe<Block_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_CommentArgs = {
  objects: Array<Block_Comment_Insert_Input>;
  on_conflict?: InputMaybe<Block_Comment_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_Comment_OneArgs = {
  object: Block_Comment_Insert_Input;
  on_conflict?: InputMaybe<Block_Comment_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_FollowerArgs = {
  objects: Array<Block_Follower_Insert_Input>;
  on_conflict?: InputMaybe<Block_Follower_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_Follower_OneArgs = {
  object: Block_Follower_Insert_Input;
  on_conflict?: InputMaybe<Block_Follower_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_LikeArgs = {
  objects: Array<Block_Like_Insert_Input>;
  on_conflict?: InputMaybe<Block_Like_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_Like_OneArgs = {
  object: Block_Like_Insert_Input;
  on_conflict?: InputMaybe<Block_Like_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_OneArgs = {
  object: Block_Insert_Input;
  on_conflict?: InputMaybe<Block_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_ReferenceArgs = {
  objects: Array<Block_Reference_Insert_Input>;
  on_conflict?: InputMaybe<Block_Reference_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_Reference_OneArgs = {
  object: Block_Reference_Insert_Input;
  on_conflict?: InputMaybe<Block_Reference_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_RepostArgs = {
  objects: Array<Block_Repost_Insert_Input>;
  on_conflict?: InputMaybe<Block_Repost_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Block_Repost_OneArgs = {
  object: Block_Repost_Insert_Input;
  on_conflict?: InputMaybe<Block_Repost_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_FollowerArgs = {
  objects: Array<Follower_Insert_Input>;
  on_conflict?: InputMaybe<Follower_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Follower_OneArgs = {
  object: Follower_Insert_Input;
  on_conflict?: InputMaybe<Follower_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_UsersArgs = {
  objects: Array<Users_Insert_Input>;
  on_conflict?: InputMaybe<Users_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Users_OneArgs = {
  object: Users_Insert_Input;
  on_conflict?: InputMaybe<Users_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_WorkspaceArgs = {
  objects: Array<Workspace_Insert_Input>;
  on_conflict?: InputMaybe<Workspace_On_Conflict>;
};


/** mutation root */
export type Mutation_RootInsert_Workspace_OneArgs = {
  object: Workspace_Insert_Input;
  on_conflict?: InputMaybe<Workspace_On_Conflict>;
};


/** mutation root */
export type Mutation_RootSave_TagsArgs = {
  tags: Array<Save_Tags_Input>;
};


/** mutation root */
export type Mutation_RootUpdateAuthProviderArgs = {
  _set?: InputMaybe<AuthProviders_Set_Input>;
  pk_columns: AuthProviders_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdateAuthProviderRequestArgs = {
  _append?: InputMaybe<AuthProviderRequests_Append_Input>;
  _delete_at_path?: InputMaybe<AuthProviderRequests_Delete_At_Path_Input>;
  _delete_elem?: InputMaybe<AuthProviderRequests_Delete_Elem_Input>;
  _delete_key?: InputMaybe<AuthProviderRequests_Delete_Key_Input>;
  _prepend?: InputMaybe<AuthProviderRequests_Prepend_Input>;
  _set?: InputMaybe<AuthProviderRequests_Set_Input>;
  pk_columns: AuthProviderRequests_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdateAuthProviderRequestsArgs = {
  _append?: InputMaybe<AuthProviderRequests_Append_Input>;
  _delete_at_path?: InputMaybe<AuthProviderRequests_Delete_At_Path_Input>;
  _delete_elem?: InputMaybe<AuthProviderRequests_Delete_Elem_Input>;
  _delete_key?: InputMaybe<AuthProviderRequests_Delete_Key_Input>;
  _prepend?: InputMaybe<AuthProviderRequests_Prepend_Input>;
  _set?: InputMaybe<AuthProviderRequests_Set_Input>;
  where: AuthProviderRequests_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdateAuthProvidersArgs = {
  _set?: InputMaybe<AuthProviders_Set_Input>;
  where: AuthProviders_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdateAuthRefreshTokenArgs = {
  _set?: InputMaybe<AuthRefreshTokens_Set_Input>;
  pk_columns: AuthRefreshTokens_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdateAuthRefreshTokensArgs = {
  _set?: InputMaybe<AuthRefreshTokens_Set_Input>;
  where: AuthRefreshTokens_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdateAuthRoleArgs = {
  _set?: InputMaybe<AuthRoles_Set_Input>;
  pk_columns: AuthRoles_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdateAuthRolesArgs = {
  _set?: InputMaybe<AuthRoles_Set_Input>;
  where: AuthRoles_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdateAuthUserProviderArgs = {
  _set?: InputMaybe<AuthUserProviders_Set_Input>;
  pk_columns: AuthUserProviders_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdateAuthUserProvidersArgs = {
  _set?: InputMaybe<AuthUserProviders_Set_Input>;
  where: AuthUserProviders_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdateAuthUserRoleArgs = {
  _set?: InputMaybe<AuthUserRoles_Set_Input>;
  pk_columns: AuthUserRoles_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdateAuthUserRolesArgs = {
  _set?: InputMaybe<AuthUserRoles_Set_Input>;
  where: AuthUserRoles_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Auth_User_PrivateArgs = {
  _set?: InputMaybe<Auth_User_Private_Set_Input>;
  where: Auth_User_Private_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_BlockArgs = {
  _append?: InputMaybe<Block_Append_Input>;
  _delete_at_path?: InputMaybe<Block_Delete_At_Path_Input>;
  _delete_elem?: InputMaybe<Block_Delete_Elem_Input>;
  _delete_key?: InputMaybe<Block_Delete_Key_Input>;
  _prepend?: InputMaybe<Block_Prepend_Input>;
  _set?: InputMaybe<Block_Set_Input>;
  where: Block_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Block_By_PkArgs = {
  _append?: InputMaybe<Block_Append_Input>;
  _delete_at_path?: InputMaybe<Block_Delete_At_Path_Input>;
  _delete_elem?: InputMaybe<Block_Delete_Elem_Input>;
  _delete_key?: InputMaybe<Block_Delete_Key_Input>;
  _prepend?: InputMaybe<Block_Prepend_Input>;
  _set?: InputMaybe<Block_Set_Input>;
  pk_columns: Block_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_Block_CommentArgs = {
  _set?: InputMaybe<Block_Comment_Set_Input>;
  where: Block_Comment_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Block_Comment_By_PkArgs = {
  _set?: InputMaybe<Block_Comment_Set_Input>;
  pk_columns: Block_Comment_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_Block_FollowerArgs = {
  _set?: InputMaybe<Block_Follower_Set_Input>;
  where: Block_Follower_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Block_Follower_By_PkArgs = {
  _set?: InputMaybe<Block_Follower_Set_Input>;
  pk_columns: Block_Follower_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_Block_LikeArgs = {
  _set?: InputMaybe<Block_Like_Set_Input>;
  where: Block_Like_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Block_Like_By_PkArgs = {
  _set?: InputMaybe<Block_Like_Set_Input>;
  pk_columns: Block_Like_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_Block_ReferenceArgs = {
  _set?: InputMaybe<Block_Reference_Set_Input>;
  where: Block_Reference_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Block_Reference_By_PkArgs = {
  _set?: InputMaybe<Block_Reference_Set_Input>;
  pk_columns: Block_Reference_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_Block_RepostArgs = {
  _set?: InputMaybe<Block_Repost_Set_Input>;
  where: Block_Repost_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Block_Repost_By_PkArgs = {
  _set?: InputMaybe<Block_Repost_Set_Input>;
  pk_columns: Block_Repost_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_FollowerArgs = {
  _set?: InputMaybe<Follower_Set_Input>;
  where: Follower_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Follower_By_PkArgs = {
  _set?: InputMaybe<Follower_Set_Input>;
  pk_columns: Follower_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_UsersArgs = {
  _append?: InputMaybe<Users_Append_Input>;
  _delete_at_path?: InputMaybe<Users_Delete_At_Path_Input>;
  _delete_elem?: InputMaybe<Users_Delete_Elem_Input>;
  _delete_key?: InputMaybe<Users_Delete_Key_Input>;
  _prepend?: InputMaybe<Users_Prepend_Input>;
  _set?: InputMaybe<Users_Set_Input>;
  where: Users_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Users_By_PkArgs = {
  _append?: InputMaybe<Users_Append_Input>;
  _delete_at_path?: InputMaybe<Users_Delete_At_Path_Input>;
  _delete_elem?: InputMaybe<Users_Delete_Elem_Input>;
  _delete_key?: InputMaybe<Users_Delete_Key_Input>;
  _prepend?: InputMaybe<Users_Prepend_Input>;
  _set?: InputMaybe<Users_Set_Input>;
  pk_columns: Users_Pk_Columns_Input;
};


/** mutation root */
export type Mutation_RootUpdate_WorkspaceArgs = {
  _set?: InputMaybe<Workspace_Set_Input>;
  where: Workspace_Bool_Exp;
};


/** mutation root */
export type Mutation_RootUpdate_Workspace_By_PkArgs = {
  _set?: InputMaybe<Workspace_Set_Input>;
  pk_columns: Workspace_Pk_Columns_Input;
};

/** column ordering options */
export enum Order_By {
  /** in ascending order, nulls last */
  Asc = 'asc',
  /** in ascending order, nulls first */
  AscNullsFirst = 'asc_nulls_first',
  /** in ascending order, nulls last */
  AscNullsLast = 'asc_nulls_last',
  /** in descending order, nulls first */
  Desc = 'desc',
  /** in descending order, nulls first */
  DescNullsFirst = 'desc_nulls_first',
  /** in descending order, nulls last */
  DescNullsLast = 'desc_nulls_last'
}

export type Query_Root = {
  __typename?: 'query_root';
  /** fetch data from the table: "auth.providers" using primary key columns */
  authProvider?: Maybe<AuthProviders>;
  /** fetch data from the table: "auth.provider_requests" using primary key columns */
  authProviderRequest?: Maybe<AuthProviderRequests>;
  /** fetch data from the table: "auth.provider_requests" */
  authProviderRequests: Array<AuthProviderRequests>;
  /** fetch aggregated fields from the table: "auth.provider_requests" */
  authProviderRequestsAggregate: AuthProviderRequests_Aggregate;
  /** fetch data from the table: "auth.providers" */
  authProviders: Array<AuthProviders>;
  /** fetch aggregated fields from the table: "auth.providers" */
  authProvidersAggregate: AuthProviders_Aggregate;
  /** fetch data from the table: "auth.refresh_tokens" using primary key columns */
  authRefreshToken?: Maybe<AuthRefreshTokens>;
  /** fetch data from the table: "auth.refresh_tokens" */
  authRefreshTokens: Array<AuthRefreshTokens>;
  /** fetch aggregated fields from the table: "auth.refresh_tokens" */
  authRefreshTokensAggregate: AuthRefreshTokens_Aggregate;
  /** fetch data from the table: "auth.roles" using primary key columns */
  authRole?: Maybe<AuthRoles>;
  /** fetch data from the table: "auth.roles" */
  authRoles: Array<AuthRoles>;
  /** fetch aggregated fields from the table: "auth.roles" */
  authRolesAggregate: AuthRoles_Aggregate;
  /** fetch data from the table: "auth.user_providers" using primary key columns */
  authUserProvider?: Maybe<AuthUserProviders>;
  /** fetch data from the table: "auth.user_providers" */
  authUserProviders: Array<AuthUserProviders>;
  /** fetch aggregated fields from the table: "auth.user_providers" */
  authUserProvidersAggregate: AuthUserProviders_Aggregate;
  /** fetch data from the table: "auth.user_roles" using primary key columns */
  authUserRole?: Maybe<AuthUserRoles>;
  /** fetch data from the table: "auth.user_roles" */
  authUserRoles: Array<AuthUserRoles>;
  /** fetch aggregated fields from the table: "auth.user_roles" */
  authUserRolesAggregate: AuthUserRoles_Aggregate;
  /** fetch data from the table: "auth.user_private" */
  auth_user_private: Array<Auth_User_Private>;
  /** fetch aggregated fields from the table: "auth.user_private" */
  auth_user_private_aggregate: Auth_User_Private_Aggregate;
  /** fetch data from the table: "block" */
  block: Array<Block>;
  /** fetch aggregated fields from the table: "block" */
  block_aggregate: Block_Aggregate;
  /** fetch data from the table: "block" using primary key columns */
  block_by_pk?: Maybe<Block>;
  /** fetch data from the table: "block_comment" */
  block_comment: Array<Block_Comment>;
  /** fetch aggregated fields from the table: "block_comment" */
  block_comment_aggregate: Block_Comment_Aggregate;
  /** fetch data from the table: "block_comment" using primary key columns */
  block_comment_by_pk?: Maybe<Block_Comment>;
  /** fetch data from the table: "block_follower" */
  block_follower: Array<Block_Follower>;
  /** fetch aggregated fields from the table: "block_follower" */
  block_follower_aggregate: Block_Follower_Aggregate;
  /** fetch data from the table: "block_follower" using primary key columns */
  block_follower_by_pk?: Maybe<Block_Follower>;
  /** fetch data from the table: "block_like" */
  block_like: Array<Block_Like>;
  /** fetch aggregated fields from the table: "block_like" */
  block_like_aggregate: Block_Like_Aggregate;
  /** fetch data from the table: "block_like" using primary key columns */
  block_like_by_pk?: Maybe<Block_Like>;
  /** fetch data from the table: "block_reference" */
  block_reference: Array<Block_Reference>;
  /** fetch aggregated fields from the table: "block_reference" */
  block_reference_aggregate: Block_Reference_Aggregate;
  /** fetch data from the table: "block_reference" using primary key columns */
  block_reference_by_pk?: Maybe<Block_Reference>;
  /** fetch data from the table: "block_repost" */
  block_repost: Array<Block_Repost>;
  /** fetch aggregated fields from the table: "block_repost" */
  block_repost_aggregate: Block_Repost_Aggregate;
  /** fetch data from the table: "block_repost" using primary key columns */
  block_repost_by_pk?: Maybe<Block_Repost>;
  /** fetch data from the table: "follower" */
  follower: Array<Follower>;
  /** fetch aggregated fields from the table: "follower" */
  follower_aggregate: Follower_Aggregate;
  /** fetch data from the table: "follower" using primary key columns */
  follower_by_pk?: Maybe<Follower>;
  /** fetch data from the table: "timeline_event" */
  timeline_event: Array<Timeline_Event>;
  /** fetch aggregated fields from the table: "timeline_event" */
  timeline_event_aggregate: Timeline_Event_Aggregate;
  /** fetch data from the table: "auth.users" */
  users: Array<Users>;
  /** fetch aggregated fields from the table: "auth.users" */
  users_aggregate: Users_Aggregate;
  /** fetch data from the table: "auth.users" using primary key columns */
  users_by_pk?: Maybe<Users>;
  /** fetch data from the table: "workspace" */
  workspace: Array<Workspace>;
  /** fetch aggregated fields from the table: "workspace" */
  workspace_aggregate: Workspace_Aggregate;
  /** fetch data from the table: "workspace" using primary key columns */
  workspace_by_pk?: Maybe<Workspace>;
};


export type Query_RootAuthProviderArgs = {
  id: Scalars['String'];
};


export type Query_RootAuthProviderRequestArgs = {
  id: Scalars['uuid'];
};


export type Query_RootAuthProviderRequestsArgs = {
  distinct_on?: InputMaybe<Array<AuthProviderRequests_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthProviderRequests_Order_By>>;
  where?: InputMaybe<AuthProviderRequests_Bool_Exp>;
};


export type Query_RootAuthProviderRequestsAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthProviderRequests_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthProviderRequests_Order_By>>;
  where?: InputMaybe<AuthProviderRequests_Bool_Exp>;
};


export type Query_RootAuthProvidersArgs = {
  distinct_on?: InputMaybe<Array<AuthProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthProviders_Order_By>>;
  where?: InputMaybe<AuthProviders_Bool_Exp>;
};


export type Query_RootAuthProvidersAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthProviders_Order_By>>;
  where?: InputMaybe<AuthProviders_Bool_Exp>;
};


export type Query_RootAuthRefreshTokenArgs = {
  refreshToken: Scalars['uuid'];
};


export type Query_RootAuthRefreshTokensArgs = {
  distinct_on?: InputMaybe<Array<AuthRefreshTokens_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRefreshTokens_Order_By>>;
  where?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
};


export type Query_RootAuthRefreshTokensAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthRefreshTokens_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRefreshTokens_Order_By>>;
  where?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
};


export type Query_RootAuthRoleArgs = {
  role: Scalars['String'];
};


export type Query_RootAuthRolesArgs = {
  distinct_on?: InputMaybe<Array<AuthRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRoles_Order_By>>;
  where?: InputMaybe<AuthRoles_Bool_Exp>;
};


export type Query_RootAuthRolesAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRoles_Order_By>>;
  where?: InputMaybe<AuthRoles_Bool_Exp>;
};


export type Query_RootAuthUserProviderArgs = {
  id: Scalars['uuid'];
};


export type Query_RootAuthUserProvidersArgs = {
  distinct_on?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserProviders_Order_By>>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};


export type Query_RootAuthUserProvidersAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserProviders_Order_By>>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};


export type Query_RootAuthUserRoleArgs = {
  id: Scalars['uuid'];
};


export type Query_RootAuthUserRolesArgs = {
  distinct_on?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserRoles_Order_By>>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};


export type Query_RootAuthUserRolesAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserRoles_Order_By>>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};


export type Query_RootAuth_User_PrivateArgs = {
  distinct_on?: InputMaybe<Array<Auth_User_Private_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Auth_User_Private_Order_By>>;
  where?: InputMaybe<Auth_User_Private_Bool_Exp>;
};


export type Query_RootAuth_User_Private_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Auth_User_Private_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Auth_User_Private_Order_By>>;
  where?: InputMaybe<Auth_User_Private_Bool_Exp>;
};


export type Query_RootBlockArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


export type Query_RootBlock_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


export type Query_RootBlock_By_PkArgs = {
  id: Scalars['uuid'];
};


export type Query_RootBlock_CommentArgs = {
  distinct_on?: InputMaybe<Array<Block_Comment_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Comment_Order_By>>;
  where?: InputMaybe<Block_Comment_Bool_Exp>;
};


export type Query_RootBlock_Comment_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Comment_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Comment_Order_By>>;
  where?: InputMaybe<Block_Comment_Bool_Exp>;
};


export type Query_RootBlock_Comment_By_PkArgs = {
  id: Scalars['uuid'];
};


export type Query_RootBlock_FollowerArgs = {
  distinct_on?: InputMaybe<Array<Block_Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Follower_Order_By>>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};


export type Query_RootBlock_Follower_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Follower_Order_By>>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};


export type Query_RootBlock_Follower_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


export type Query_RootBlock_LikeArgs = {
  distinct_on?: InputMaybe<Array<Block_Like_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Like_Order_By>>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};


export type Query_RootBlock_Like_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Like_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Like_Order_By>>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};


export type Query_RootBlock_Like_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


export type Query_RootBlock_ReferenceArgs = {
  distinct_on?: InputMaybe<Array<Block_Reference_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Reference_Order_By>>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};


export type Query_RootBlock_Reference_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Reference_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Reference_Order_By>>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};


export type Query_RootBlock_Reference_By_PkArgs = {
  block_referee_id: Scalars['uuid'];
  block_referrer_id: Scalars['uuid'];
};


export type Query_RootBlock_RepostArgs = {
  distinct_on?: InputMaybe<Array<Block_Repost_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Repost_Order_By>>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};


export type Query_RootBlock_Repost_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Repost_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Repost_Order_By>>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};


export type Query_RootBlock_Repost_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


export type Query_RootFollowerArgs = {
  distinct_on?: InputMaybe<Array<Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Follower_Order_By>>;
  where?: InputMaybe<Follower_Bool_Exp>;
};


export type Query_RootFollower_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Follower_Order_By>>;
  where?: InputMaybe<Follower_Bool_Exp>;
};


export type Query_RootFollower_By_PkArgs = {
  follower_id: Scalars['uuid'];
  following_id: Scalars['uuid'];
};


export type Query_RootTimeline_EventArgs = {
  distinct_on?: InputMaybe<Array<Timeline_Event_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Timeline_Event_Order_By>>;
  where?: InputMaybe<Timeline_Event_Bool_Exp>;
};


export type Query_RootTimeline_Event_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Timeline_Event_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Timeline_Event_Order_By>>;
  where?: InputMaybe<Timeline_Event_Bool_Exp>;
};


export type Query_RootUsersArgs = {
  distinct_on?: InputMaybe<Array<Users_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Users_Order_By>>;
  where?: InputMaybe<Users_Bool_Exp>;
};


export type Query_RootUsers_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Users_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Users_Order_By>>;
  where?: InputMaybe<Users_Bool_Exp>;
};


export type Query_RootUsers_By_PkArgs = {
  id: Scalars['uuid'];
};


export type Query_RootWorkspaceArgs = {
  distinct_on?: InputMaybe<Array<Workspace_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Workspace_Order_By>>;
  where?: InputMaybe<Workspace_Bool_Exp>;
};


export type Query_RootWorkspace_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Workspace_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Workspace_Order_By>>;
  where?: InputMaybe<Workspace_Bool_Exp>;
};


export type Query_RootWorkspace_By_PkArgs = {
  id: Scalars['uuid'];
};

export type Save_Tags_Input = {
  block_id: Scalars['uuid'];
  value: Scalars['String'];
};

export type Save_Tags_Output = {
  __typename?: 'save_tags_output';
  msg?: Maybe<Scalars['String']>;
};

export type Subscription_Root = {
  __typename?: 'subscription_root';
  /** fetch data from the table: "auth.providers" using primary key columns */
  authProvider?: Maybe<AuthProviders>;
  /** fetch data from the table: "auth.provider_requests" using primary key columns */
  authProviderRequest?: Maybe<AuthProviderRequests>;
  /** fetch data from the table: "auth.provider_requests" */
  authProviderRequests: Array<AuthProviderRequests>;
  /** fetch aggregated fields from the table: "auth.provider_requests" */
  authProviderRequestsAggregate: AuthProviderRequests_Aggregate;
  /** fetch data from the table: "auth.providers" */
  authProviders: Array<AuthProviders>;
  /** fetch aggregated fields from the table: "auth.providers" */
  authProvidersAggregate: AuthProviders_Aggregate;
  /** fetch data from the table: "auth.refresh_tokens" using primary key columns */
  authRefreshToken?: Maybe<AuthRefreshTokens>;
  /** fetch data from the table: "auth.refresh_tokens" */
  authRefreshTokens: Array<AuthRefreshTokens>;
  /** fetch aggregated fields from the table: "auth.refresh_tokens" */
  authRefreshTokensAggregate: AuthRefreshTokens_Aggregate;
  /** fetch data from the table: "auth.roles" using primary key columns */
  authRole?: Maybe<AuthRoles>;
  /** fetch data from the table: "auth.roles" */
  authRoles: Array<AuthRoles>;
  /** fetch aggregated fields from the table: "auth.roles" */
  authRolesAggregate: AuthRoles_Aggregate;
  /** fetch data from the table: "auth.user_providers" using primary key columns */
  authUserProvider?: Maybe<AuthUserProviders>;
  /** fetch data from the table: "auth.user_providers" */
  authUserProviders: Array<AuthUserProviders>;
  /** fetch aggregated fields from the table: "auth.user_providers" */
  authUserProvidersAggregate: AuthUserProviders_Aggregate;
  /** fetch data from the table: "auth.user_roles" using primary key columns */
  authUserRole?: Maybe<AuthUserRoles>;
  /** fetch data from the table: "auth.user_roles" */
  authUserRoles: Array<AuthUserRoles>;
  /** fetch aggregated fields from the table: "auth.user_roles" */
  authUserRolesAggregate: AuthUserRoles_Aggregate;
  /** fetch data from the table: "auth.user_private" */
  auth_user_private: Array<Auth_User_Private>;
  /** fetch aggregated fields from the table: "auth.user_private" */
  auth_user_private_aggregate: Auth_User_Private_Aggregate;
  /** fetch data from the table: "block" */
  block: Array<Block>;
  /** fetch aggregated fields from the table: "block" */
  block_aggregate: Block_Aggregate;
  /** fetch data from the table: "block" using primary key columns */
  block_by_pk?: Maybe<Block>;
  /** fetch data from the table: "block_comment" */
  block_comment: Array<Block_Comment>;
  /** fetch aggregated fields from the table: "block_comment" */
  block_comment_aggregate: Block_Comment_Aggregate;
  /** fetch data from the table: "block_comment" using primary key columns */
  block_comment_by_pk?: Maybe<Block_Comment>;
  /** fetch data from the table: "block_follower" */
  block_follower: Array<Block_Follower>;
  /** fetch aggregated fields from the table: "block_follower" */
  block_follower_aggregate: Block_Follower_Aggregate;
  /** fetch data from the table: "block_follower" using primary key columns */
  block_follower_by_pk?: Maybe<Block_Follower>;
  /** fetch data from the table: "block_like" */
  block_like: Array<Block_Like>;
  /** fetch aggregated fields from the table: "block_like" */
  block_like_aggregate: Block_Like_Aggregate;
  /** fetch data from the table: "block_like" using primary key columns */
  block_like_by_pk?: Maybe<Block_Like>;
  /** fetch data from the table: "block_reference" */
  block_reference: Array<Block_Reference>;
  /** fetch aggregated fields from the table: "block_reference" */
  block_reference_aggregate: Block_Reference_Aggregate;
  /** fetch data from the table: "block_reference" using primary key columns */
  block_reference_by_pk?: Maybe<Block_Reference>;
  /** fetch data from the table: "block_repost" */
  block_repost: Array<Block_Repost>;
  /** fetch aggregated fields from the table: "block_repost" */
  block_repost_aggregate: Block_Repost_Aggregate;
  /** fetch data from the table: "block_repost" using primary key columns */
  block_repost_by_pk?: Maybe<Block_Repost>;
  /** fetch data from the table: "follower" */
  follower: Array<Follower>;
  /** fetch aggregated fields from the table: "follower" */
  follower_aggregate: Follower_Aggregate;
  /** fetch data from the table: "follower" using primary key columns */
  follower_by_pk?: Maybe<Follower>;
  /** fetch data from the table: "timeline_event" */
  timeline_event: Array<Timeline_Event>;
  /** fetch aggregated fields from the table: "timeline_event" */
  timeline_event_aggregate: Timeline_Event_Aggregate;
  /** fetch data from the table: "auth.users" */
  users: Array<Users>;
  /** fetch aggregated fields from the table: "auth.users" */
  users_aggregate: Users_Aggregate;
  /** fetch data from the table: "auth.users" using primary key columns */
  users_by_pk?: Maybe<Users>;
  /** fetch data from the table: "workspace" */
  workspace: Array<Workspace>;
  /** fetch aggregated fields from the table: "workspace" */
  workspace_aggregate: Workspace_Aggregate;
  /** fetch data from the table: "workspace" using primary key columns */
  workspace_by_pk?: Maybe<Workspace>;
};


export type Subscription_RootAuthProviderArgs = {
  id: Scalars['String'];
};


export type Subscription_RootAuthProviderRequestArgs = {
  id: Scalars['uuid'];
};


export type Subscription_RootAuthProviderRequestsArgs = {
  distinct_on?: InputMaybe<Array<AuthProviderRequests_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthProviderRequests_Order_By>>;
  where?: InputMaybe<AuthProviderRequests_Bool_Exp>;
};


export type Subscription_RootAuthProviderRequestsAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthProviderRequests_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthProviderRequests_Order_By>>;
  where?: InputMaybe<AuthProviderRequests_Bool_Exp>;
};


export type Subscription_RootAuthProvidersArgs = {
  distinct_on?: InputMaybe<Array<AuthProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthProviders_Order_By>>;
  where?: InputMaybe<AuthProviders_Bool_Exp>;
};


export type Subscription_RootAuthProvidersAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthProviders_Order_By>>;
  where?: InputMaybe<AuthProviders_Bool_Exp>;
};


export type Subscription_RootAuthRefreshTokenArgs = {
  refreshToken: Scalars['uuid'];
};


export type Subscription_RootAuthRefreshTokensArgs = {
  distinct_on?: InputMaybe<Array<AuthRefreshTokens_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRefreshTokens_Order_By>>;
  where?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
};


export type Subscription_RootAuthRefreshTokensAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthRefreshTokens_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRefreshTokens_Order_By>>;
  where?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
};


export type Subscription_RootAuthRoleArgs = {
  role: Scalars['String'];
};


export type Subscription_RootAuthRolesArgs = {
  distinct_on?: InputMaybe<Array<AuthRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRoles_Order_By>>;
  where?: InputMaybe<AuthRoles_Bool_Exp>;
};


export type Subscription_RootAuthRolesAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRoles_Order_By>>;
  where?: InputMaybe<AuthRoles_Bool_Exp>;
};


export type Subscription_RootAuthUserProviderArgs = {
  id: Scalars['uuid'];
};


export type Subscription_RootAuthUserProvidersArgs = {
  distinct_on?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserProviders_Order_By>>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};


export type Subscription_RootAuthUserProvidersAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserProviders_Order_By>>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};


export type Subscription_RootAuthUserRoleArgs = {
  id: Scalars['uuid'];
};


export type Subscription_RootAuthUserRolesArgs = {
  distinct_on?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserRoles_Order_By>>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};


export type Subscription_RootAuthUserRolesAggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserRoles_Order_By>>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};


export type Subscription_RootAuth_User_PrivateArgs = {
  distinct_on?: InputMaybe<Array<Auth_User_Private_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Auth_User_Private_Order_By>>;
  where?: InputMaybe<Auth_User_Private_Bool_Exp>;
};


export type Subscription_RootAuth_User_Private_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Auth_User_Private_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Auth_User_Private_Order_By>>;
  where?: InputMaybe<Auth_User_Private_Bool_Exp>;
};


export type Subscription_RootBlockArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


export type Subscription_RootBlock_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


export type Subscription_RootBlock_By_PkArgs = {
  id: Scalars['uuid'];
};


export type Subscription_RootBlock_CommentArgs = {
  distinct_on?: InputMaybe<Array<Block_Comment_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Comment_Order_By>>;
  where?: InputMaybe<Block_Comment_Bool_Exp>;
};


export type Subscription_RootBlock_Comment_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Comment_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Comment_Order_By>>;
  where?: InputMaybe<Block_Comment_Bool_Exp>;
};


export type Subscription_RootBlock_Comment_By_PkArgs = {
  id: Scalars['uuid'];
};


export type Subscription_RootBlock_FollowerArgs = {
  distinct_on?: InputMaybe<Array<Block_Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Follower_Order_By>>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};


export type Subscription_RootBlock_Follower_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Follower_Order_By>>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};


export type Subscription_RootBlock_Follower_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


export type Subscription_RootBlock_LikeArgs = {
  distinct_on?: InputMaybe<Array<Block_Like_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Like_Order_By>>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};


export type Subscription_RootBlock_Like_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Like_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Like_Order_By>>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};


export type Subscription_RootBlock_Like_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


export type Subscription_RootBlock_ReferenceArgs = {
  distinct_on?: InputMaybe<Array<Block_Reference_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Reference_Order_By>>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};


export type Subscription_RootBlock_Reference_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Reference_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Reference_Order_By>>;
  where?: InputMaybe<Block_Reference_Bool_Exp>;
};


export type Subscription_RootBlock_Reference_By_PkArgs = {
  block_referee_id: Scalars['uuid'];
  block_referrer_id: Scalars['uuid'];
};


export type Subscription_RootBlock_RepostArgs = {
  distinct_on?: InputMaybe<Array<Block_Repost_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Repost_Order_By>>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};


export type Subscription_RootBlock_Repost_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Repost_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Repost_Order_By>>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};


export type Subscription_RootBlock_Repost_By_PkArgs = {
  block_id: Scalars['uuid'];
  user_id: Scalars['uuid'];
};


export type Subscription_RootFollowerArgs = {
  distinct_on?: InputMaybe<Array<Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Follower_Order_By>>;
  where?: InputMaybe<Follower_Bool_Exp>;
};


export type Subscription_RootFollower_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Follower_Order_By>>;
  where?: InputMaybe<Follower_Bool_Exp>;
};


export type Subscription_RootFollower_By_PkArgs = {
  follower_id: Scalars['uuid'];
  following_id: Scalars['uuid'];
};


export type Subscription_RootTimeline_EventArgs = {
  distinct_on?: InputMaybe<Array<Timeline_Event_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Timeline_Event_Order_By>>;
  where?: InputMaybe<Timeline_Event_Bool_Exp>;
};


export type Subscription_RootTimeline_Event_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Timeline_Event_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Timeline_Event_Order_By>>;
  where?: InputMaybe<Timeline_Event_Bool_Exp>;
};


export type Subscription_RootUsersArgs = {
  distinct_on?: InputMaybe<Array<Users_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Users_Order_By>>;
  where?: InputMaybe<Users_Bool_Exp>;
};


export type Subscription_RootUsers_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Users_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Users_Order_By>>;
  where?: InputMaybe<Users_Bool_Exp>;
};


export type Subscription_RootUsers_By_PkArgs = {
  id: Scalars['uuid'];
};


export type Subscription_RootWorkspaceArgs = {
  distinct_on?: InputMaybe<Array<Workspace_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Workspace_Order_By>>;
  where?: InputMaybe<Workspace_Bool_Exp>;
};


export type Subscription_RootWorkspace_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Workspace_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Workspace_Order_By>>;
  where?: InputMaybe<Workspace_Bool_Exp>;
};


export type Subscription_RootWorkspace_By_PkArgs = {
  id: Scalars['uuid'];
};

/** columns and relationships of "timeline_event" */
export type Timeline_Event = {
  __typename?: 'timeline_event';
  created_at?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  type?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** aggregated selection of "timeline_event" */
export type Timeline_Event_Aggregate = {
  __typename?: 'timeline_event_aggregate';
  aggregate?: Maybe<Timeline_Event_Aggregate_Fields>;
  nodes: Array<Timeline_Event>;
};

/** aggregate fields of "timeline_event" */
export type Timeline_Event_Aggregate_Fields = {
  __typename?: 'timeline_event_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Timeline_Event_Max_Fields>;
  min?: Maybe<Timeline_Event_Min_Fields>;
};


/** aggregate fields of "timeline_event" */
export type Timeline_Event_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Timeline_Event_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** Boolean expression to filter rows from the table "timeline_event". All fields are combined with a logical 'AND'. */
export type Timeline_Event_Bool_Exp = {
  _and?: InputMaybe<Array<Timeline_Event_Bool_Exp>>;
  _not?: InputMaybe<Timeline_Event_Bool_Exp>;
  _or?: InputMaybe<Array<Timeline_Event_Bool_Exp>>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  type?: InputMaybe<String_Comparison_Exp>;
  user_id?: InputMaybe<Uuid_Comparison_Exp>;
};

/** aggregate max on columns */
export type Timeline_Event_Max_Fields = {
  __typename?: 'timeline_event_max_fields';
  created_at?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  type?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** aggregate min on columns */
export type Timeline_Event_Min_Fields = {
  __typename?: 'timeline_event_min_fields';
  created_at?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  type?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** Ordering options when selecting data from "timeline_event". */
export type Timeline_Event_Order_By = {
  created_at?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  type?: InputMaybe<Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** select columns of table "timeline_event" */
export enum Timeline_Event_Select_Column {
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  Id = 'id',
  /** column name */
  Type = 'type',
  /** column name */
  UserId = 'user_id'
}

/** Boolean expression to compare columns of type "timestamptz". All fields are combined with logical 'AND'. */
export type Timestamptz_Comparison_Exp = {
  _eq?: InputMaybe<Scalars['timestamptz']>;
  _gt?: InputMaybe<Scalars['timestamptz']>;
  _gte?: InputMaybe<Scalars['timestamptz']>;
  _in?: InputMaybe<Array<Scalars['timestamptz']>>;
  _is_null?: InputMaybe<Scalars['Boolean']>;
  _lt?: InputMaybe<Scalars['timestamptz']>;
  _lte?: InputMaybe<Scalars['timestamptz']>;
  _neq?: InputMaybe<Scalars['timestamptz']>;
  _nin?: InputMaybe<Array<Scalars['timestamptz']>>;
};

/** columns and relationships of "auth.users" */
export type Users = {
  __typename?: 'users';
  activeMfaType?: Maybe<Scalars['String']>;
  avatarUrl: Scalars['String'];
  banner_url?: Maybe<Scalars['String']>;
  bio?: Maybe<Scalars['String']>;
  birthdate?: Maybe<Scalars['date']>;
  /** An array relationship */
  block_comments: Array<Block_Comment>;
  /** An aggregate relationship */
  block_comments_aggregate: Block_Comment_Aggregate;
  /** An array relationship */
  block_followers: Array<Block_Follower>;
  /** An aggregate relationship */
  block_followers_aggregate: Block_Follower_Aggregate;
  /** An array relationship */
  block_likes: Array<Block_Like>;
  /** An aggregate relationship */
  block_likes_aggregate: Block_Like_Aggregate;
  /** An array relationship */
  block_reposts: Array<Block_Repost>;
  /** An aggregate relationship */
  block_reposts_aggregate: Block_Repost_Aggregate;
  /** An array relationship */
  blocks: Array<Block>;
  /** An aggregate relationship */
  blocks_aggregate: Block_Aggregate;
  createdAt: Scalars['timestamptz'];
  defaultRole: Scalars['String'];
  /** An object relationship */
  defaultRoleByRole: AuthRoles;
  disabled: Scalars['Boolean'];
  displayName: Scalars['String'];
  email?: Maybe<Scalars['citext']>;
  emailVerified: Scalars['Boolean'];
  facebook_id?: Maybe<Scalars['String']>;
  /** An array relationship */
  followers: Array<Follower>;
  /** An aggregate relationship */
  followers_aggregate: Follower_Aggregate;
  /** An array relationship */
  following: Array<Follower>;
  /** An aggregate relationship */
  following_aggregate: Follower_Aggregate;
  gender?: Maybe<Scalars['String']>;
  github_username?: Maybe<Scalars['String']>;
  google_id?: Maybe<Scalars['String']>;
  id: Scalars['uuid'];
  isAnonymous: Scalars['Boolean'];
  is_username_set?: Maybe<Scalars['Boolean']>;
  lastSeen?: Maybe<Scalars['timestamptz']>;
  locale: Scalars['String'];
  location?: Maybe<Scalars['String']>;
  metadata?: Maybe<Scalars['jsonb']>;
  newEmail?: Maybe<Scalars['citext']>;
  otpHash?: Maybe<Scalars['String']>;
  otpHashExpiresAt: Scalars['timestamptz'];
  otpMethodLastUsed?: Maybe<Scalars['String']>;
  passwordHash?: Maybe<Scalars['String']>;
  phoneNumber?: Maybe<Scalars['String']>;
  phoneNumberVerified: Scalars['Boolean'];
  /** An object relationship */
  private?: Maybe<Auth_User_Private>;
  /** An array relationship */
  refreshTokens: Array<AuthRefreshTokens>;
  /** An aggregate relationship */
  refreshTokens_aggregate: AuthRefreshTokens_Aggregate;
  /** An array relationship */
  roles: Array<AuthUserRoles>;
  /** An aggregate relationship */
  roles_aggregate: AuthUserRoles_Aggregate;
  ticket?: Maybe<Scalars['String']>;
  ticketExpiresAt: Scalars['timestamptz'];
  totpSecret?: Maybe<Scalars['String']>;
  twitter_username?: Maybe<Scalars['String']>;
  updatedAt: Scalars['timestamptz'];
  /** An array relationship */
  userProviders: Array<AuthUserProviders>;
  /** An aggregate relationship */
  userProviders_aggregate: AuthUserProviders_Aggregate;
  username?: Maybe<Scalars['String']>;
  website_url?: Maybe<Scalars['String']>;
  /** An object relationship */
  workspace?: Maybe<Workspace>;
};


/** columns and relationships of "auth.users" */
export type UsersBlock_CommentsArgs = {
  distinct_on?: InputMaybe<Array<Block_Comment_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Comment_Order_By>>;
  where?: InputMaybe<Block_Comment_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlock_Comments_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Comment_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Comment_Order_By>>;
  where?: InputMaybe<Block_Comment_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlock_FollowersArgs = {
  distinct_on?: InputMaybe<Array<Block_Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Follower_Order_By>>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlock_Followers_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Follower_Order_By>>;
  where?: InputMaybe<Block_Follower_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlock_LikesArgs = {
  distinct_on?: InputMaybe<Array<Block_Like_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Like_Order_By>>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlock_Likes_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Like_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Like_Order_By>>;
  where?: InputMaybe<Block_Like_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlock_RepostsArgs = {
  distinct_on?: InputMaybe<Array<Block_Repost_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Repost_Order_By>>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlock_Reposts_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Repost_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Repost_Order_By>>;
  where?: InputMaybe<Block_Repost_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlocksArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersBlocks_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Block_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Block_Order_By>>;
  where?: InputMaybe<Block_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersFollowersArgs = {
  distinct_on?: InputMaybe<Array<Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Follower_Order_By>>;
  where?: InputMaybe<Follower_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersFollowers_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Follower_Order_By>>;
  where?: InputMaybe<Follower_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersFollowingArgs = {
  distinct_on?: InputMaybe<Array<Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Follower_Order_By>>;
  where?: InputMaybe<Follower_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersFollowing_AggregateArgs = {
  distinct_on?: InputMaybe<Array<Follower_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<Follower_Order_By>>;
  where?: InputMaybe<Follower_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersMetadataArgs = {
  path?: InputMaybe<Scalars['String']>;
};


/** columns and relationships of "auth.users" */
export type UsersRefreshTokensArgs = {
  distinct_on?: InputMaybe<Array<AuthRefreshTokens_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRefreshTokens_Order_By>>;
  where?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersRefreshTokens_AggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthRefreshTokens_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthRefreshTokens_Order_By>>;
  where?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersRolesArgs = {
  distinct_on?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserRoles_Order_By>>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersRoles_AggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthUserRoles_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserRoles_Order_By>>;
  where?: InputMaybe<AuthUserRoles_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersUserProvidersArgs = {
  distinct_on?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserProviders_Order_By>>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};


/** columns and relationships of "auth.users" */
export type UsersUserProviders_AggregateArgs = {
  distinct_on?: InputMaybe<Array<AuthUserProviders_Select_Column>>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  order_by?: InputMaybe<Array<AuthUserProviders_Order_By>>;
  where?: InputMaybe<AuthUserProviders_Bool_Exp>;
};

/** aggregated selection of "auth.users" */
export type Users_Aggregate = {
  __typename?: 'users_aggregate';
  aggregate?: Maybe<Users_Aggregate_Fields>;
  nodes: Array<Users>;
};

/** aggregate fields of "auth.users" */
export type Users_Aggregate_Fields = {
  __typename?: 'users_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Users_Max_Fields>;
  min?: Maybe<Users_Min_Fields>;
};


/** aggregate fields of "auth.users" */
export type Users_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Users_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "auth.users" */
export type Users_Aggregate_Order_By = {
  count?: InputMaybe<Order_By>;
  max?: InputMaybe<Users_Max_Order_By>;
  min?: InputMaybe<Users_Min_Order_By>;
};

/** append existing jsonb value of filtered columns with new jsonb value */
export type Users_Append_Input = {
  metadata?: InputMaybe<Scalars['jsonb']>;
};

/** input type for inserting array relation for remote table "auth.users" */
export type Users_Arr_Rel_Insert_Input = {
  data: Array<Users_Insert_Input>;
  /** on conflict condition */
  on_conflict?: InputMaybe<Users_On_Conflict>;
};

/** Boolean expression to filter rows from the table "auth.users". All fields are combined with a logical 'AND'. */
export type Users_Bool_Exp = {
  _and?: InputMaybe<Array<Users_Bool_Exp>>;
  _not?: InputMaybe<Users_Bool_Exp>;
  _or?: InputMaybe<Array<Users_Bool_Exp>>;
  activeMfaType?: InputMaybe<String_Comparison_Exp>;
  avatarUrl?: InputMaybe<String_Comparison_Exp>;
  banner_url?: InputMaybe<String_Comparison_Exp>;
  bio?: InputMaybe<String_Comparison_Exp>;
  birthdate?: InputMaybe<Date_Comparison_Exp>;
  block_comments?: InputMaybe<Block_Comment_Bool_Exp>;
  block_followers?: InputMaybe<Block_Follower_Bool_Exp>;
  block_likes?: InputMaybe<Block_Like_Bool_Exp>;
  block_reposts?: InputMaybe<Block_Repost_Bool_Exp>;
  blocks?: InputMaybe<Block_Bool_Exp>;
  createdAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  defaultRole?: InputMaybe<String_Comparison_Exp>;
  defaultRoleByRole?: InputMaybe<AuthRoles_Bool_Exp>;
  disabled?: InputMaybe<Boolean_Comparison_Exp>;
  displayName?: InputMaybe<String_Comparison_Exp>;
  email?: InputMaybe<Citext_Comparison_Exp>;
  emailVerified?: InputMaybe<Boolean_Comparison_Exp>;
  facebook_id?: InputMaybe<String_Comparison_Exp>;
  followers?: InputMaybe<Follower_Bool_Exp>;
  following?: InputMaybe<Follower_Bool_Exp>;
  gender?: InputMaybe<String_Comparison_Exp>;
  github_username?: InputMaybe<String_Comparison_Exp>;
  google_id?: InputMaybe<String_Comparison_Exp>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  isAnonymous?: InputMaybe<Boolean_Comparison_Exp>;
  is_username_set?: InputMaybe<Boolean_Comparison_Exp>;
  lastSeen?: InputMaybe<Timestamptz_Comparison_Exp>;
  locale?: InputMaybe<String_Comparison_Exp>;
  location?: InputMaybe<String_Comparison_Exp>;
  metadata?: InputMaybe<Jsonb_Comparison_Exp>;
  newEmail?: InputMaybe<Citext_Comparison_Exp>;
  otpHash?: InputMaybe<String_Comparison_Exp>;
  otpHashExpiresAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  otpMethodLastUsed?: InputMaybe<String_Comparison_Exp>;
  passwordHash?: InputMaybe<String_Comparison_Exp>;
  phoneNumber?: InputMaybe<String_Comparison_Exp>;
  phoneNumberVerified?: InputMaybe<Boolean_Comparison_Exp>;
  private?: InputMaybe<Auth_User_Private_Bool_Exp>;
  refreshTokens?: InputMaybe<AuthRefreshTokens_Bool_Exp>;
  roles?: InputMaybe<AuthUserRoles_Bool_Exp>;
  ticket?: InputMaybe<String_Comparison_Exp>;
  ticketExpiresAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  totpSecret?: InputMaybe<String_Comparison_Exp>;
  twitter_username?: InputMaybe<String_Comparison_Exp>;
  updatedAt?: InputMaybe<Timestamptz_Comparison_Exp>;
  userProviders?: InputMaybe<AuthUserProviders_Bool_Exp>;
  username?: InputMaybe<String_Comparison_Exp>;
  website_url?: InputMaybe<String_Comparison_Exp>;
  workspace?: InputMaybe<Workspace_Bool_Exp>;
};

/** unique or primary key constraints on table "auth.users" */
export enum Users_Constraint {
  /** unique or primary key constraint */
  UsersEmailKey = 'users_email_key',
  /** unique or primary key constraint */
  UsersPhoneNumberKey = 'users_phone_number_key',
  /** unique or primary key constraint */
  UsersPkey = 'users_pkey',
  /** unique or primary key constraint */
  UsersUsernameKey = 'users_username_key'
}

/** delete the field or element with specified path (for JSON arrays, negative integers count from the end) */
export type Users_Delete_At_Path_Input = {
  metadata?: InputMaybe<Array<Scalars['String']>>;
};

/** delete the array element with specified index (negative integers count from the end). throws an error if top level container is not an array */
export type Users_Delete_Elem_Input = {
  metadata?: InputMaybe<Scalars['Int']>;
};

/** delete key/value pair or string element. key/value pairs are matched based on their key value */
export type Users_Delete_Key_Input = {
  metadata?: InputMaybe<Scalars['String']>;
};

/** input type for inserting data into table "auth.users" */
export type Users_Insert_Input = {
  activeMfaType?: InputMaybe<Scalars['String']>;
  avatarUrl?: InputMaybe<Scalars['String']>;
  banner_url?: InputMaybe<Scalars['String']>;
  bio?: InputMaybe<Scalars['String']>;
  birthdate?: InputMaybe<Scalars['date']>;
  block_comments?: InputMaybe<Block_Comment_Arr_Rel_Insert_Input>;
  block_followers?: InputMaybe<Block_Follower_Arr_Rel_Insert_Input>;
  block_likes?: InputMaybe<Block_Like_Arr_Rel_Insert_Input>;
  block_reposts?: InputMaybe<Block_Repost_Arr_Rel_Insert_Input>;
  blocks?: InputMaybe<Block_Arr_Rel_Insert_Input>;
  createdAt?: InputMaybe<Scalars['timestamptz']>;
  defaultRole?: InputMaybe<Scalars['String']>;
  defaultRoleByRole?: InputMaybe<AuthRoles_Obj_Rel_Insert_Input>;
  disabled?: InputMaybe<Scalars['Boolean']>;
  displayName?: InputMaybe<Scalars['String']>;
  email?: InputMaybe<Scalars['citext']>;
  emailVerified?: InputMaybe<Scalars['Boolean']>;
  facebook_id?: InputMaybe<Scalars['String']>;
  followers?: InputMaybe<Follower_Arr_Rel_Insert_Input>;
  following?: InputMaybe<Follower_Arr_Rel_Insert_Input>;
  gender?: InputMaybe<Scalars['String']>;
  github_username?: InputMaybe<Scalars['String']>;
  google_id?: InputMaybe<Scalars['String']>;
  id?: InputMaybe<Scalars['uuid']>;
  isAnonymous?: InputMaybe<Scalars['Boolean']>;
  is_username_set?: InputMaybe<Scalars['Boolean']>;
  lastSeen?: InputMaybe<Scalars['timestamptz']>;
  locale?: InputMaybe<Scalars['String']>;
  location?: InputMaybe<Scalars['String']>;
  metadata?: InputMaybe<Scalars['jsonb']>;
  newEmail?: InputMaybe<Scalars['citext']>;
  otpHash?: InputMaybe<Scalars['String']>;
  otpHashExpiresAt?: InputMaybe<Scalars['timestamptz']>;
  otpMethodLastUsed?: InputMaybe<Scalars['String']>;
  passwordHash?: InputMaybe<Scalars['String']>;
  phoneNumber?: InputMaybe<Scalars['String']>;
  phoneNumberVerified?: InputMaybe<Scalars['Boolean']>;
  private?: InputMaybe<Auth_User_Private_Obj_Rel_Insert_Input>;
  refreshTokens?: InputMaybe<AuthRefreshTokens_Arr_Rel_Insert_Input>;
  roles?: InputMaybe<AuthUserRoles_Arr_Rel_Insert_Input>;
  ticket?: InputMaybe<Scalars['String']>;
  ticketExpiresAt?: InputMaybe<Scalars['timestamptz']>;
  totpSecret?: InputMaybe<Scalars['String']>;
  twitter_username?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['timestamptz']>;
  userProviders?: InputMaybe<AuthUserProviders_Arr_Rel_Insert_Input>;
  username?: InputMaybe<Scalars['String']>;
  website_url?: InputMaybe<Scalars['String']>;
  workspace?: InputMaybe<Workspace_Obj_Rel_Insert_Input>;
};

/** aggregate max on columns */
export type Users_Max_Fields = {
  __typename?: 'users_max_fields';
  activeMfaType?: Maybe<Scalars['String']>;
  avatarUrl?: Maybe<Scalars['String']>;
  banner_url?: Maybe<Scalars['String']>;
  bio?: Maybe<Scalars['String']>;
  birthdate?: Maybe<Scalars['date']>;
  createdAt?: Maybe<Scalars['timestamptz']>;
  defaultRole?: Maybe<Scalars['String']>;
  displayName?: Maybe<Scalars['String']>;
  email?: Maybe<Scalars['citext']>;
  facebook_id?: Maybe<Scalars['String']>;
  gender?: Maybe<Scalars['String']>;
  github_username?: Maybe<Scalars['String']>;
  google_id?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['uuid']>;
  lastSeen?: Maybe<Scalars['timestamptz']>;
  locale?: Maybe<Scalars['String']>;
  location?: Maybe<Scalars['String']>;
  newEmail?: Maybe<Scalars['citext']>;
  otpHash?: Maybe<Scalars['String']>;
  otpHashExpiresAt?: Maybe<Scalars['timestamptz']>;
  otpMethodLastUsed?: Maybe<Scalars['String']>;
  passwordHash?: Maybe<Scalars['String']>;
  phoneNumber?: Maybe<Scalars['String']>;
  ticket?: Maybe<Scalars['String']>;
  ticketExpiresAt?: Maybe<Scalars['timestamptz']>;
  totpSecret?: Maybe<Scalars['String']>;
  twitter_username?: Maybe<Scalars['String']>;
  updatedAt?: Maybe<Scalars['timestamptz']>;
  username?: Maybe<Scalars['String']>;
  website_url?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "auth.users" */
export type Users_Max_Order_By = {
  activeMfaType?: InputMaybe<Order_By>;
  avatarUrl?: InputMaybe<Order_By>;
  banner_url?: InputMaybe<Order_By>;
  bio?: InputMaybe<Order_By>;
  birthdate?: InputMaybe<Order_By>;
  createdAt?: InputMaybe<Order_By>;
  defaultRole?: InputMaybe<Order_By>;
  displayName?: InputMaybe<Order_By>;
  email?: InputMaybe<Order_By>;
  facebook_id?: InputMaybe<Order_By>;
  gender?: InputMaybe<Order_By>;
  github_username?: InputMaybe<Order_By>;
  google_id?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  lastSeen?: InputMaybe<Order_By>;
  locale?: InputMaybe<Order_By>;
  location?: InputMaybe<Order_By>;
  newEmail?: InputMaybe<Order_By>;
  otpHash?: InputMaybe<Order_By>;
  otpHashExpiresAt?: InputMaybe<Order_By>;
  otpMethodLastUsed?: InputMaybe<Order_By>;
  passwordHash?: InputMaybe<Order_By>;
  phoneNumber?: InputMaybe<Order_By>;
  ticket?: InputMaybe<Order_By>;
  ticketExpiresAt?: InputMaybe<Order_By>;
  totpSecret?: InputMaybe<Order_By>;
  twitter_username?: InputMaybe<Order_By>;
  updatedAt?: InputMaybe<Order_By>;
  username?: InputMaybe<Order_By>;
  website_url?: InputMaybe<Order_By>;
};

/** aggregate min on columns */
export type Users_Min_Fields = {
  __typename?: 'users_min_fields';
  activeMfaType?: Maybe<Scalars['String']>;
  avatarUrl?: Maybe<Scalars['String']>;
  banner_url?: Maybe<Scalars['String']>;
  bio?: Maybe<Scalars['String']>;
  birthdate?: Maybe<Scalars['date']>;
  createdAt?: Maybe<Scalars['timestamptz']>;
  defaultRole?: Maybe<Scalars['String']>;
  displayName?: Maybe<Scalars['String']>;
  email?: Maybe<Scalars['citext']>;
  facebook_id?: Maybe<Scalars['String']>;
  gender?: Maybe<Scalars['String']>;
  github_username?: Maybe<Scalars['String']>;
  google_id?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['uuid']>;
  lastSeen?: Maybe<Scalars['timestamptz']>;
  locale?: Maybe<Scalars['String']>;
  location?: Maybe<Scalars['String']>;
  newEmail?: Maybe<Scalars['citext']>;
  otpHash?: Maybe<Scalars['String']>;
  otpHashExpiresAt?: Maybe<Scalars['timestamptz']>;
  otpMethodLastUsed?: Maybe<Scalars['String']>;
  passwordHash?: Maybe<Scalars['String']>;
  phoneNumber?: Maybe<Scalars['String']>;
  ticket?: Maybe<Scalars['String']>;
  ticketExpiresAt?: Maybe<Scalars['timestamptz']>;
  totpSecret?: Maybe<Scalars['String']>;
  twitter_username?: Maybe<Scalars['String']>;
  updatedAt?: Maybe<Scalars['timestamptz']>;
  username?: Maybe<Scalars['String']>;
  website_url?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "auth.users" */
export type Users_Min_Order_By = {
  activeMfaType?: InputMaybe<Order_By>;
  avatarUrl?: InputMaybe<Order_By>;
  banner_url?: InputMaybe<Order_By>;
  bio?: InputMaybe<Order_By>;
  birthdate?: InputMaybe<Order_By>;
  createdAt?: InputMaybe<Order_By>;
  defaultRole?: InputMaybe<Order_By>;
  displayName?: InputMaybe<Order_By>;
  email?: InputMaybe<Order_By>;
  facebook_id?: InputMaybe<Order_By>;
  gender?: InputMaybe<Order_By>;
  github_username?: InputMaybe<Order_By>;
  google_id?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  lastSeen?: InputMaybe<Order_By>;
  locale?: InputMaybe<Order_By>;
  location?: InputMaybe<Order_By>;
  newEmail?: InputMaybe<Order_By>;
  otpHash?: InputMaybe<Order_By>;
  otpHashExpiresAt?: InputMaybe<Order_By>;
  otpMethodLastUsed?: InputMaybe<Order_By>;
  passwordHash?: InputMaybe<Order_By>;
  phoneNumber?: InputMaybe<Order_By>;
  ticket?: InputMaybe<Order_By>;
  ticketExpiresAt?: InputMaybe<Order_By>;
  totpSecret?: InputMaybe<Order_By>;
  twitter_username?: InputMaybe<Order_By>;
  updatedAt?: InputMaybe<Order_By>;
  username?: InputMaybe<Order_By>;
  website_url?: InputMaybe<Order_By>;
};

/** response of any mutation on the table "auth.users" */
export type Users_Mutation_Response = {
  __typename?: 'users_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Users>;
};

/** input type for inserting object relation for remote table "auth.users" */
export type Users_Obj_Rel_Insert_Input = {
  data: Users_Insert_Input;
  /** on conflict condition */
  on_conflict?: InputMaybe<Users_On_Conflict>;
};

/** on conflict condition type for table "auth.users" */
export type Users_On_Conflict = {
  constraint: Users_Constraint;
  update_columns?: Array<Users_Update_Column>;
  where?: InputMaybe<Users_Bool_Exp>;
};

/** Ordering options when selecting data from "auth.users". */
export type Users_Order_By = {
  activeMfaType?: InputMaybe<Order_By>;
  avatarUrl?: InputMaybe<Order_By>;
  banner_url?: InputMaybe<Order_By>;
  bio?: InputMaybe<Order_By>;
  birthdate?: InputMaybe<Order_By>;
  block_comments_aggregate?: InputMaybe<Block_Comment_Aggregate_Order_By>;
  block_followers_aggregate?: InputMaybe<Block_Follower_Aggregate_Order_By>;
  block_likes_aggregate?: InputMaybe<Block_Like_Aggregate_Order_By>;
  block_reposts_aggregate?: InputMaybe<Block_Repost_Aggregate_Order_By>;
  blocks_aggregate?: InputMaybe<Block_Aggregate_Order_By>;
  createdAt?: InputMaybe<Order_By>;
  defaultRole?: InputMaybe<Order_By>;
  defaultRoleByRole?: InputMaybe<AuthRoles_Order_By>;
  disabled?: InputMaybe<Order_By>;
  displayName?: InputMaybe<Order_By>;
  email?: InputMaybe<Order_By>;
  emailVerified?: InputMaybe<Order_By>;
  facebook_id?: InputMaybe<Order_By>;
  followers_aggregate?: InputMaybe<Follower_Aggregate_Order_By>;
  following_aggregate?: InputMaybe<Follower_Aggregate_Order_By>;
  gender?: InputMaybe<Order_By>;
  github_username?: InputMaybe<Order_By>;
  google_id?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  isAnonymous?: InputMaybe<Order_By>;
  is_username_set?: InputMaybe<Order_By>;
  lastSeen?: InputMaybe<Order_By>;
  locale?: InputMaybe<Order_By>;
  location?: InputMaybe<Order_By>;
  metadata?: InputMaybe<Order_By>;
  newEmail?: InputMaybe<Order_By>;
  otpHash?: InputMaybe<Order_By>;
  otpHashExpiresAt?: InputMaybe<Order_By>;
  otpMethodLastUsed?: InputMaybe<Order_By>;
  passwordHash?: InputMaybe<Order_By>;
  phoneNumber?: InputMaybe<Order_By>;
  phoneNumberVerified?: InputMaybe<Order_By>;
  private?: InputMaybe<Auth_User_Private_Order_By>;
  refreshTokens_aggregate?: InputMaybe<AuthRefreshTokens_Aggregate_Order_By>;
  roles_aggregate?: InputMaybe<AuthUserRoles_Aggregate_Order_By>;
  ticket?: InputMaybe<Order_By>;
  ticketExpiresAt?: InputMaybe<Order_By>;
  totpSecret?: InputMaybe<Order_By>;
  twitter_username?: InputMaybe<Order_By>;
  updatedAt?: InputMaybe<Order_By>;
  userProviders_aggregate?: InputMaybe<AuthUserProviders_Aggregate_Order_By>;
  username?: InputMaybe<Order_By>;
  website_url?: InputMaybe<Order_By>;
  workspace?: InputMaybe<Workspace_Order_By>;
};

/** primary key columns input for table: users */
export type Users_Pk_Columns_Input = {
  id: Scalars['uuid'];
};

/** prepend existing jsonb value of filtered columns with new jsonb value */
export type Users_Prepend_Input = {
  metadata?: InputMaybe<Scalars['jsonb']>;
};

/** select columns of table "auth.users" */
export enum Users_Select_Column {
  /** column name */
  ActiveMfaType = 'activeMfaType',
  /** column name */
  AvatarUrl = 'avatarUrl',
  /** column name */
  BannerUrl = 'banner_url',
  /** column name */
  Bio = 'bio',
  /** column name */
  Birthdate = 'birthdate',
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  DefaultRole = 'defaultRole',
  /** column name */
  Disabled = 'disabled',
  /** column name */
  DisplayName = 'displayName',
  /** column name */
  Email = 'email',
  /** column name */
  EmailVerified = 'emailVerified',
  /** column name */
  FacebookId = 'facebook_id',
  /** column name */
  Gender = 'gender',
  /** column name */
  GithubUsername = 'github_username',
  /** column name */
  GoogleId = 'google_id',
  /** column name */
  Id = 'id',
  /** column name */
  IsAnonymous = 'isAnonymous',
  /** column name */
  IsUsernameSet = 'is_username_set',
  /** column name */
  LastSeen = 'lastSeen',
  /** column name */
  Locale = 'locale',
  /** column name */
  Location = 'location',
  /** column name */
  Metadata = 'metadata',
  /** column name */
  NewEmail = 'newEmail',
  /** column name */
  OtpHash = 'otpHash',
  /** column name */
  OtpHashExpiresAt = 'otpHashExpiresAt',
  /** column name */
  OtpMethodLastUsed = 'otpMethodLastUsed',
  /** column name */
  PasswordHash = 'passwordHash',
  /** column name */
  PhoneNumber = 'phoneNumber',
  /** column name */
  PhoneNumberVerified = 'phoneNumberVerified',
  /** column name */
  Ticket = 'ticket',
  /** column name */
  TicketExpiresAt = 'ticketExpiresAt',
  /** column name */
  TotpSecret = 'totpSecret',
  /** column name */
  TwitterUsername = 'twitter_username',
  /** column name */
  UpdatedAt = 'updatedAt',
  /** column name */
  Username = 'username',
  /** column name */
  WebsiteUrl = 'website_url'
}

/** input type for updating data in table "auth.users" */
export type Users_Set_Input = {
  activeMfaType?: InputMaybe<Scalars['String']>;
  avatarUrl?: InputMaybe<Scalars['String']>;
  banner_url?: InputMaybe<Scalars['String']>;
  bio?: InputMaybe<Scalars['String']>;
  birthdate?: InputMaybe<Scalars['date']>;
  createdAt?: InputMaybe<Scalars['timestamptz']>;
  defaultRole?: InputMaybe<Scalars['String']>;
  disabled?: InputMaybe<Scalars['Boolean']>;
  displayName?: InputMaybe<Scalars['String']>;
  email?: InputMaybe<Scalars['citext']>;
  emailVerified?: InputMaybe<Scalars['Boolean']>;
  facebook_id?: InputMaybe<Scalars['String']>;
  gender?: InputMaybe<Scalars['String']>;
  github_username?: InputMaybe<Scalars['String']>;
  google_id?: InputMaybe<Scalars['String']>;
  id?: InputMaybe<Scalars['uuid']>;
  isAnonymous?: InputMaybe<Scalars['Boolean']>;
  is_username_set?: InputMaybe<Scalars['Boolean']>;
  lastSeen?: InputMaybe<Scalars['timestamptz']>;
  locale?: InputMaybe<Scalars['String']>;
  location?: InputMaybe<Scalars['String']>;
  metadata?: InputMaybe<Scalars['jsonb']>;
  newEmail?: InputMaybe<Scalars['citext']>;
  otpHash?: InputMaybe<Scalars['String']>;
  otpHashExpiresAt?: InputMaybe<Scalars['timestamptz']>;
  otpMethodLastUsed?: InputMaybe<Scalars['String']>;
  passwordHash?: InputMaybe<Scalars['String']>;
  phoneNumber?: InputMaybe<Scalars['String']>;
  phoneNumberVerified?: InputMaybe<Scalars['Boolean']>;
  ticket?: InputMaybe<Scalars['String']>;
  ticketExpiresAt?: InputMaybe<Scalars['timestamptz']>;
  totpSecret?: InputMaybe<Scalars['String']>;
  twitter_username?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['timestamptz']>;
  username?: InputMaybe<Scalars['String']>;
  website_url?: InputMaybe<Scalars['String']>;
};

/** update columns of table "auth.users" */
export enum Users_Update_Column {
  /** column name */
  ActiveMfaType = 'activeMfaType',
  /** column name */
  AvatarUrl = 'avatarUrl',
  /** column name */
  BannerUrl = 'banner_url',
  /** column name */
  Bio = 'bio',
  /** column name */
  Birthdate = 'birthdate',
  /** column name */
  CreatedAt = 'createdAt',
  /** column name */
  DefaultRole = 'defaultRole',
  /** column name */
  Disabled = 'disabled',
  /** column name */
  DisplayName = 'displayName',
  /** column name */
  Email = 'email',
  /** column name */
  EmailVerified = 'emailVerified',
  /** column name */
  FacebookId = 'facebook_id',
  /** column name */
  Gender = 'gender',
  /** column name */
  GithubUsername = 'github_username',
  /** column name */
  GoogleId = 'google_id',
  /** column name */
  Id = 'id',
  /** column name */
  IsAnonymous = 'isAnonymous',
  /** column name */
  IsUsernameSet = 'is_username_set',
  /** column name */
  LastSeen = 'lastSeen',
  /** column name */
  Locale = 'locale',
  /** column name */
  Location = 'location',
  /** column name */
  Metadata = 'metadata',
  /** column name */
  NewEmail = 'newEmail',
  /** column name */
  OtpHash = 'otpHash',
  /** column name */
  OtpHashExpiresAt = 'otpHashExpiresAt',
  /** column name */
  OtpMethodLastUsed = 'otpMethodLastUsed',
  /** column name */
  PasswordHash = 'passwordHash',
  /** column name */
  PhoneNumber = 'phoneNumber',
  /** column name */
  PhoneNumberVerified = 'phoneNumberVerified',
  /** column name */
  Ticket = 'ticket',
  /** column name */
  TicketExpiresAt = 'ticketExpiresAt',
  /** column name */
  TotpSecret = 'totpSecret',
  /** column name */
  TwitterUsername = 'twitter_username',
  /** column name */
  UpdatedAt = 'updatedAt',
  /** column name */
  Username = 'username',
  /** column name */
  WebsiteUrl = 'website_url'
}

/** Boolean expression to compare columns of type "uuid". All fields are combined with logical 'AND'. */
export type Uuid_Comparison_Exp = {
  _eq?: InputMaybe<Scalars['uuid']>;
  _gt?: InputMaybe<Scalars['uuid']>;
  _gte?: InputMaybe<Scalars['uuid']>;
  _in?: InputMaybe<Array<Scalars['uuid']>>;
  _is_null?: InputMaybe<Scalars['Boolean']>;
  _lt?: InputMaybe<Scalars['uuid']>;
  _lte?: InputMaybe<Scalars['uuid']>;
  _neq?: InputMaybe<Scalars['uuid']>;
  _nin?: InputMaybe<Array<Scalars['uuid']>>;
};

/** columns and relationships of "workspace" */
export type Workspace = {
  __typename?: 'workspace';
  created_at: Scalars['timestamptz'];
  id: Scalars['uuid'];
  name: Scalars['String'];
  updated_at: Scalars['timestamptz'];
  /** An object relationship */
  user: Users;
  user_id: Scalars['uuid'];
};

/** aggregated selection of "workspace" */
export type Workspace_Aggregate = {
  __typename?: 'workspace_aggregate';
  aggregate?: Maybe<Workspace_Aggregate_Fields>;
  nodes: Array<Workspace>;
};

/** aggregate fields of "workspace" */
export type Workspace_Aggregate_Fields = {
  __typename?: 'workspace_aggregate_fields';
  count: Scalars['Int'];
  max?: Maybe<Workspace_Max_Fields>;
  min?: Maybe<Workspace_Min_Fields>;
};


/** aggregate fields of "workspace" */
export type Workspace_Aggregate_FieldsCountArgs = {
  columns?: InputMaybe<Array<Workspace_Select_Column>>;
  distinct?: InputMaybe<Scalars['Boolean']>;
};

/** Boolean expression to filter rows from the table "workspace". All fields are combined with a logical 'AND'. */
export type Workspace_Bool_Exp = {
  _and?: InputMaybe<Array<Workspace_Bool_Exp>>;
  _not?: InputMaybe<Workspace_Bool_Exp>;
  _or?: InputMaybe<Array<Workspace_Bool_Exp>>;
  created_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  id?: InputMaybe<Uuid_Comparison_Exp>;
  name?: InputMaybe<String_Comparison_Exp>;
  updated_at?: InputMaybe<Timestamptz_Comparison_Exp>;
  user?: InputMaybe<Users_Bool_Exp>;
  user_id?: InputMaybe<Uuid_Comparison_Exp>;
};

/** unique or primary key constraints on table "workspace" */
export enum Workspace_Constraint {
  /** unique or primary key constraint */
  WorkspacePkey = 'workspace_pkey'
}

/** input type for inserting data into table "workspace" */
export type Workspace_Insert_Input = {
  created_at?: InputMaybe<Scalars['timestamptz']>;
  id?: InputMaybe<Scalars['uuid']>;
  name?: InputMaybe<Scalars['String']>;
  updated_at?: InputMaybe<Scalars['timestamptz']>;
  user?: InputMaybe<Users_Obj_Rel_Insert_Input>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** aggregate max on columns */
export type Workspace_Max_Fields = {
  __typename?: 'workspace_max_fields';
  created_at?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  name?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamptz']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** aggregate min on columns */
export type Workspace_Min_Fields = {
  __typename?: 'workspace_min_fields';
  created_at?: Maybe<Scalars['timestamptz']>;
  id?: Maybe<Scalars['uuid']>;
  name?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamptz']>;
  user_id?: Maybe<Scalars['uuid']>;
};

/** response of any mutation on the table "workspace" */
export type Workspace_Mutation_Response = {
  __typename?: 'workspace_mutation_response';
  /** number of rows affected by the mutation */
  affected_rows: Scalars['Int'];
  /** data from the rows affected by the mutation */
  returning: Array<Workspace>;
};

/** input type for inserting object relation for remote table "workspace" */
export type Workspace_Obj_Rel_Insert_Input = {
  data: Workspace_Insert_Input;
  /** on conflict condition */
  on_conflict?: InputMaybe<Workspace_On_Conflict>;
};

/** on conflict condition type for table "workspace" */
export type Workspace_On_Conflict = {
  constraint: Workspace_Constraint;
  update_columns?: Array<Workspace_Update_Column>;
  where?: InputMaybe<Workspace_Bool_Exp>;
};

/** Ordering options when selecting data from "workspace". */
export type Workspace_Order_By = {
  created_at?: InputMaybe<Order_By>;
  id?: InputMaybe<Order_By>;
  name?: InputMaybe<Order_By>;
  updated_at?: InputMaybe<Order_By>;
  user?: InputMaybe<Users_Order_By>;
  user_id?: InputMaybe<Order_By>;
};

/** primary key columns input for table: workspace */
export type Workspace_Pk_Columns_Input = {
  id: Scalars['uuid'];
};

/** select columns of table "workspace" */
export enum Workspace_Select_Column {
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  Id = 'id',
  /** column name */
  Name = 'name',
  /** column name */
  UpdatedAt = 'updated_at',
  /** column name */
  UserId = 'user_id'
}

/** input type for updating data in table "workspace" */
export type Workspace_Set_Input = {
  created_at?: InputMaybe<Scalars['timestamptz']>;
  id?: InputMaybe<Scalars['uuid']>;
  name?: InputMaybe<Scalars['String']>;
  updated_at?: InputMaybe<Scalars['timestamptz']>;
  user_id?: InputMaybe<Scalars['uuid']>;
};

/** update columns of table "workspace" */
export enum Workspace_Update_Column {
  /** column name */
  CreatedAt = 'created_at',
  /** column name */
  Id = 'id',
  /** column name */
  Name = 'name',
  /** column name */
  UpdatedAt = 'updated_at',
  /** column name */
  UserId = 'user_id'
}

export type InsertProviderRequestMutationVariables = Exact<{
  providerRequest: AuthProviderRequests_Insert_Input;
}>;


export type InsertProviderRequestMutation = { __typename?: 'mutation_root', insertAuthProviderRequest?: { __typename?: 'authProviderRequests', id: any, options?: any | null } | null };

export type DeleteProviderRequestMutationVariables = Exact<{
  id: Scalars['uuid'];
}>;


export type DeleteProviderRequestMutation = { __typename?: 'mutation_root', deleteAuthProviderRequest?: { __typename?: 'authProviderRequests', id: any, options?: any | null } | null };

export type ProviderRequestQueryVariables = Exact<{
  id: Scalars['uuid'];
}>;


export type ProviderRequestQuery = { __typename?: 'query_root', authProviderRequest?: { __typename?: 'authProviderRequests', id: any, options?: any | null } | null };

export type InsertRefreshTokenMutationVariables = Exact<{
  refreshToken: AuthRefreshTokens_Insert_Input;
}>;


export type InsertRefreshTokenMutation = { __typename?: 'mutation_root', insertAuthRefreshToken?: { __typename?: 'authRefreshTokens', refreshToken: any } | null };

export type DeleteRefreshTokenMutationVariables = Exact<{
  refreshToken: Scalars['uuid'];
}>;


export type DeleteRefreshTokenMutation = { __typename?: 'mutation_root', deleteAuthRefreshToken?: { __typename?: 'authRefreshTokens', refreshToken: any, expiresAt: any } | null };

export type DeleteUserRefreshTokensMutationVariables = Exact<{
  userId: Scalars['uuid'];
}>;


export type DeleteUserRefreshTokensMutation = { __typename?: 'mutation_root', deleteAuthRefreshTokens?: { __typename?: 'authRefreshTokens_mutation_response', affected_rows: number } | null };

export type DeleteExpiredRefreshTokensMutationVariables = Exact<{ [key: string]: never; }>;


export type DeleteExpiredRefreshTokensMutation = { __typename?: 'mutation_root', deleteAuthRefreshTokens?: { __typename?: 'authRefreshTokens_mutation_response', affected_rows: number } | null };

export type AuthUserProvidersQueryVariables = Exact<{
  provider: Scalars['String'];
  providerUserId: Scalars['String'];
}>;


export type AuthUserProvidersQuery = { __typename?: 'query_root', authUserProviders: Array<{ __typename?: 'authUserProviders', id: any, user: { __typename?: 'users', id: any } }> };

export type UserProviderQueryVariables = Exact<{
  userId: Scalars['uuid'];
  providerId: Scalars['String'];
}>;


export type UserProviderQuery = { __typename?: 'query_root', authUserProviders: Array<{ __typename?: 'authUserProviders', id: any, refreshToken?: string | null }> };

export type UpdateAuthUserproviderMutationVariables = Exact<{
  id: Scalars['uuid'];
  authUserProvider: AuthUserProviders_Set_Input;
}>;


export type UpdateAuthUserproviderMutation = { __typename?: 'mutation_root', updateAuthUserProvider?: { __typename?: 'authUserProviders', id: any } | null };

export type InsertUserRolesMutationVariables = Exact<{
  userRoles: Array<AuthUserRoles_Insert_Input> | AuthUserRoles_Insert_Input;
}>;


export type InsertUserRolesMutation = { __typename?: 'mutation_root', insertAuthUserRoles?: { __typename?: 'authUserRoles_mutation_response', affected_rows: number } | null };

export type DeleteUserRolesByUserIdMutationVariables = Exact<{
  userId: Scalars['uuid'];
}>;


export type DeleteUserRolesByUserIdMutation = { __typename?: 'mutation_root', deleteAuthUserRoles?: { __typename?: 'authUserRoles_mutation_response', affected_rows: number } | null };

export type UserFieldsFragment = { __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> };

export type UserQueryVariables = Exact<{
  id: Scalars['uuid'];
}>;


export type UserQuery = { __typename?: 'query_root', users_by_pk?: { __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> } | null };

export type UsersQueryVariables = Exact<{
  where: Users_Bool_Exp;
}>;


export type UsersQuery = { __typename?: 'query_root', users: Array<{ __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> }> };

export type GetUsersByRefreshTokenAndUpdateRefreshTokenExpiresAtMutationVariables = Exact<{
  refreshToken: Scalars['uuid'];
  expiresAt: Scalars['timestamptz'];
}>;


export type GetUsersByRefreshTokenAndUpdateRefreshTokenExpiresAtMutation = { __typename?: 'mutation_root', updateAuthRefreshTokens?: { __typename?: 'authRefreshTokens_mutation_response', returning: Array<{ __typename?: 'authRefreshTokens', refreshToken: any, user: { __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> } }> } | null };

export type GetUsersByRefreshTokenOldQueryVariables = Exact<{
  refreshToken: Scalars['uuid'];
}>;


export type GetUsersByRefreshTokenOldQuery = { __typename?: 'query_root', authRefreshTokens: Array<{ __typename?: 'authRefreshTokens', refreshToken: any, user: { __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> } }> };

export type UpdateUserMutationVariables = Exact<{
  id: Scalars['uuid'];
  user: Users_Set_Input;
}>;


export type UpdateUserMutation = { __typename?: 'mutation_root', update_users_by_pk?: { __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> } | null };

export type UpdateUserWhereMutationVariables = Exact<{
  where: Users_Bool_Exp;
  user: Users_Set_Input;
}>;


export type UpdateUserWhereMutation = { __typename?: 'mutation_root', update_users?: { __typename?: 'users_mutation_response', affected_rows: number } | null };

export type RotateUsersTicketMutationVariables = Exact<{
  oldTicket: Scalars['String'];
  newTicket: Scalars['String'];
  newTicketExpiresAt: Scalars['timestamptz'];
}>;


export type RotateUsersTicketMutation = { __typename?: 'mutation_root', update_users?: { __typename?: 'users_mutation_response', affected_rows: number } | null };

export type ChangeEmailsByTicketMutationVariables = Exact<{
  ticket: Scalars['String'];
  email: Scalars['citext'];
  newTicket: Scalars['String'];
  now: Scalars['timestamptz'];
}>;


export type ChangeEmailsByTicketMutation = { __typename?: 'mutation_root', update_users?: { __typename?: 'users_mutation_response', returning: Array<{ __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> }> } | null };

export type InsertUserMutationVariables = Exact<{
  user: Users_Insert_Input;
}>;


export type InsertUserMutation = { __typename?: 'mutation_root', insert_users_one?: { __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> } | null };

export type DeleteUserMutationVariables = Exact<{
  userId: Scalars['uuid'];
}>;


export type DeleteUserMutation = { __typename?: 'mutation_root', deleteAuthUserRoles?: { __typename?: 'authUserRoles_mutation_response', affected_rows: number } | null, delete_users_by_pk?: { __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> } | null };

export type DeanonymizeUserMutationVariables = Exact<{
  userId: Scalars['uuid'];
  avatarUrl?: InputMaybe<Scalars['String']>;
  role: Scalars['String'];
}>;


export type DeanonymizeUserMutation = { __typename?: 'mutation_root', updateAuthUserRoles?: { __typename?: 'authUserRoles_mutation_response', affected_rows: number } | null, update_users_by_pk?: { __typename?: 'users', id: any } | null };

export type InsertUserProviderToUserMutationVariables = Exact<{
  userProvider: AuthUserProviders_Insert_Input;
}>;


export type InsertUserProviderToUserMutation = { __typename?: 'mutation_root', insertAuthUserProvider?: { __typename?: 'authUserProviders', id: any } | null };

export type GetUserByTicketQueryVariables = Exact<{
  ticket: Scalars['String'];
}>;


export type GetUserByTicketQuery = { __typename?: 'query_root', users: Array<{ __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> }> };

export type UpdateUsersByTicketMutationVariables = Exact<{
  ticket: Scalars['String'];
  user: Users_Set_Input;
}>;


export type UpdateUsersByTicketMutation = { __typename?: 'mutation_root', update_users?: { __typename?: 'users_mutation_response', affected_rows: number, returning: Array<{ __typename?: 'users', id: any, createdAt: any, disabled: boolean, displayName: string, avatarUrl: string, email?: any | null, passwordHash?: string | null, emailVerified: boolean, phoneNumber?: string | null, phoneNumberVerified: boolean, defaultRole: string, isAnonymous: boolean, ticket?: string | null, otpHash?: string | null, totpSecret?: string | null, activeMfaType?: string | null, newEmail?: any | null, locale: string, metadata?: any | null, roles: Array<{ __typename?: 'authUserRoles', role: string }> }> } | null };

export const UserFieldsFragmentDoc = gql`
    fragment userFields on users {
  id
  createdAt
  disabled
  displayName
  avatarUrl
  email
  passwordHash
  emailVerified
  phoneNumber
  phoneNumberVerified
  defaultRole
  isAnonymous
  ticket
  otpHash
  totpSecret
  activeMfaType
  newEmail
  locale
  metadata
  roles {
    role
  }
}
    `;
export const InsertProviderRequestDocument = gql`
    mutation insertProviderRequest($providerRequest: authProviderRequests_insert_input!) {
  insertAuthProviderRequest(object: $providerRequest) {
    id
    options
  }
}
    `;
export const DeleteProviderRequestDocument = gql`
    mutation deleteProviderRequest($id: uuid!) {
  deleteAuthProviderRequest(id: $id) {
    id
    options
  }
}
    `;
export const ProviderRequestDocument = gql`
    query providerRequest($id: uuid!) {
  authProviderRequest(id: $id) {
    id
    options
  }
}
    `;
export const InsertRefreshTokenDocument = gql`
    mutation insertRefreshToken($refreshToken: authRefreshTokens_insert_input!) {
  insertAuthRefreshToken(object: $refreshToken) {
    refreshToken
  }
}
    `;
export const DeleteRefreshTokenDocument = gql`
    mutation deleteRefreshToken($refreshToken: uuid!) {
  deleteAuthRefreshToken(refreshToken: $refreshToken) {
    refreshToken
    expiresAt
  }
}
    `;
export const DeleteUserRefreshTokensDocument = gql`
    mutation deleteUserRefreshTokens($userId: uuid!) {
  deleteAuthRefreshTokens(where: {user: {id: {_eq: $userId}}}) {
    affected_rows
  }
}
    `;
export const DeleteExpiredRefreshTokensDocument = gql`
    mutation deleteExpiredRefreshTokens {
  deleteAuthRefreshTokens(where: {expiresAt: {_lt: now}}) {
    affected_rows
  }
}
    `;
export const AuthUserProvidersDocument = gql`
    query authUserProviders($provider: String!, $providerUserId: String!) {
  authUserProviders(
    where: {_and: {provider: {id: {_eq: $provider}}, providerUserId: {_eq: $providerUserId}}}
    limit: 1
  ) {
    id
    user {
      id
    }
  }
}
    `;
export const UserProviderDocument = gql`
    query userProvider($userId: uuid!, $providerId: String!) {
  authUserProviders(
    where: {_and: [{userId: {_eq: $userId}}, {providerId: {_eq: $providerId}}]}
    limit: 1
  ) {
    id
    refreshToken
  }
}
    `;
export const UpdateAuthUserproviderDocument = gql`
    mutation updateAuthUserprovider($id: uuid!, $authUserProvider: authUserProviders_set_input!) {
  updateAuthUserProvider(pk_columns: {id: $id}, _set: $authUserProvider) {
    id
  }
}
    `;
export const InsertUserRolesDocument = gql`
    mutation insertUserRoles($userRoles: [authUserRoles_insert_input!]!) {
  insertAuthUserRoles(objects: $userRoles) {
    affected_rows
  }
}
    `;
export const DeleteUserRolesByUserIdDocument = gql`
    mutation deleteUserRolesByUserId($userId: uuid!) {
  deleteAuthUserRoles(where: {userId: {_eq: $userId}}) {
    affected_rows
  }
}
    `;
export const UserDocument = gql`
    query user($id: uuid!) {
  users_by_pk(id: $id) {
    ...userFields
  }
}
    ${UserFieldsFragmentDoc}`;
export const UsersDocument = gql`
    query users($where: users_bool_exp!) {
  users(where: $where) {
    ...userFields
  }
}
    ${UserFieldsFragmentDoc}`;
export const GetUsersByRefreshTokenAndUpdateRefreshTokenExpiresAtDocument = gql`
    mutation getUsersByRefreshTokenAndUpdateRefreshTokenExpiresAt($refreshToken: uuid!, $expiresAt: timestamptz!) {
  updateAuthRefreshTokens(
    _set: {expiresAt: $expiresAt}
    where: {_and: [{refreshToken: {_eq: $refreshToken}}, {user: {disabled: {_eq: false}}}, {expiresAt: {_gte: now}}]}
  ) {
    returning {
      refreshToken
      user {
        ...userFields
      }
    }
  }
}
    ${UserFieldsFragmentDoc}`;
export const GetUsersByRefreshTokenOldDocument = gql`
    query getUsersByRefreshTokenOld($refreshToken: uuid!) {
  authRefreshTokens(
    where: {_and: [{refreshToken: {_eq: $refreshToken}}, {user: {disabled: {_eq: false}}}, {expiresAt: {_gte: now}}]}
  ) {
    refreshToken
    user {
      ...userFields
    }
  }
}
    ${UserFieldsFragmentDoc}`;
export const UpdateUserDocument = gql`
    mutation updateUser($id: uuid!, $user: users_set_input!) {
  update_users_by_pk(pk_columns: {id: $id}, _set: $user) {
    ...userFields
  }
}
    ${UserFieldsFragmentDoc}`;
export const UpdateUserWhereDocument = gql`
    mutation updateUserWhere($where: users_bool_exp!, $user: users_set_input!) {
  update_users(where: $where, _set: $user) {
    affected_rows
  }
}
    `;
export const RotateUsersTicketDocument = gql`
    mutation rotateUsersTicket($oldTicket: String!, $newTicket: String!, $newTicketExpiresAt: timestamptz!) {
  update_users(
    _set: {ticket: $newTicket, ticketExpiresAt: $newTicketExpiresAt}
    where: {ticket: {_eq: $oldTicket}}
  ) {
    affected_rows
  }
}
    `;
export const ChangeEmailsByTicketDocument = gql`
    mutation changeEmailsByTicket($ticket: String!, $email: citext!, $newTicket: String!, $now: timestamptz!) {
  update_users(
    where: {_and: [{ticket: {_eq: $ticket}}, {ticketExpiresAt: {_gt: $now}}]}
    _set: {email: $email, newEmail: null, ticket: $newTicket, ticketExpiresAt: $now}
  ) {
    returning {
      ...userFields
    }
  }
}
    ${UserFieldsFragmentDoc}`;
export const InsertUserDocument = gql`
    mutation insertUser($user: users_insert_input!) {
  insert_users_one(object: $user) {
    ...userFields
  }
}
    ${UserFieldsFragmentDoc}`;
export const DeleteUserDocument = gql`
    mutation deleteUser($userId: uuid!) {
  deleteAuthUserRoles(where: {userId: {_eq: $userId}}) {
    affected_rows
  }
  delete_users_by_pk(id: $userId) {
    ...userFields
  }
}
    ${UserFieldsFragmentDoc}`;
export const DeanonymizeUserDocument = gql`
    mutation deanonymizeUser($userId: uuid!, $avatarUrl: String, $role: String!) {
  updateAuthUserRoles(where: {user: {id: {_eq: $userId}}}, _set: {role: $role}) {
    affected_rows
  }
  update_users_by_pk(
    pk_columns: {id: $userId}
    _set: {avatarUrl: $avatarUrl, defaultRole: $role}
  ) {
    id
  }
}
    `;
export const InsertUserProviderToUserDocument = gql`
    mutation insertUserProviderToUser($userProvider: authUserProviders_insert_input!) {
  insertAuthUserProvider(object: $userProvider) {
    id
  }
}
    `;
export const GetUserByTicketDocument = gql`
    query getUserByTicket($ticket: String!) {
  users(where: {ticket: {_eq: $ticket}}) {
    ...userFields
  }
}
    ${UserFieldsFragmentDoc}`;
export const UpdateUsersByTicketDocument = gql`
    mutation updateUsersByTicket($ticket: String!, $user: users_set_input!) {
  update_users(
    where: {_and: [{ticket: {_eq: $ticket}}, {ticketExpiresAt: {_gt: now}}]}
    _set: $user
  ) {
    affected_rows
    returning {
      ...userFields
    }
  }
}
    ${UserFieldsFragmentDoc}`;

export type SdkFunctionWrapper = <T>(action: (requestHeaders?:Record<string, string>) => Promise<T>, operationName: string, operationType?: string) => Promise<T>;


const defaultWrapper: SdkFunctionWrapper = (action, _operationName, _operationType) => action();

export function getSdk(client: GraphQLClient, withWrapper: SdkFunctionWrapper = defaultWrapper) {
  return {
    insertProviderRequest(variables: InsertProviderRequestMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<InsertProviderRequestMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<InsertProviderRequestMutation>(InsertProviderRequestDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'insertProviderRequest', 'mutation');
    },
    deleteProviderRequest(variables: DeleteProviderRequestMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<DeleteProviderRequestMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<DeleteProviderRequestMutation>(DeleteProviderRequestDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'deleteProviderRequest', 'mutation');
    },
    providerRequest(variables: ProviderRequestQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<ProviderRequestQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<ProviderRequestQuery>(ProviderRequestDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'providerRequest', 'query');
    },
    insertRefreshToken(variables: InsertRefreshTokenMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<InsertRefreshTokenMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<InsertRefreshTokenMutation>(InsertRefreshTokenDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'insertRefreshToken', 'mutation');
    },
    deleteRefreshToken(variables: DeleteRefreshTokenMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<DeleteRefreshTokenMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<DeleteRefreshTokenMutation>(DeleteRefreshTokenDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'deleteRefreshToken', 'mutation');
    },
    deleteUserRefreshTokens(variables: DeleteUserRefreshTokensMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<DeleteUserRefreshTokensMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<DeleteUserRefreshTokensMutation>(DeleteUserRefreshTokensDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'deleteUserRefreshTokens', 'mutation');
    },
    deleteExpiredRefreshTokens(variables?: DeleteExpiredRefreshTokensMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<DeleteExpiredRefreshTokensMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<DeleteExpiredRefreshTokensMutation>(DeleteExpiredRefreshTokensDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'deleteExpiredRefreshTokens', 'mutation');
    },
    authUserProviders(variables: AuthUserProvidersQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<AuthUserProvidersQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<AuthUserProvidersQuery>(AuthUserProvidersDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'authUserProviders', 'query');
    },
    userProvider(variables: UserProviderQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<UserProviderQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<UserProviderQuery>(UserProviderDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'userProvider', 'query');
    },
    updateAuthUserprovider(variables: UpdateAuthUserproviderMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<UpdateAuthUserproviderMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<UpdateAuthUserproviderMutation>(UpdateAuthUserproviderDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'updateAuthUserprovider', 'mutation');
    },
    insertUserRoles(variables: InsertUserRolesMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<InsertUserRolesMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<InsertUserRolesMutation>(InsertUserRolesDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'insertUserRoles', 'mutation');
    },
    deleteUserRolesByUserId(variables: DeleteUserRolesByUserIdMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<DeleteUserRolesByUserIdMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<DeleteUserRolesByUserIdMutation>(DeleteUserRolesByUserIdDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'deleteUserRolesByUserId', 'mutation');
    },
    user(variables: UserQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<UserQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<UserQuery>(UserDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'user', 'query');
    },
    users(variables: UsersQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<UsersQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<UsersQuery>(UsersDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'users', 'query');
    },
    getUsersByRefreshTokenAndUpdateRefreshTokenExpiresAt(variables: GetUsersByRefreshTokenAndUpdateRefreshTokenExpiresAtMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<GetUsersByRefreshTokenAndUpdateRefreshTokenExpiresAtMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetUsersByRefreshTokenAndUpdateRefreshTokenExpiresAtMutation>(GetUsersByRefreshTokenAndUpdateRefreshTokenExpiresAtDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'getUsersByRefreshTokenAndUpdateRefreshTokenExpiresAt', 'mutation');
    },
    getUsersByRefreshTokenOld(variables: GetUsersByRefreshTokenOldQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<GetUsersByRefreshTokenOldQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetUsersByRefreshTokenOldQuery>(GetUsersByRefreshTokenOldDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'getUsersByRefreshTokenOld', 'query');
    },
    updateUser(variables: UpdateUserMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<UpdateUserMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<UpdateUserMutation>(UpdateUserDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'updateUser', 'mutation');
    },
    updateUserWhere(variables: UpdateUserWhereMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<UpdateUserWhereMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<UpdateUserWhereMutation>(UpdateUserWhereDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'updateUserWhere', 'mutation');
    },
    rotateUsersTicket(variables: RotateUsersTicketMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<RotateUsersTicketMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<RotateUsersTicketMutation>(RotateUsersTicketDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'rotateUsersTicket', 'mutation');
    },
    changeEmailsByTicket(variables: ChangeEmailsByTicketMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<ChangeEmailsByTicketMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<ChangeEmailsByTicketMutation>(ChangeEmailsByTicketDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'changeEmailsByTicket', 'mutation');
    },
    insertUser(variables: InsertUserMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<InsertUserMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<InsertUserMutation>(InsertUserDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'insertUser', 'mutation');
    },
    deleteUser(variables: DeleteUserMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<DeleteUserMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<DeleteUserMutation>(DeleteUserDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'deleteUser', 'mutation');
    },
    deanonymizeUser(variables: DeanonymizeUserMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<DeanonymizeUserMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<DeanonymizeUserMutation>(DeanonymizeUserDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'deanonymizeUser', 'mutation');
    },
    insertUserProviderToUser(variables: InsertUserProviderToUserMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<InsertUserProviderToUserMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<InsertUserProviderToUserMutation>(InsertUserProviderToUserDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'insertUserProviderToUser', 'mutation');
    },
    getUserByTicket(variables: GetUserByTicketQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<GetUserByTicketQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetUserByTicketQuery>(GetUserByTicketDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'getUserByTicket', 'query');
    },
    updateUsersByTicket(variables: UpdateUsersByTicketMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<UpdateUsersByTicketMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<UpdateUsersByTicketMutation>(UpdateUsersByTicketDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'updateUsersByTicket', 'mutation');
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;