CREATE TABLE keto_uuid_mappings
(
    id                       CHAR(36) NOT NULL,
    string_representation    TEXT        NOT NULL CHECK (string_representation <> ''),

    PRIMARY KEY (id)
);
