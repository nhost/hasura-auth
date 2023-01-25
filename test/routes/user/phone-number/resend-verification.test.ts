import { Client } from 'pg';
import * as faker from 'faker';
import { StatusCodes } from 'http-status-codes';

import { ENV } from '../../../../src/utils/env';
import { request } from '../../../server';
import { SignInResponse } from '@/types';

describe('resend phone verification', () => {
  let client: Client;
  let accessToken: string | undefined;
  let body: SignInResponse | undefined;
  const email = faker.internet.email();
  const password = faker.internet.password(8);

  beforeAll(async () => {
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
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_SMS_PASSWORDLESS_ENABLED: true,
      AUTH_SMS_TEST_PHONE_NUMBERS: '+359837025545,+359837025555',
      AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED: false,
    });

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect(StatusCodes.OK);

    const response = await request
      .post('/signin/email-password')
      .send({ email, password })
      .expect(StatusCodes.OK);
    body = response.body;

    if (!body?.session) {
      throw new Error('session is not set');
    }

    accessToken = body.session.accessToken;

    await client.query(
      `UPDATE auth.users SET phone_number = '+359837025545', phone_number_verified = false WHERE id = '${body.session.user?.id}';`
    );
  });

  it('should be able to resend verification code', async () => {
    expect(body?.session).toBeTruthy();

    await request
      .post('/user/phone-number/resend-verification')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({
        phoneNumber: '+359837025545',
      })
      .expect(StatusCodes.OK);
  });

  it('should not be able to resend verification code if phone is verified', async () => {
    expect(body?.session).toBeTruthy();

    await client.query(
      `UPDATE auth.users SET phone_number = '+359837025545', phone_number_verified = true WHERE id = '${body?.session?.user?.id}';`
    );

    await request
      .post('/user/phone-number/resend-verification')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({
        phoneNumber: '+359837025545',
      })
      .expect(StatusCodes.CONFLICT);
  });

  it('should not be able to resend verification code if passing different phone number', async () => {
    expect(body?.session).toBeTruthy();

    await request
      .post('/user/phone-number/resend-verification')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({
        phoneNumber: '+359837025555',
      })
      .expect(StatusCodes.CONFLICT);
  });

  it('should be able to send verification code on new phone number', async () => {
    expect(body?.session).toBeTruthy();

    await request
      .post('/user/phone-number/change')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({
        newPhoneNumber: '+359837025555',
      })
      .expect(StatusCodes.OK);

    await request
      .post('/user/phone-number/resend-verification')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({
        phoneNumber: '+359837025555',
      })
      .expect(StatusCodes.OK);
  });
});
