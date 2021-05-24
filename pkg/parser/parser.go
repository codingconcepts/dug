package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/codingconcepts/oates/pkg/connectors"
	"github.com/codingconcepts/oates/pkg/models"
	"gopkg.in/yaml.v2"
)

// ParseConnectors loads a file from disk and returns a collection ConnectFetches,
// that are ready to connect and start fetching schema inforamtion.
func ParseConnectors(path string) ([]models.ConnectFetcher, error) {
	// Read the contents of the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	// Parse the contents of the file into a top-level sources object.
	var s models.Sources
	if err = yaml.Unmarshal(content, &s); err != nil {
		return nil, fmt.Errorf("unmarshalling config file: %w", err)
	}

	var cfs []models.ConnectFetcher
	for _, source := range s.Sources {
		switch strings.ToLower(source.Type) {
		case "cassandra":
			var connector connectors.CassandraConnector
			if err := source.Connector.Unmarshal(&connector); err != nil {
				return nil, fmt.Errorf("parsing cassandra connector: %w", err)
			}
			cfs = append(cfs, &connector)
		}
	}

	return cfs, nil
}
