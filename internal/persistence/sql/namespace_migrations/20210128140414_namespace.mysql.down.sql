DROP INDEX {{ identifier .Parameters.tableName }}_user_set_idx ON {{ identifier .Parameters.tableName }};

DROP INDEX {{ identifier .Parameters.tableName }}_object_idx ON {{ identifier .Parameters.tableName }};

DROP TABLE {{ identifier .Parameters.tableName }};
