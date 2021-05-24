# oates
A dead simple data discovery tool.

## Installation

From source using the Go toolchain:
```
$ go get -u github.com/codingconcepts/oates
```

## Usage

These steps use the _examples/cassandra example.

**Step 1** Create a database:

```
$ docker run -it --rm --name cassandra -p 9042:9042 bitnami/cassandra:latest
```

**Step 2** Create a database using the CQL in _examples/cassandra/create.sql.

**Step 3** Run oates:

```
$ oates -config _examples/cassandra/config.yaml
&{session:0xc000091180 Host:localhost:9042 Keyspace:my_keyspace Username:cassandra Password:cassandra}
        {Database:my_keyspace Table:table_one Column:date_of_birth Type:timestamp}
        {Database:my_keyspace Table:table_one Column:first_name Type:text}
        {Database:my_keyspace Table:table_one Column:id Type:uuid}
        {Database:my_keyspace Table:table_one Column:last_name Type:text}
        {Database:my_keyspace Table:table_one Column:metadata Type:blob}
```