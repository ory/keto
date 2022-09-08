ALTER TABLE keto_uuid_mappings ADD CONSTRAINT keto_uuid_mappings_chk_1 CHECK (string_representation <> '');
