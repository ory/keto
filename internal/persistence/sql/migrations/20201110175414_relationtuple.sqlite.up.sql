CREATE TABLE keto_namespace
(
    id             INTEGER PRIMARY KEY,
    name           VARCHAR,
    schema_version INTEGER NOT NULL
);

CREATE UNIQUE INDEX keto_namespace_names ON keto_namespace (name);
