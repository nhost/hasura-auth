import {
  gqlSdk,
  InsertUserMutation,
  InsertUserMutationVariables,
} from '../gql-sdk';

type UserInput = InsertUserMutationVariables['user'];
type UserOutput = NonNullable<InsertUserMutation['insertUser']>;

export const insertUser = async (user: UserInput): Promise<UserOutput> => {
  const { insertUser } = await gqlSdk.insertUser({
    user,
  });
  if (!insertUser) {
    throw new Error('Could not insert user');
  }
  return insertUser;
};
