CREATE TABLE keto_relation_tuples
(
    shard_id                 CHAR(36)    NOT NULL,
    nid                      CHAR(36)    NOT NULL,
    namespace_id             INT8        NOT NULL,
    object                   CHAR(36)    NOT NULL,
    relation                 VARCHAR(64) NOT NULL,
    subject_id               CHAR(36) NULL,
    subject_set_namespace_id INT8 NULL,
    subject_set_object       CHAR(36) NULL,
    subject_set_relation     VARCHAR(64) NULL,
    commit_time              TIMESTAMP   NOT NULL,
    PRIMARY KEY (shard_id ASC, nid ASC),
    CONSTRAINT keto_relation_tuples_nid_fk FOREIGN KEY (nid) REFERENCES networks (id),

    -- mysql has no partial indexes so we can only use the full one
    INDEX                    keto_relation_tuples_full_idx (nid, namespace_id, object, relation, subject_id, subject_set_namespace_id, subject_set_object, subject_set_relation, commit_time),
    INDEX                    keto_relation_tuples_reverse_subject_idx (nid, subject_id, subject_set_namespace_id, subject_set_object, subject_set_relation, relation, namespace_id),

    CONSTRAINT chk_keto_rt_subject_type CHECK (((((subject_id IS NULL) AND (subject_set_namespace_id IS NOT NULL)) AND
                                                      (subject_set_object IS NOT NULL)) AND
                                                     (subject_set_relation IS NOT NULL)) OR
                                                    ((((subject_id IS NOT NULL) AND (subject_set_namespace_id IS NULL)) AND
                                                      (subject_set_object IS NULL)) AND (subject_set_relation IS NULL)))
);
