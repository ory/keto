CREATE INDEX keto_relation_tuples_reverse_subject_ids_idx ON keto_relation_tuples (nid,
                                                                                   subject_id,
                                                                                   relation,
                                                                                   namespace_id
    ) WHERE subject_set_namespace_id IS NULL AND subject_set_object IS NULL AND subject_set_relation IS NULL;

CREATE INDEX keto_relation_tuples_reverse_subject_sets_idx ON keto_relation_tuples (nid,
                                                                                    subject_set_namespace_id,
                                                                                    subject_set_object,
                                                                                    subject_set_relation,
                                                                                    relation,
                                                                                    namespace_id
    ) WHERE subject_id IS NULL;
