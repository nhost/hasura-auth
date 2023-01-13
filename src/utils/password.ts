/**
 * Password hashing function.
 * @param password Password to hash.
 */
export const hashPassword = async (password: string): Promise<string> => {
  const { hash } = await import('bcrypt');
  return await hash(password, 10);
};
