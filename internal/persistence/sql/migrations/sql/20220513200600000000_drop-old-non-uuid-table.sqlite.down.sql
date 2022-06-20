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
    CONSTRAINT "primary" PRIMARY KEY (shard_id ASC, nid ASC),
    CONSTRAINT keto_relation_tuples_nid_fk FOREIGN KEY (nid) REFERENCES networks (id),
    CONSTRAINT chk_keto_rt_subject_type CHECK (((((subject_id IS NULL) AND (subject_set_namespace_id IS NOT NULL)) AND
                                                 (subject_set_object IS NOT NULL)) AND
                                                (subject_set_relation IS NOT NULL)) OR
                                               ((((subject_id IS NOT NULL) AND (subject_set_namespace_id IS NULL)) AND
                                                 (subject_set_object IS NULL)) AND (subject_set_relation IS NULL)))
);

CREATE INDEX keto_relation_tuples_subject_ids_idx ON keto_relation_tuples (nid ASC, namespace_id ASC, object ASC, relation ASC, subject_id ASC) WHERE ((subject_set_namespace_id IS NULL) AND (subject_set_object IS NULL)) AND (subject_set_relation IS NULL);
CREATE INDEX keto_relation_tuples_subject_sets_idx ON keto_relation_tuples (nid ASC, namespace_id ASC, object ASC,
                                                                            relation ASC, subject_set_namespace_id ASC,
                                                                            subject_set_object ASC, subject_set_relation
                                                                            ASC) WHERE subject_id IS NULL;
CREATE INDEX keto_relation_tuples_full_idx ON keto_relation_tuples (nid ASC, namespace_id ASC, object ASC, relation ASC,
                                                                    subject_id ASC, subject_set_namespace_id ASC,
                                                                    subject_set_object ASC, subject_set_relation ASC,
                                                                    commit_time ASC);
CREATE INDEX keto_relation_tuples_reverse_subject_ids_idx ON keto_relation_tuples (nid ASC, subject_id ASC, relation ASC, namespace_id ASC) WHERE ((subject_set_namespace_id IS NULL) AND (subject_set_object IS NULL)) AND (subject_set_relation IS NULL);
CREATE INDEX keto_relation_tuples_reverse_subject_sets_idx ON keto_relation_tuples (nid ASC, subject_set_namespace_id
                                                                                    ASC, subject_set_object ASC,
                                                                                    subject_set_relation ASC, relation
                                                                                    ASC, namespace_id
                                                                                    ASC) WHERE subject_id IS NULL;
