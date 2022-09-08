--ALTER TABLE keto_uuid_mappings DROP CONSTRAINT IF EXISTS check_string_representation;
ALTER TABLE keto_uuid_mappings RENAME TO keto_uuid_mapping_with_old_check_constraint;
CREATE TABLE keto_uuid_mappings
(
    id                       UUID NOT NULL PRIMARY KEY,
    string_representation    TEXT NOT NULL
);
INSERT INTO keto_uuid_mappings SELECT * FROM keto_uuid_mapping_with_old_check_constraint;
DROP TABLE keto_uuid_mapping_with_old_check_constraint;
