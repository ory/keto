CREATE TABLE keto_uuid_mappings
(
    id                       UUID NOT NULL PRIMARY KEY,
    string_representation    TEXT NOT NULL CHECK (string_representation <> '')
);
