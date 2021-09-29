import { Client } from 'pg';

import { ENV } from '../../../src/utils/env';
import { request } from '../../server';
import { mailHogSearch, deleteAllMailHogEmails } from '../../utils';
import { trackTable, setTableCustomization } from '../../../src/metadata';

describe('email-password', () => {
  let client: any;

  beforeAll(async () => {
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });
    await client.connect();
    await deleteAllMailHogEmails();
  });

  afterAll(() => {
    client.end();
  });

  beforeEach(async () => {
    await client.query(`DELETE FROM auth.users;`);
  });

  it('should sign up user', async () => {
    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_SIGNUP_PROFILE_FIELDS: '',
      AUTH_EMAILS_ENABLED: true,
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(200);
  });

  it('should fail to sign up with same email', async () => {
    // set env vars
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(409);
  });

  it('should fail with weak password', async () => {
    // set env vars
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_HIBP_ENABLED: true,
      AUTH_ACCESS_CONTROL_ALLOW_LIST: '',
      AUTH_ACCESS_CONTROL_BLOCK_LIST: '',
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(400);
  });

  it('should succeed to sign up with different emails', async () => {
    // set env vars
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_HIBP_ENABLED: false,
      AUTH_ACCESS_CONTROL_ALLOW_LIST: '',
      AUTH_ACCESS_CONTROL_BLOCK_LIST: '',
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoes@example.com', password: '123456' })
      .expect(200);
  });

  it('should success with SMTP settings', async () => {
    // set env vars
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_HIBP_ENABLED: false,
      AUTH_SIGNIN_EMAIL_VERIFIED_REQUIRED: true,
    });

    const email = 'joedoe@example.com';

    await request
      .post('/signup/email-password')
      .send({ email, password: '123456' })
      .expect(200);

    // fetch email from mailhog and check ticket
    const [message] = await mailHogSearch(email);
    expect(message).toBeTruthy();
    const ticket = message.Content.Headers['X-Ticket'][0];
    expect(ticket.startsWith('verifyEmail:')).toBeTruthy();
  });

  it('default role must be part of allowed roles', async () => {
    const email = 'joedoe@example.com';

    await request
      .post('/signup/email-password')
      .send({
        email,
        password: '123456',
        defaultRole: 'user',
        allowedRoles: ['editor'],
      })
      .expect(400);
  });

  it('allowed roles must be subset of env var ALLOWED_USER_ROLES', async () => {
    // set env vars
    await await request.post('/change-env').send({
      ALLOWED_USER_ROLES: 'user,editor',
    });

    const email = 'joedoe@example.com';

    await request
      .post('/signup/email-password')
      .send({
        email,
        password: '123456',
        defaultRole: 'user',
        allowedRoles: ['user', 'some-other-role'],
      })
      .expect(400);
  });

  it('user must verify email before being able to sign in', async () => {
    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_SIGNIN_EMAIL_VERIFIED_REQUIRED: true,
    });

    const email = 'joedoe@example.com';
    const password = '123123';

    await request
      .post('/signup/email-password')
      .send({ email, password })
      .expect(200);

    await request
      .post('/signin/email-password')
      .send({ email, password })
      .expect(401);

    // get ticket from email
    const [message] = await mailHogSearch(email);
    expect(message).toBeTruthy();
    const ticket = message.Content.Headers['X-Ticket'][0];
    expect(ticket.startsWith('verifyEmail:')).toBeTruthy();

    // use ticket to verify email
    await request
      .post('/user/email/verify')
      .send({ email, ticket })
      .expect(200);

    // sign in should now work

    await request
      .post('/signin/email-password')
      .send({ email, password })
      .expect(200);
  });
});

describe('email-password access control lists', () => {
  let client: any;
  beforeAll(async () => {
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });
    await client.connect();
    await deleteAllMailHogEmails();
  });

  afterAll(() => {
    client.end();
  });

  beforeEach(async () => {
    await client.query(`DELETE FROM auth.users;`);
  });

  it('should only allow emails that are allowed', async () => {
    // set env vars
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ACCESS_CONTROL_ALLOWED_EMAILS: 'aaa@nhost.io, bbb@nhost.io',
      AUTH_ACCESS_CONTROL_ALLOWED_EMAIL_DOMAINS: '',
      AUTH_ACCESS_CONTROL_BLOCKED_EMAILS: '',
      AUTH_ACCESS_CONTROL_BLOCKED_EMAIL_DOMAINS: '',
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'aaa@nhost.io', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'bbb@nhost.io', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'ccc@nhost.io', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'aaa@other.io', password: '123456' })
      .expect(403);
  });

  it('should only allow email domains that are allowed', async () => {
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ACCESS_CONTROL_ALLOWED_EMAILS: '',
      AUTH_ACCESS_CONTROL_ALLOWED_EMAIL_DOMAINS: 'nhost.io',
      AUTH_ACCESS_CONTROL_BLOCKED_EMAILS: '',
      AUTH_ACCESS_CONTROL_BLOCKED_EMAIL_DOMAINS: '',
    });

    console.log('1');

    let res;
    res = await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(403);

    console.log(res);

    console.log('2');
    res = await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@nhost.io', password: '123456' })
      .expect(200);

    console.log('3');
    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@other.io', password: '123456' })
      .expect(403);
  });

  it('should block blocked emails', async () => {
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ACCESS_CONTROL_ALLOWED_EMAILS: '',
      AUTH_ACCESS_CONTROL_ALLOWED_EMAIL_DOMAINS: '',
      AUTH_ACCESS_CONTROL_BLOCKED_EMAILS: 'aaa@nhost.io, bbb@nhost.io',
      AUTH_ACCESS_CONTROL_BLOCKED_EMAIL_DOMAINS: '',
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'aaa@nhost.io', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'bbb@nhost.io', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'ccc@nhost.io', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@other.io', password: '123456' })
      .expect(200);
  });

  it('should block blocked email domains', async () => {
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ACCESS_CONTROL_ALLOWED_EMAILS: '',
      AUTH_ACCESS_CONTROL_ALLOWED_EMAIL_DOMAINS: '',
      AUTH_ACCESS_CONTROL_BLOCKED_EMAILS: '',
      AUTH_ACCESS_CONTROL_BLOCKED_EMAIL_DOMAINS: 'tesla.com, ikea.se',
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'aaa@nhost.io', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'bbb@nhost.io', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'aaa@tesla.com', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'bbb@tesla.com', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'aaa@ikea.se', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'bbb@ikea.se', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'aaa@ikeaa.se', password: '123456' })
      .expect(200);
  });

  it('should fail sign up if email is not allowed', async () => {
    // set env vars
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ACCESS_CONTROL_ALLOW_LIST: 'vip@example.com',
      AUTH_ACCESS_CONTROL_BLOCK_LIST: '',
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'vip@example.com', password: '123456' })
      .expect(200);
  });

  it('should fail sign up if email domain is blocked', async () => {
    // set env vars
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ACCESS_CONTROL_ALLOW_LIST: '',
      AUTH_ACCESS_CONTROL_BLOCK_LIST: '*@example.com',
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@nhost.io', password: '123456' })
      .expect(200);
  });

  it('should fail sign up if email is blocked', async () => {
    // set env vars
    await await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_ACCESS_CONTROL_ALLOW_LIST: '',
      AUTH_ACCESS_CONTROL_BLOCK_LIST: 'joedoe@example.com',
    });

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@example.com', password: '123456' })
      .expect(403);

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe2@example.com', password: '123456' })
      .expect(200);

    await request
      .post('/signup/email-password')
      .send({ email: 'joedoe@nhost.io', password: '123456' })
      .expect(200);
  });
});

describe('email-password with profile table', () => {
  let client: any;

  beforeAll(async () => {
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });
    await client.connect();
    await deleteAllMailHogEmails();

    // create profile table
    console.log('drop profiles table:');
    await client.query(`DROP TABLE IF EXISTS public.profiles;`);

    console.log('create profiles table');
    await client.query(`
    CREATE TABLE public.profiles (
      user_id uuid PRIMARY KEY,
      company_id int NOT NULL,
      foreign key(user_id) references auth.users(id) on delete cascade
    );
    `);

    // track table
    await trackTable({
      source: 'default',
      table: { schema: 'public', name: 'profiles' },
    });

    // set profile customization
    await setTableCustomization({
      source: 'default',
      table: {
        schema: 'public',
        name: 'profiles',
      },
      configuration: {
        identifier: 'profile',
        custom_root_fields: {
          select: 'profiles',
          select_by_pk: 'profile',
          select_aggregate: 'profilesAggregat',
          insert: 'insertProfiles',
          insert_one: 'insertProfile',
          update: 'updateProfiles',
          update_by_pk: 'updateProfile',
          delete: 'deleteProfiles',
          delete_by_pk: 'deleteProfile',
        },
        custom_column_names: {
          user_id: 'userId',
          company_id: 'companyId',
        },
      },
    });
  });

  afterAll(() => {
    client.end();
  });

  beforeEach(async () => {
    // clear database
    await client.query(`DELETE FROM auth.users;`);
    await client.query(`DELETE FROM public.profiles;`);
  });

  it('should sign up user with profile data', async () => {
    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_SIGNUP_PROFILE_FIELDS: 'companyId',
    });

    await request
      .post('/signup/email-password')
      .send({
        email: 'joedoe@example.com',
        password: '123456',
        profile: { companyId: 1337 },
      })
      .expect(200);
  });

  it('should fail to sign up user with extra profile data', async () => {
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_SIGNUP_PROFILE_FIELDS: 'companyId',
    });

    await request
      .post('/signup/email-password')
      .send({
        email: 'joedoe@example.com',
        password: '123456',
        profile: { extra: true, companyId: 1337 },
      })
      .expect(400);
  });

  it('should fail to sign up because registration custom fields does not exist in database', async () => {
    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_SIGNUP_PROFILE_FIELDS: 'incorrect',
    });

    await request
      .post('/signup/email-password')
      .send({
        email: 'joedoe@example.com',
        password: '123456',
        profile: { incorrect: 1337 },
      })
      .expect(500);
  });

  it('should fail to sign up user with no profile data', async () => {
    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_SIGNUP_PROFILE_FIELDS: 'companyId',
    });

    await request
      .post('/signup/email-password')
      .send({
        email: 'joedoe@example.com',
        password: '123456',
      })
      .expect(400, {
        statusCode: 400,
        error: 'Bad Request',
        message: 'Profile required',
      });
  });

  it('should fail to insert profile', async () => {
    await client.query(`ALTER TABLE public.profiles DROP COLUMN company_id`);

    // set env vars
    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_SIGNUP_PROFILE_FIELDS: 'companyId',
    });

    await request
      .post('/signup/email-password')
      .send({
        email: 'joedoe@example.com',
        password: '123456',
        profile: {
          companyId: 1336,
        },
      })
      .expect(500);
  });
});
