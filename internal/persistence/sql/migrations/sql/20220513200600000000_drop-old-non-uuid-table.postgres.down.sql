CREATE TABLE keto_relation_tuples
(
    shard_id                 UUID        NOT NULL,
    nid                      UUID        NOT NULL,
    namespace_id             INT8        NOT NULL,
    object                   VARCHAR(64) NOT NULL,
    relation                 VARCHAR(64) NOT NULL,
    subject_id               VARCHAR(64) NULL,
    subject_set_namespace_id INT8 NULL,
    subject_set_object       VARCHAR(64) NULL,
    subject_set_relation     VARCHAR(64) NULL,
    commit_time              TIMESTAMP   NOT NULL,
    PRIMARY KEY (shard_id, nid),
    CONSTRAINT keto_relation_tuples_nid_fk FOREIGN KEY (nid) REFERENCES networks (id),
    CONSTRAINT chk_keto_rt_subject_type CHECK (((((subject_id IS NULL) AND (subject_set_namespace_id IS NOT NULL)) AND
                                                 (subject_set_object IS NOT NULL)) AND
                                                (subject_set_relation IS NOT NULL)) OR
                                               ((((subject_id IS NOT NULL) AND (subject_set_namespace_id IS NULL)) AND
                                                 (subject_set_object IS NULL)) AND (subject_set_relation IS NULL)))
);
CREATE INDEX keto_relation_tuples_subject_ids_idx ON keto_relation_tuples (nid, namespace_id, object, relation, subject_id) WHERE ((subject_set_namespace_id IS NULL) AND (subject_set_object IS NULL)) AND (subject_set_relation IS NULL);
CREATE INDEX keto_relation_tuples_subject_sets_idx ON keto_relation_tuples (nid, namespace_id, object, relation,
                                                                            subject_set_namespace_id,
                                                                            subject_set_object,
                                                                            subject_set_relation) WHERE subject_id IS NULL;
CREATE INDEX keto_relation_tuples_full_idx ON keto_relation_tuples (nid, namespace_id, object, relation, subject_id,
                                                                    subject_set_namespace_id, subject_set_object,
                                                                    subject_set_relation, commit_time);
CREATE INDEX keto_relation_tuples_reverse_subject_ids_idx ON keto_relation_tuples (nid, subject_id, relation, namespace_id) WHERE ((subject_set_namespace_id IS NULL) AND (subject_set_object IS NULL)) AND (subject_set_relation IS NULL);
CREATE INDEX keto_relation_tuples_reverse_subject_sets_idx ON keto_relation_tuples (nid, subject_set_namespace_id,
                                                                                    subject_set_object,
                                                                                    subject_set_relation, relation,
                                                                                    namespace_id) WHERE subject_id IS NULL;

