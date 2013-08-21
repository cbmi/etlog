package stores

// Store for text files. The location of the data is defined by lines
// and chars.
type Text struct {
    File

    // The line, line range or series of lines and ranges.
    Line string `json:"line,omitempty" bson:"line,omitempty"`
}
