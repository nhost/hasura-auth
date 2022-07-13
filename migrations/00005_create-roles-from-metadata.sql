CREATE OR REPLACE FUNCTION auth.update_roles_from_metadata() RETURNS trigger LANGUAGE plpgsql AS $$ BEGIN
INSERT INTO auth.roles(role)
SELECT DISTINCT jsonb_array_elements_text(
        jsonb_path_query_array(
            new.metadata::jsonb,
            '$.sources[*].tables[*].select_permissions[*].role'
        ) || jsonb_path_query_array(
            new.metadata::jsonb,
            '$.sources[*].tables[*].insert_permissions[*].role'
        ) || jsonb_path_query_array(
            new.metadata::jsonb,
            '$.sources[*].tables[*].update_permissions[*].role'
        ) || jsonb_path_query_array(
            new.metadata::jsonb,
            '$.sources[*].tables[*].delete_permissions[*].role'
        )
    ) as role ON CONFLICT DO NOTHING;
RETURN new;
END;
$$;
DROP TRIGGER IF EXISTS update_roles_from_metadata ON hdb_catalog.hdb_metadata;
CREATE TRIGGER update_roles_from_metadata BEFORE
UPDATE ON hdb_catalog.hdb_metadata FOR EACH ROW EXECUTE FUNCTION auth.update_roles_from_metadata ();