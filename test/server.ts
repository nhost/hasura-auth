import { createApp } from '@/app';
import supertest, { SuperTest, Test } from 'supertest';

let _request: SuperTest<Test> | null;

export const getRequestClient = async () => {
  if (!_request) {
    const app = await createApp();
    _request = supertest(app);
  }
  return _request;
};
