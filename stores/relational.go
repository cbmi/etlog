package stores

// Store representing a data based on the relational model. Examples include
// PostgreSQL, SQLite, Oracle, MySQL.
type Relational struct {
	Store

	// The database name
	Database string `json:"database,omitempty" bson:"database,omitempty"`

	// The schema name for databases that support them, e.g. PostgreSQL
	Schema string `json:"schema,omitempty" bson:"schema,omitempty"`

	// The table name
	Table string `json:"table,omitempty" bson:"table,omitempty"`

	// Object containing the lookup of the row
	Row map[string]interface{} `json:"row,omitempty" bson:"row,omitempty"`

	// The column name
	Column string `json:"column,omitempty" column:"column,omitempty"`
}
