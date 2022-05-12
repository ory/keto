CREATE INDEX keto_relation_tuples_reverse_subject_idx ON keto_relation_tuples (nid,
                                                                               subject_id,
                                                                               subject_set_namespace_id,
                                                                               subject_set_object,
                                                                               subject_set_relation,
                                                                               relation,
                                                                               namespace_id
    );
