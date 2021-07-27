CREATE TABLE "keto_namespace_ids" (
"id" UUID NOT NULL,
"nid" UUID NOT NULL,
"serial_id" integer NOT NULL,
"created_at" timestamp NOT NULL,
"updated_at" timestamp NOT NULL,
PRIMARY KEY("nid", "id")
);