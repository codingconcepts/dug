package models

// Column describes a column in a database schema.
type Column struct {
	Database string `json:"database"`
	Table    string `json:"table"`
	Column   string `json:"column"`
	Type     string `json:"data_type"`
}
