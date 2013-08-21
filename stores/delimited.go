package stores

// Store for delimited text files. The location is denoted by the line and
// column fields. Examples includes CSV and tab-delimited files.
type Delimited struct {
	Text

	// The delimiter used which denotes the structure
	Delimiter string `json:"delimiter,omitempty" bson:"delimiter,omitempty"`

	// A line number or range of lines representing the header. Note, this
	// should only be supplied if the columns are guaranteed consistent for
	// all rows in the file.
	Header string `json:"header,omitempty" bson:"header,omitempty"`

	// The column index or name (if a header exists) where the data is
	// defined. For indexes, this can be a range. For column names this
	// can be a comma-separated list of names.
	Column string `json:"column,omitempty" bson:"column,omitempty"`
}
