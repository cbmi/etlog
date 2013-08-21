package stores

// Store for JSON-based data
type Document struct {
    Store

    // A forward slash-delimited path to the value
    Path string `json:"path,omitempty" bson:"path,omitempty"`
}
