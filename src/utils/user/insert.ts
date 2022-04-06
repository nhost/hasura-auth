import { gqlSdk } from "../gql-sdk";
import { InsertUserMutation, InsertUserMutationVariables } from "../__generated__/graphql-request";

type UserInput = InsertUserMutationVariables['user'];
type UserOutput = NonNullable<InsertUserMutation['insert_users_one']>;

export const insertUser = async (user: UserInput): Promise<UserOutput> => {
  const { insert_users_one: insertUser } = await gqlSdk.insertUser({
    user,
  });
  if (!insertUser) {
    throw new Error('Could not insert user');
  }
  return insertUser;
};
