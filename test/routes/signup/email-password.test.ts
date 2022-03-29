import { Client } from 'pg';
import * as faker from 'faker';

import { ENV } from '../../../src/utils/env';
import { request } from '../../server';
import { mailHogSearch, deleteAllMailHogEmails } from '../../utils';

describe('email-password', () => {
  let client: Client;

  beforeAll(async () => {
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });
    await client.connect();
    await deleteAllMailHogEmails();
  });

  afterAll(async () => {
    await client.end();
  });

  beforeEach(async () => {
    await client.query(`DELETE FROM auth.users;`);
  });

  it('should sign up user', async () => {
    const email = faker.internet.email();
    const password = faker.internet.password();

    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
    });

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect(200);
  });

  it('should fail to sign up with same email', async () => {
    const email = faker.internet.email();
    const password = faker.internet.password();

    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
    });

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect(409);
  });

  it('should fail with weak password', async () => {
    const email = faker.internet.email();
    const password = '123456';
    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_PASSWORD_HIBP_ENABLED: true,
    });

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect(400);
  });

  it('should succeed to sign up with different emails', async () => {
    const emailA = faker.internet.email();
    const passwordA = faker.internet.password();
    const emailB = faker.internet.email();
    const passwordB = faker.internet.password();
    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_PASSWORD_HIBP_ENABLED: false,
    });

    await request
      .post('/signup/email-password')
      .send({ email: emailA, password: passwordA })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: emailB, password: passwordB })
      .expect(200);
  });

  it('should success with SMTP settings', async () => {
    const email = faker.internet.email();
    const password = faker.internet.password();

    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_HIBP_ENABLED: false,
      AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED: true,
    });

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect(200);

    // fetch email from mailhog and check ticket
    const [message] = await mailHogSearch(email);
    expect(message).toBeTruthy();
    const ticket = message.Content.Headers['X-Ticket'][0];
    expect(ticket.startsWith('verifyEmail:')).toBeTruthy();
  });

  it('default role must be part of allowed roles', async () => {
    const email = faker.internet.email();
    const password = faker.internet.password();

    await request
      .post('/signup/email-password')
      .send({
        email,
        password,
        defaultRole: 'user',
        allowedRoles: ['editor'],
      })
      .expect(400);
  });

  it('allowed roles must be subset of env var AUTH_USER_DEFAULT_ALLOWED_ROLES', async () => {
    const email = faker.internet.email();
    const password = faker.internet.password();

    // set env vars
    await request.post('/change-env').send({
      ALLOWED_USER_ROLES: 'user,editor',
    });

    await request
      .post('/signup/email-password')
      .send({
        email,
        password,
        defaultRole: 'user',
        allowedRoles: ['user', 'some-other-role'],
      })
      .expect(400);
  });

  it('user must verify email before being able to sign in', async () => {
    const email = faker.internet.email();
    const password = faker.internet.password();

    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED: true,
      AUTH_USER_DEFAULT_ALLOWED_ROLES: '',
    });

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect((res) => (res.status != 200 ? console.log(res.body) : 0));

    await request
      .post('/signin/email-password')
      .send({ email, password })
      .expect(401);

    // get ticket from email
    const [message] = await mailHogSearch(email);
    expect(message).toBeTruthy();
    const ticket = message.Content.Headers['X-Ticket'][0];
    const redirectTo = message.Content.Headers['X-Redirect-To'][0];

    // use ticket to verify email
    await request
      .get(
        `/verify?ticket=${ticket}&type=signinPasswordless&redirectTo=${redirectTo}`
      )
      .expect(302);

    // sign in should now work
    await request
      .post('/signin/email-password')
      .send({ email, password })
      .expect(200);
  });
});
