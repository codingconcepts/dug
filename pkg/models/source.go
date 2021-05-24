package models

// Sources is the top-level of a configuration file and contains a collection
// of Source objects.
type Sources struct {
	Sources []Source `yaml:"sources"`
}

// Source is the configuration for a single discoverer for a given database source.
type Source struct {
	Name      string     `yaml:"name"`
	Type      string     `yaml:"type"`
	Connector RawMessage `yaml:"connector"`
}

// ConnectFetcher describes the behaviour of a source connector. It must be able
// connect to its database and fetch column information.
type ConnectFetcher interface {
	Connect() error
	Fetch() ([]Metadata, error)
}
