ALTER TABLE keto_uuid_mappings ADD CONSTRAINT check_string_representation CHECK (string_representation <> '');
