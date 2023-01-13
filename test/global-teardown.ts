require('tsconfig-paths/register');

import { pgPool } from '@/utils';

export default async (): Promise<void> => {
  await pgPool.end();
};
