
CREATE INDEX keto_relation_tuples_subject_sets_idx ON keto_relation_tuples (nid,
                                                                            namespace_id,
                                                                            object,
                                                                            relation,
                                                                            subject_set_namespace_id,
                                                                            subject_set_object,
                                                                            subject_set_relation) WHERE subject_id IS NULL;