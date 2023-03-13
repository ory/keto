ALTER TABLE
  keto_relation_tuples DROP CONSTRAINT keto_relation_tuples_uuid_nid_fk,
ADD
  CONSTRAINT keto_relation_tuples_nid_fk FOREIGN KEY (nid) REFERENCES networks (id) ON UPDATE RESTRICT ON DELETE CASCADE;