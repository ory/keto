CREATE TABLE keto_relation_tuples_uuid
(
    shard_id                 UUID        NOT NULL,
    nid                      UUID        NOT NULL,
    namespace                VARCHAR(200) NOT NULL,
    object                   UUID        NOT NULL,
    relation                 VARCHAR(64) NOT NULL,
    subject_id               UUID NULL,
    subject_set_namespace    VARCHAR(200) NULL,
    subject_set_object       UUID NULL,
    subject_set_relation     VARCHAR(64) NULL,
    commit_time              TIMESTAMP   NOT NULL,
    PRIMARY KEY (shard_id, nid),
    CONSTRAINT keto_relation_tuples_uuid_nid_fk FOREIGN KEY (nid) REFERENCES networks (id),
    CONSTRAINT chk_keto_rt_uuid_subject_type CHECK (((((subject_id IS NULL) AND (subject_set_namespace IS NOT NULL)) AND
                                                      (subject_set_object IS NOT NULL)) AND
                                                     (subject_set_relation IS NOT NULL)) OR
                                                    ((((subject_id IS NOT NULL) AND (subject_set_namespace IS NULL)) AND
                                                      (subject_set_object IS NULL)) AND (subject_set_relation IS NULL)))
);
CREATE INDEX keto_relation_tuples_uuid_subject_ids_idx ON keto_relation_tuples_uuid (nid, namespace, object, relation, subject_id) WHERE ((subject_set_namespace IS NULL) AND (subject_set_object IS NULL)) AND (subject_set_relation IS NULL);
CREATE INDEX keto_relation_tuples_uuid_subject_sets_idx ON keto_relation_tuples_uuid (nid, namespace, object, relation, subject_set_namespace, subject_set_object, subject_set_relation) WHERE subject_id IS NULL;
CREATE INDEX keto_relation_tuples_uuid_full_idx ON keto_relation_tuples_uuid (nid, namespace, object, relation, subject_id, subject_set_namespace, subject_set_object, subject_set_relation, commit_time);
CREATE INDEX keto_relation_tuples_uuid_reverse_subject_ids_idx ON keto_relation_tuples_uuid (nid, subject_id, relation, namespace) WHERE ((subject_set_namespace IS NULL) AND (subject_set_object IS NULL)) AND (subject_set_relation IS NULL);
CREATE INDEX keto_relation_tuples_uuid_reverse_subject_sets_idx ON keto_relation_tuples_uuid (nid, subject_set_namespace, subject_set_object, subject_set_relation, relation, namespace) WHERE subject_id IS NULL;
