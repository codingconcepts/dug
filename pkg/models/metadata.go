package models

// Metadata describes a piece of information in a database table.
type Metadata struct {
	Database string `json:"database"`
	Table    string `json:"table"`
	Column   string `json:"column"`
	Type     string `json:"data_type"`
}
