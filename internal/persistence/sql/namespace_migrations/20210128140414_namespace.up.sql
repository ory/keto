CREATE TABLE {{ identifier .Parameters.tableName }}
(
    shard_id    varchar(64),
    object      varchar(64),
    relation    varchar(64),
    subject     varchar(256), /* can be <namespace:object#relation> or <user_id> */
    commit_time timestamp,

    PRIMARY KEY (shard_id, object, relation, subject, commit_time)
);

CREATE INDEX {{ identifier .Parameters.tableName }}_object_idx ON {{ identifier .Parameters.tableName }} (object);

CREATE INDEX {{ identifier .Parameters.tableName }}_user_set_idx ON {{ identifier .Parameters.tableName }} (object, relation);
