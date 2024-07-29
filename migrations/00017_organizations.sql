CREATE table auth.organizations (
  id uuid DEFAULT public.gen_random_uuid () NOT NULL PRIMARY KEY,
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone DEFAULT now() NOT NULL,
  name text NOT NULL UNIQUE,
  metadata jsonb
);


CREATE table auth.organizations_roles (
  value text PRIMARY KEY,
  comment text
);


INSERT INTO auth.organizations_roles
  (value, comment)
VALUES
  ('admin', 'Admin'),
  ('user', 'User');


CREATE table auth.organizations_members (
  organization_id uuid NOT NULL REFERENCES auth.organizations (id) ON DELETE CASCADE ON UPDATE CASCADE,
  user_id uuid NOT NULL REFERENCES auth.users (id) ON DELETE CASCADE ON UPDATE CASCADE UNIQUE,
  role text NOT NULL REFERENCES auth.organizations_roles (value),
  PRIMARY KEY (organization_id, user_id)
);


CREATE table auth.organizations_invites (
    id uuid DEFAULT public.gen_random_uuid () NOT NULL PRIMARY KEY,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    organization_id uuid NOT NULL REFERENCES auth.organizations (id),
    email text NOT NULL,
    role text NOT NULL REFERENCES auth.organizations_roles (value)
);
