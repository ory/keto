CREATE TABLE relation_tuples
(
    shard_id    varchar(64),
    object_id   varchar(64),
    relation    varchar(64),
    subject     varchar(128), /* can be object_id#relation or user_id */
    commit_time timestamp,
    PRIMARY KEY (shard_id, object_id, relation, subject, commit_time)
);

CREATE INDEX object_id_idx ON relation_tuples (object_id);

CREATE INDEX user_set_idx ON relation_tuples (object_id, relation);
