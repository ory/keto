CREATE TABLE keto_relation_tuples_uuid
(
    shard_id                 UUID        NOT NULL,
    nid                      UUID        NOT NULL,
    namespace_id             INT8        NOT NULL,
    object                   UUID        NOT NULL,
    relation                 VARCHAR(64) NOT NULL,
    subject_id               UUID NULL,
    subject_set_namespace_id INT8 NULL,
    subject_set_object       UUID NULL,
    subject_set_relation     VARCHAR(64) NULL,
    commit_time              TIMESTAMP   NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (shard_id ASC, nid ASC),
    CONSTRAINT keto_relation_tuples_nid_fk FOREIGN KEY (nid) REFERENCES public.networks (id),
    INDEX                    keto_relation_tuples_subject_ids_idx (nid ASC, namespace_id ASC, object ASC, relation ASC, subject_id ASC) WHERE ((subject_set_namespace_id IS NULL) AND (subject_set_object IS NULL)) AND (subject_set_relation IS NULL),
    INDEX                    keto_relation_tuples_subject_sets_idx (nid ASC, namespace_id ASC, object ASC, relation ASC, subject_set_namespace_id ASC, subject_set_object ASC, subject_set_relation ASC) WHERE subject_id IS NULL,
    INDEX                    keto_relation_tuples_full_idx (nid ASC, namespace_id ASC, object ASC, relation ASC, subject_id ASC, subject_set_namespace_id ASC, subject_set_object ASC, subject_set_relation ASC, commit_time ASC),
    INDEX                    keto_relation_tuples_reverse_subject_ids_idx (nid ASC, subject_id ASC, relation ASC, namespace_id ASC) WHERE ((subject_set_namespace_id IS NULL) AND (subject_set_object IS NULL)) AND (subject_set_relation IS NULL),
    INDEX                    keto_relation_tuples_reverse_subject_sets_idx (nid ASC, subject_set_namespace_id ASC, subject_set_object ASC, subject_set_relation ASC, relation ASC, namespace_id ASC) WHERE subject_id IS NULL,
    FAMILY                   "primary" (shard_id, nid, namespace_id, object, relation, subject_id, subject_set_namespace_id, subject_set_object, subject_set_relation, commit_time),
    CONSTRAINT chk_keto_rt_subject_type CHECK (((((subject_id IS NULL) AND (subject_set_namespace_id IS NOT NULL)) AND
                                                 (subject_set_object IS NOT NULL)) AND
                                                (subject_set_relation IS NOT NULL)) OR
                                               ((((subject_id IS NOT NULL) AND (subject_set_namespace_id IS NULL)) AND
                                                 (subject_set_object IS NULL)) AND (subject_set_relation IS NULL)))
);
