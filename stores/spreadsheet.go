package stores

// Store for spreadsheet files such as Microsoft Excel, Google Docs
// Spreadsheet, and OpenOffice Spreadsheet. Data is tabular.
type Spreadsheet struct {
    File

    // The index or name of the sheet
    Sheet string `json:"sheet",bson:"sheet"`

	// A line number or range of lines representing the header.
	Header string `json:"header",bson:"header"`

    // The line, line range or series of lines and ranges.
    Row string `json:"row",bson:"row"`

	// The column index or name (if a header exists) where the data is
	// defined. For indexes, this can be a range. For column names this
	// can be a comma-separated list of names.
	Column string `json:"column",bson:"column"`
}
