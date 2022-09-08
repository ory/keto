ALTER TABLE keto_uuid_mappings ADD CONSTRAINT IF NOT EXISTS check_string_representation CHECK (string_representation <> '');
