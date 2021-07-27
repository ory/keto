
CREATE INDEX keto_relation_tuples_subject_ids_idx ON keto_relation_tuples (nid,
                                                                           namespace_id,
                                                                           object,
                                                                           relation,
                                                                           subject_id) WHERE subject_set_namespace_id IS NULL AND subject_set_object IS NULL AND subject_set_relation IS NULL;