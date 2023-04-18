import { ENV } from '@/utils';
import * as faker from 'faker';
import { StatusCodes } from 'http-status-codes';
import { Client } from 'pg';
import { request } from '../../server';
import { deleteAllMailHogEmails } from '../../utils';

describe('personal access token', () => {
  let client: Client;
  const email = faker.internet.email();
  const password = faker.internet.password();

  beforeAll(async () => {
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });

    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED: false,
    });

    await client.connect();
    await deleteAllMailHogEmails();
  });

  afterAll(async () => {
    await client.end();
  });

  beforeEach(async () => {
    await client.query(`DELETE FROM auth.users;`);
    await client.query(`DELETE FROM auth.refresh_tokens;`);

    await request.post('/signup/email-password').send({
      email,
      password,
    });
  });

  test('should not allow unauthenticated users to create a personal access token', async () => {
    await request
      .post('/pat')
      .send({ expiresAt: new Date() })
      .expect(StatusCodes.UNAUTHORIZED);
  });

  test('should not allow an expiry date that is less than 7 days from now', async () => {
    const response = await request.post('/signin/email-password').send({
      email,
      password,
    });

    const { accessToken } = response.body?.session;

    await request
      .post('/pat')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({ expiresAt: new Date() })
      .expect(StatusCodes.BAD_REQUEST);
  });

  test('should be able to add metadata to a personal access token', async () => {
    const response = await request.post('/signin/email-password').send({
      email,
      password,
    });

    const { accessToken } = response.body?.session;

    const patResponse = await request
      .post('/pat')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({
        expiresAt: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000),
        metadata: { name: 'Test PAT' },
      })
      .expect(StatusCodes.OK);

    const { rows } = await client.query(
      'SELECT * FROM auth.refresh_tokens WHERE refresh_token=$1;',
      [patResponse.body?.personalAccessToken]
    );

    expect(rows[0].metadata).toMatchObject({ name: 'Test PAT' });
  });

  test('should authenticate using the PAT workflow', async () => {
    const response = await request.post('/signin/email-password').send({
      email,
      password,
    });

    const { accessToken } = response.body?.session;

    const patResponse = await request
      .post('/pat')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({
        expiresAt: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000),
      })
      .expect(StatusCodes.OK);

    const { personalAccessToken } = patResponse.body;

    expect(personalAccessToken).toBeDefined();

    const patSignInResponse = await request
      .post('/signin/pat')
      .send({ personalAccessToken })
      .expect(StatusCodes.OK);

    expect(patSignInResponse.body?.session?.accessToken).toBeDefined();
  });
});
