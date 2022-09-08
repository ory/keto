ALTER TABLE keto_uuid_mappings ADD CONSTRAINT keto_uuid_mappings_string_representation_check CHECK (string_representation <> '');
