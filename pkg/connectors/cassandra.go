package connectors

import (
	"fmt"

	"github.com/codingconcepts/oates/pkg/models"
	"github.com/gocql/gocql"
)

// CassandraConnector is an implementation of ConnectFetcher that operates against
// a Cassandra database.
type CassandraConnector struct {
	session  *gocql.Session
	Host     string `yaml:"host"`
	Keyspace string `yaml:"keyspace"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Connect opens a Cassandra session.
func (cf *CassandraConnector) Connect() error {
	cluster := gocql.NewCluster(cf.Host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cf.Username,
		Password: cf.Password,
	}

	cluster.Consistency = gocql.LocalQuorum
	cluster.Keyspace = cf.Keyspace

	var err error
	if cf.session, err = cluster.CreateSession(); err != nil {
		return fmt.Errorf("creating cassandra session: %w", err)
	}

	return nil
}

// Fetch fetches schema information from a configured keyspace.
func (cf *CassandraConnector) Fetch() ([]models.Metadata, error) {
	stmt := `SELECT
		"keyspace_name",
		"table_name",
		"column_name",
		"type"
	FROM system_schema.columns WHERE keyspace_name = ?`

	var columns []models.Metadata
	var eg models.Metadata

	iter := cf.session.Query(stmt, cf.Keyspace).Iter()
	for iter.Scan(&eg.Database, &eg.Table, &eg.Column, &eg.Type) {
		columns = append(columns, eg)
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("scanning rows: %w", err)
	}

	return columns, nil
}
