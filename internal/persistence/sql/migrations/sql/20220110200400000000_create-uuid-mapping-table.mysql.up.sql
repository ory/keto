CREATE TABLE keto_uuid_mappings
(
    id                       VARCHAR(64) NOT NULL,
    string_representation    VARCHAR(64) NOT NULL CHECK (string_representation <> ''),

    PRIMARY KEY (id),

    -- enforce uniqueness
    CONSTRAINT chk_keto_uuid_map_uniq UNIQUE (id, string_representation)
);