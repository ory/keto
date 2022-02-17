-- These tuples can't be retrieved over any API (and also not inserted by any API) because we always add a `WHERE nid = ?` clause to the queries.
-- This DELETE statement is here to ensure that we don't run into problems where the `nid` column is not properly set. Should not happen if only the official API is used.
DELETE FROM keto_relation_tuples WHERE nid NOT IN (SELECT nid FROM networks);

ALTER TABLE keto_relation_tuples
    ADD CONSTRAINT keto_relation_tuples_nid_fk
        FOREIGN KEY (nid) REFERENCES networks (id);
