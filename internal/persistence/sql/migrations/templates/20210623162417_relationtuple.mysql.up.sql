-- we can not fizz this migration because there is no check constraint support in fizz
-- and it gets cumbersome with sqlite; having a working SQL version is actually way easier
CREATE TABLE keto_relation_tuples
(
    shard_id                 char(36)    NOT NULL,
    nid                      char(36)    NOT NULL,
    namespace_id             INTEGER     NOT NULL,
    object                   VARCHAR(64) NOT NULL,
    relation                 VARCHAR(64) NOT NULL,
    subject_id               VARCHAR(64) NULL,
    subject_set_namespace_id INTEGER NULL,
    subject_set_object       VARCHAR(64) NULL,
    subject_set_relation     VARCHAR(64) NULL,
    commit_time              TIMESTAMP   NOT NULL,

    PRIMARY KEY (shard_id, nid),

    -- enforce to have exactly one of subject_id or subject_set
    CONSTRAINT chk_keto_rt_subject_type CHECK
        ((subject_id = null AND
          NOT subject_set_namespace_id = null AND NOT subject_set_object = null AND NOT subject_set_relation = null)
            OR
         (NOT subject_id = null AND
          subject_set_namespace_id = null AND subject_set_object = null AND subject_set_relation = null))
);
