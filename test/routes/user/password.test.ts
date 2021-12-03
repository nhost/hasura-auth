import { Client } from 'pg';
// import { v4 as uuidv4 } from 'uuid';

import { request } from '../../server';
import { ENV } from '../../../src/utils/env';
import { SignInResponse } from '../../../src/types';
import { mailHogSearch } from '../../utils';

describe('user password', () => {
  let client: any;
  let accessToken: string | undefined;
  let body: SignInResponse | undefined;
  const email = 'asdasd@asdasd.com';
  const password = '123123123';

  beforeAll(async () => {
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });
    await client.connect();
  });

  afterAll(() => {
    client.end();
  });

  beforeEach(async () => {
    await client.query(`DELETE FROM auth.users;`);
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED: false,
    });

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect(200);

    const response = await request
      .post('/signin/email-password')
      .send({ email, password })
      .expect(200);

    body = response.body;

    if (!body?.session) {
      throw new Error('session is not set');
    }
    accessToken = body.session.accessToken;
  });

  it('should change password with old password', async () => {
    const oldPassword = password;
    const newPassword = '543543543';

    await request
      .post('/user/password')
      .set('Authorization', `Bearer ${accessToken}`)
      .send({ oldPassword, newPassword })
      .expect(200);

    await request
      .post('/signin/email-password')
      .send({ email, password: oldPassword })
      .expect(401);

    await request
      .post('/signin/email-password')
      .send({ email, password: newPassword })
      .expect(200);
  });

  it('should change password with ticket', async () => {
    await request.post('/user/password/reset').send({ email }).expect(200);

    // get ticket from email
    const [message] = await mailHogSearch(email);
    expect(message).toBeTruthy();

    const ticket = message.Content.Headers['X-Ticket'][0];
    const redirectTo = message.Content.Headers['X-Redirect-To'][0];

    // use password reset link
    await request
      .get(
        `/verify?ticket=${ticket}&type=signinPasswordless&redirectTo=${redirectTo}`
      )
      .expect(302);

    // TODO
    // get refershToken from previous request

    // request new access token

    // use access token to update password

    // const oldPassword = password;
    // const newPassword = '543543543';

    // await request
    //   .post('/user/password')
    //   .set('Authorization', `Bearer ${accessToken}`)
    //   .send({ ticket: 'incorrect', newPassword })
    //   .expect(400);

    // await request
    //   .post('/user/password')
    //   .set('Authorization', `Bearer ${accessToken}`)
    //   .send({ ticket: `passwordReset:${uuidv4()}`, newPassword })
    //   .expect(401);

    // await request
    //   .post('/user/password')
    //   .set('Authorization', `Bearer ${accessToken}`)
    //   .send({ ticket, newPassword })
    //   .expect(200);

    // await request
    //   .post('/signin/email-password')
    //   .send({ email, password: oldPassword })
    //   .expect(401);

    // await request
    //   .post('/signin/email-password')
    //   .send({ email, password: newPassword })
    //   .expect(200);
  });

  it('should be able to pass "redirectTo" when changing password with ticket when ', async () => {
    const options = {
      redirectTo: 'http://localhost:3000/change-password-redirect',
    };

    await request
      .post('/user/password/reset')
      .send({ email, options })
      .expect(200);

    // get ticket from email
    const [message] = await mailHogSearch(email);
    expect(message).toBeTruthy();

    const ticket = message.Content.Headers['X-Ticket'][0];
    const redirectTo = message.Content.Headers['X-Redirect-To'][0];

    // use password reset link
    await request
      .get(
        `/verify?ticket=${ticket}&type=signinPasswordless&redirectTo=${redirectTo}`
      )
      .expect(302);

    expect(redirectTo).toStrictEqual(options.redirectTo);
  });
});
