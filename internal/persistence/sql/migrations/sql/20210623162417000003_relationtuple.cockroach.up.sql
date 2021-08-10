
CREATE INDEX keto_relation_tuples_full_idx ON keto_relation_tuples (nid,
                                                                    namespace_id,
                                                                    object,
                                                                    relation,
                                                                    subject_id,
                                                                    subject_set_namespace_id,
                                                                    subject_set_object,
                                                                    subject_set_relation,
                                                                    commit_time);