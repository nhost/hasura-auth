import { StatusCodes } from 'http-status-codes';
import { Client } from 'pg';
import { SuperTest, Test } from 'supertest';

import { ENV } from '../../../src/utils/env';
import { getRequestClient } from '../../server';
import { isValidAccessToken } from '../../utils';

describe('token', () => {
  let client: Client;
  let request: SuperTest<Test>;

  beforeAll(async () => {
    request = await getRequestClient();
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });
    await client.connect();
  });

  afterAll(async () => {
    await client.end();
  });

  beforeEach(async () => {
    await client.query(`DELETE FROM auth.users;`);
  });

  it('should should sign in and get access token with standard claims', async () => {
    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ANONYMOUS_USERS_ENABLED: true,
    });

    const { body } = await request
      .post('/signin/anonymous')
      .send()
      .expect(StatusCodes.OK);

    const { accessToken, accessTokenExpiresIn, refreshToken } = body.session;
    const { mfa } = body;

    expect(await isValidAccessToken(accessToken)).toBe(true);
    expect(typeof accessTokenExpiresIn).toBe('number');
    expect(typeof refreshToken).toBe('string');
    expect(mfa).toBe(null);
  });
});
