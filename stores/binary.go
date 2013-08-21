package stores

// Store for binary files or text files which are not line-based. The
// location of the data is defined by byte positions and ranges.
type Binary struct {
    File

    // The byte position, range or series of bytes and ranges.
    Bytes string `json:"bytes,omitempty" bson:"bytes,omitempty"`
}
