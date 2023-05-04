import { authMetadataSource } from '@/metadata';
import { ENV, escapeValueToPg, overrideMetadata } from '@/utils';
import * as faker from 'faker';
import { Client } from 'pg';
import { request } from '../../server';
import { decodeAccessToken } from '../../utils';

describe('custom JWT claims', () => {
  let client: Client;
  const organisationId = faker.datatype.uuid();
  const projects = [...Array(3).keys()].map(faker.datatype.uuid);

  beforeAll(async () => {
    client = new Client({
      connectionString: ENV.HASURA_GRAPHQL_DATABASE_URL,
    });
    await client.connect();
    try {
      await client.query(`
      CREATE TABLE IF NOT EXISTS public.profiles (
        id uuid PRIMARY KEY
            CONSTRAINT fk_user REFERENCES auth.users(id) 
            ON UPDATE CASCADE ON DELETE CASCADE,
        organisation_id uuid);
      CREATE TABLE public.organisations (id uuid primary key default gen_random_uuid());
      ALTER TABLE public.profiles ADD FOREIGN KEY (organisation_id) REFERENCES public.organisations;
      INSERT INTO public.organisations(id) VALUES ('${organisationId}');
      CREATE TABLE public.projects (id uuid primary key);
      CREATE TABLE public.project_members (id uuid primary key default gen_random_uuid(), user_id uuid, project_id uuid);
      ALTER TABLE public.project_members ADD FOREIGN KEY (project_id) REFERENCES public.projects;
      ALTER TABLE public.project_members ADD FOREIGN KEY (user_id) REFERENCES public.profiles;
      INSERT INTO public.projects VALUES ${projects
        .map((id) => "('" + id + "')")
        .join(',')};
      `);
    } catch (e) {
      console.log('cannot create custom jwt test tables');
    }

    await overrideMetadata({
      ...authMetadataSource,
      tables: [
        ...authMetadataSource.tables.filter(
          ({ table }) => table.name !== 'users'
        ),
        { table: { schema: 'public', name: 'projects' } },
        {
          table: { schema: 'public', name: 'project_members' },
          object_relationships: [
            {
              name: 'project',
              using: {
                foreign_key_constraint_on: 'project_id',
              },
            },
          ],
        },
        { table: { schema: 'public', name: 'organisations' } },
        {
          table: { schema: 'public', name: 'profiles' },
          object_relationships: [
            {
              name: 'organisation',
              using: {
                foreign_key_constraint_on: 'organisation_id',
              },
            },
          ],
          array_relationships: [
            {
              name: 'contributesTo',
              using: {
                foreign_key_constraint_on: {
                  table: {
                    schema: 'public',
                    name: 'project_members',
                  },
                  column: 'user_id',
                },
              },
            },
          ],
        },
        {
          table: {
            name: 'users',
            schema: 'auth',
          },
          configuration: {
            column_config: {
              active_mfa_type: {
                custom_name: 'activeMfaType',
              },
              avatar_url: {
                custom_name: 'avatarUrl',
              },
              created_at: {
                custom_name: 'createdAt',
              },
              default_role: {
                custom_name: 'defaultRole',
              },
              disabled: {
                custom_name: 'disabled',
              },
              display_name: {
                custom_name: 'displayName',
              },
              email: {
                custom_name: 'email',
              },
              email_verified: {
                custom_name: 'emailVerified',
              },
              id: {
                custom_name: 'id',
              },
              is_anonymous: {
                custom_name: 'isAnonymous',
              },
              last_seen: {
                custom_name: 'lastSeen',
              },
              locale: {
                custom_name: 'locale',
              },
              new_email: {
                custom_name: 'newEmail',
              },
              otp_hash: {
                custom_name: 'otpHash',
              },
              otp_hash_expires_at: {
                custom_name: 'otpHashExpiresAt',
              },
              otp_method_last_used: {
                custom_name: 'otpMethodLastUsed',
              },
              password_hash: {
                custom_name: 'passwordHash',
              },
              phone_number: {
                custom_name: 'phoneNumber',
              },
              phone_number_verified: {
                custom_name: 'phoneNumberVerified',
              },
              ticket: {
                custom_name: 'ticket',
              },
              ticket_expires_at: {
                custom_name: 'ticketExpiresAt',
              },
              totp_secret: {
                custom_name: 'totpSecret',
              },
              updated_at: {
                custom_name: 'updatedAt',
              },
              webauthn_current_challenge: {
                custom_name: 'currentChallenge',
              },
            },
            custom_column_names: {
              active_mfa_type: 'activeMfaType',
              avatar_url: 'avatarUrl',
              created_at: 'createdAt',
              default_role: 'defaultRole',
              disabled: 'disabled',
              display_name: 'displayName',
              email: 'email',
              email_verified: 'emailVerified',
              id: 'id',
              is_anonymous: 'isAnonymous',
              last_seen: 'lastSeen',
              locale: 'locale',
              new_email: 'newEmail',
              otp_hash: 'otpHash',
              otp_hash_expires_at: 'otpHashExpiresAt',
              otp_method_last_used: 'otpMethodLastUsed',
              password_hash: 'passwordHash',
              phone_number: 'phoneNumber',
              phone_number_verified: 'phoneNumberVerified',
              ticket: 'ticket',
              ticket_expires_at: 'ticketExpiresAt',
              totp_secret: 'totpSecret',
              updated_at: 'updatedAt',
              webauthn_current_challenge: 'currentChallenge',
            },
            custom_name: 'users',
            custom_root_fields: {
              delete: 'deleteUsers',
              delete_by_pk: 'deleteUser',
              insert: 'insertUsers',
              insert_one: 'insertUser',
              select: 'users',
              select_aggregate: 'usersAggregate',
              select_by_pk: 'user',
              update: 'updateUsers',
              update_by_pk: 'updateUser',
            },
          },
          object_relationships: [
            {
              name: 'defaultRoleByRole',
              using: {
                foreign_key_constraint_on: 'default_role',
              },
            },
            {
              name: 'profile',
              using: {
                foreign_key_constraint_on: {
                  column: 'id',
                  table: {
                    schema: 'public',
                    name: 'profiles',
                  },
                },
              },
            },
          ],
          array_relationships: [
            {
              name: 'refreshTokens',
              using: {
                foreign_key_constraint_on: {
                  column: 'user_id',
                  table: {
                    name: 'refresh_tokens',
                    schema: 'auth',
                  },
                },
              },
            },
            {
              name: 'roles',
              using: {
                foreign_key_constraint_on: {
                  column: 'user_id',
                  table: {
                    name: 'user_roles',
                    schema: 'auth',
                  },
                },
              },
            },
            {
              name: 'securityKeys',
              using: {
                foreign_key_constraint_on: {
                  column: 'user_id',
                  table: {
                    name: 'user_security_keys',
                    schema: 'auth',
                  },
                },
              },
            },
            {
              name: 'userProviders',
              using: {
                foreign_key_constraint_on: {
                  column: 'user_id',
                  table: {
                    name: 'user_providers',
                    schema: 'auth',
                  },
                },
              },
            },
          ],
        },
      ],
    });

    await request.post('/change-env').send({
      AUTH_DISABLE_NEW_USERS: false,
      AUTH_EMAIL_SIGNIN_EMAIL_VERIFIED_REQUIRED: false,
    });
  }, 10000);

  afterAll(async () => {
    await overrideMetadata(authMetadataSource);
    await client.query(
      `DROP TABLE public.project_members;
        DROP TABLE public.profiles;
        DROP TABLE public.organisations;
        DROP TABLE public.projects;
        `
    );
    await client.end();
  });

  beforeEach(async () => {
    // await client.query(`DELETE FROM auth.users;`);
  });

  // * Insert a user with a profile and optional projects, and return their decoded JWT
  const insertUserProfile = async (
    projs?: string[],
    metaToken = faker.datatype.uuid()
  ) => {
    const email = faker.internet.email();
    const password = faker.internet.password();

    const {
      body: {
        session: { user },
      },
    } = await request.post('/signup/email-password').send({
      email,
      password,
      options: {
        metadata: {
          token: metaToken,
        },
      },
    });
    let query = `INSERT INTO public.profiles(id) VALUES('${user.id}');`;
    if (projs && projs.length > 0) {
      query += `INSERT INTO public.project_members(project_id, user_id) VALUES ${projs
        .map((id) => "('" + id + "', '" + user.id + "')")
        .join(',')};`;
    }
    await client.query(query);
    const { body } = await request.post('/signin/email-password').send({
      email,
      password,
    });

    expect(body).toBeDefined();
    const jwt = await decodeAccessToken(body.session.accessToken);

    expect(jwt).not.toBeNull();
    expect(jwt!['https://hasura.io/jwt/claims']).toBeObject();
    return jwt!;
  };

  it('should add a custom claim from a nested object relationship', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS: '{"organisation-id":"profile.organisation.id"}',
    });
    const email = faker.internet.email();
    const password = faker.internet.password();

    const {
      body: {
        session: { user },
      },
    } = await request.post('/signup/email-password').send({
      email,
      password,
    });

    await client.query(
      `INSERT INTO public.profiles(id, organisation_id) VALUES('${user.id}', '${organisationId}');`
    );
    const { body } = await request.post('/signin/email-password').send({
      email,
      password,
    });

    const jwt = await decodeAccessToken(body.session.accessToken);
    expect(jwt).toBeObject();
    if (jwt) {
      expect(jwt['https://hasura.io/jwt/claims']).toBeObject();
      expect(
        jwt['https://hasura.io/jwt/claims']['x-hasura-organisation-id']
      ).toEqual(organisationId);
    }
  });

  it('should add a custom claim from a nested array relationship', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS:
        '{"project-ids":"profile.contributesTo[].project.id"}',
    });

    const jwt = await insertUserProfile(projects);

    expect(jwt['https://hasura.io/jwt/claims']['x-hasura-project-ids']).toEqual(
      escapeValueToPg(projects)
    );
  });

  it('should add a custom claim from a field of the metadata column', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS: '{"token":"metadata.token"}',
    });
    const token = faker.datatype.uuid();
    const jwt = await insertUserProfile([], token);

    expect(jwt['https://hasura.io/jwt/claims']['x-hasura-token']).toEqual(
      token
    );
  });

  it('should add a custom claim from a nested array relationship while using the ambiguous non [] notation', async () => {
    // * Kept for backwards compatibility
    // * See https://docs.jsonata.org/predicate#singleton-array-and-value-equivalence
    // * Not recommended as it will return incorrect values when the array only has one element
    // * See https://github.com/nhost/hasura-auth/issues/222
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS:
        '{"project-ids":"profile.contributesTo[].project.id"}',
    });

    const jwt = await insertUserProfile(projects);

    expect(jwt['https://hasura.io/jwt/claims']['x-hasura-project-ids']).toEqual(
      escapeValueToPg(projects)
    );
  });

  it('should add a custom claim from a nested array relationship without element in array', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS:
        '{"project-ids":"profile.contributesTo[].project.id"}',
    });

    const jwt = await insertUserProfile();

    expect(jwt['https://hasura.io/jwt/claims']['x-hasura-project-ids']).toEqual(
      escapeValueToPg([])
    );
  });

  it('should add a custom claim from a nested array relationship with only one element in array', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS:
        '{"project-ids":"profile.contributesTo[].project.id"}',
    });
    const userProjects = projects.slice(0, 1);
    const jwt = await insertUserProfile(userProjects);

    expect(jwt['https://hasura.io/jwt/claims']['x-hasura-project-ids']).toEqual(
      escapeValueToPg(userProjects)
    );
  });

  it('should handle an invalid configuration (unparsable)', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS: '{"invalid-json": unquoted value }',
    });

    const jwt = await insertUserProfile();

    expect(jwt['https://hasura.io/jwt/claims']).toBeObject();
    expect(
      jwt['https://hasura.io/jwt/claims']['x-hasura-invalid-json']
    ).toBeUndefined();
  });

  it('should handle an invalid configuration (parsable, but not an object)', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS: 'string value',
    });
    const email = faker.internet.email();
    const password = faker.internet.password();
    const {
      body: {
        session: { user },
      },
    } = await request.post('/signup/email-password').send({
      email,
      password,
    });
    expect(user?.id).toBeString();
  });

  it('should handle an valid configuration with invalid GraphQL path', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS: '{"key": "path.does.not-exist" }',
    });
    const jwt = await insertUserProfile();
    expect(jwt['https://hasura.io/jwt/claims']).toBeObject();
    expect(jwt['https://hasura.io/jwt/claims']['x-hasura-key']).toBeUndefined();
  });

  it('should handle an valid configuration with invalid JSONata path', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS: '{"key": "invalid jsonata path!?!" }',
    });
    const jwt = await insertUserProfile();
    expect(jwt['https://hasura.io/jwt/claims']).toBeObject();
    expect(jwt['https://hasura.io/jwt/claims']['x-hasura-key']).toBeUndefined();
  });

  it('should handle a custom claim on a hard-coded JSON field', async () => {
    await request.post('/change-env').send({
      AUTH_JWT_CUSTOM_CLAIMS: '{"name": "metadata.name" }',
    });
    const email = faker.internet.email();
    const password = faker.internet.password();
    const name = faker.name.firstName();
    const {
      body: { session },
    } = await request.post('/signup/email-password').send({
      email,
      password,
      options: {
        metadata: { name },
      },
    });
    expect(session?.user?.id).toBeString();
    const jwt = await decodeAccessToken(session.accessToken);
    if (jwt) {
      expect(jwt['https://hasura.io/jwt/claims']).toBeObject();
      expect(jwt['https://hasura.io/jwt/claims']['x-hasura-name']).toEqual(
        name
      );
    }
  });
});
