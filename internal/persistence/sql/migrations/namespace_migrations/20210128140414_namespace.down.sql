DROP INDEX {{ identifier .Parameters.tableName }}_user_set_idx;

DROP INDEX {{ identifier .Parameters.tableName }}_object_idx;

DROP TABLE {{ identifier .Parameters.tableName }};
