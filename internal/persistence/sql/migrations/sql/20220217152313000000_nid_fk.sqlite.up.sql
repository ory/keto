-- These tuples can't be retrieved over any API (and also not inserted by any API) because we always add a `WHERE nid = ?` clause to the queries.
-- This DELETE statement is here to ensure that we don't run into problems where the `nid` column is not properly set. Should not happen if only the official API is used.
DELETE FROM keto_relation_tuples WHERE nid NOT IN (SELECT nid FROM networks);

CREATE TABLE keto_relation_tuples_with_fk
(
    shard_id                 TEXT        NOT NULL,
    nid                      TEXT        NOT NULL,
    namespace_id             INTEGER     NOT NULL,
    object                   VARCHAR(64) NOT NULL,
    relation                 VARCHAR(64) NOT NULL,
    subject_id               VARCHAR(64) NULL,
    subject_set_namespace_id INTEGER NULL,
    subject_set_object       VARCHAR(64) NULL,
    subject_set_relation     VARCHAR(64) NULL,
    commit_time              TIMESTAMP   NOT NULL,

    PRIMARY KEY (shard_id, nid),

    CONSTRAINT keto_relation_tuples_nid_fk FOREIGN KEY (nid) REFERENCES networks (id),

    -- enforce to have exactly one of subject_id or subject_set
    CONSTRAINT chk_keto_rt_subject_type CHECK
        ((subject_id IS NULL AND
          subject_set_namespace_id IS NOT NULL AND subject_set_object IS NOT NULL AND subject_set_relation IS NOT NULL)
            OR
         (subject_id IS NOT NULL AND
          subject_set_namespace_id IS NULL AND subject_set_object IS NULL AND subject_set_relation IS NULL))
);

INSERT INTO keto_relation_tuples_with_fk (shard_id, nid, namespace_id, object, relation, subject_id, subject_set_namespace_id, subject_set_object, subject_set_relation, commit_time) SELECT * FROM keto_relation_tuples;

DROP TABLE keto_relation_tuples;

ALTER TABLE keto_relation_tuples_with_fk RENAME TO keto_relation_tuples;

CREATE INDEX keto_relation_tuples_subject_ids_idx ON keto_relation_tuples (nid,
                                                                           namespace_id,
                                                                           object,
                                                                           relation,
                                                                           subject_id) WHERE subject_set_namespace_id IS NULL AND subject_set_object IS NULL AND subject_set_relation IS NULL;

CREATE INDEX keto_relation_tuples_subject_sets_idx ON keto_relation_tuples (nid,
                                                                            namespace_id,
                                                                            object,
                                                                            relation,
                                                                            subject_set_namespace_id,
                                                                            subject_set_object,
                                                                            subject_set_relation) WHERE subject_id IS NULL;

CREATE INDEX keto_relation_tuples_full_idx ON keto_relation_tuples (nid,
                                                                    namespace_id,
                                                                    object,
                                                                    relation,
                                                                    subject_id,
                                                                    subject_set_namespace_id,
                                                                    subject_set_object,
                                                                    subject_set_relation,
                                                                    commit_time);
