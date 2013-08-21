package stores

// The base type for file-based stores
type File struct {
    Store

    // The name of the file. If the uri is supplied with a path, the name of
    // the file will be extracted if not supplied.
    Name string `json:"name,omitempty" bson:"name,omitempty"`
}
