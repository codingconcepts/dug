-- Use this script to create the database that will be used in this example.
CREATE KEYSPACE IF NOT EXISTS my_keyspace
    WITH REPLICATION = {
        'class' : 'SimpleStrategy',
        'replication_factor' : 1
    };

CREATE TABLE IF NOT EXISTS my_keyspace.table_one (
	"id" uuid,
	"first_name" text,
	"last_name" text,
	"date_of_birth" timestamp,
	"metadata" blob,
	PRIMARY KEY (("id"), "first_name", "last_name")
);

SELECT * FROM system_schema.columns WHERE keyspace_name = 'my_keyspace';