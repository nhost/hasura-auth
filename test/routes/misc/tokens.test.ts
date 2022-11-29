import { v4 as uuidv4 } from 'uuid';
import { StatusCodes } from 'http-status-codes';
import { Client } from 'pg';
import crypto from 'crypto';

import { ENV } from '../../../src/utils/env';
import { request } from '../../server';
import { decodeAccessToken, isValidAccessToken } from '../../utils';

describe('token', () => {
  let client: Client;

  beforeAll(async () => {
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });
    await client.connect();
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ANONYMOUS_USERS_ENABLED: true,
    });
  });

  afterAll(async () => {
    await client.end();
  });

  beforeEach(async () => {
    await client.query(`DELETE FROM auth.users;`);
  });

  it('should should sign in and get access token with standard claims', async () => {
    // set env vars

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

  it('should refresh with a token stored as plain text', async () => {
    const { body } = await request
      .post('/signin/anonymous')
      .send()
      .expect(StatusCodes.OK);

    const refreshToken = uuidv4();
    const id = (await decodeAccessToken(body.session.accessToken))?.[
      'https://hasura.io/jwt/claims'
    ]['x-hasura-user-id'];
    await client.query(
      `INSERT INTO "auth"."refresh_tokens" (refresh_token, user_id, expires_at) VALUES ($1, $2, $3);`,
      [
        refreshToken,
        id,
        new Date(Date.now() + 1000 * 60 * 60 * 24 * 30).toISOString(),
      ]
    );

    await request.post('/token').send({ refreshToken }).expect(StatusCodes.OK);
  });

  it('should refresh with a token stored as SHA256', async () => {
    const { body } = await request
      .post('/signin/anonymous')
      .send()
      .expect(StatusCodes.OK);

    const refreshToken = uuidv4();
    const hashedToken = crypto
      .createHash('sha256')
      .update(refreshToken)
      .digest('hex');
    const id = (await decodeAccessToken(body.session.accessToken))?.[
      'https://hasura.io/jwt/claims'
    ]['x-hasura-user-id'];
    await client.query(
      `INSERT INTO "auth"."refresh_tokens" (refresh_token, user_id, expires_at) VALUES ($1, $2, $3);`,
      [
        hashedToken,
        id,
        new Date(Date.now() + 1000 * 60 * 60 * 24 * 30).toISOString(),
      ]
    );

    await request.post('/token').send({ refreshToken }).expect(StatusCodes.OK);
  });
});
