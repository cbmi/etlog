package stores

// Store for documents stored in MongoDB databases
type MongoDocument struct {
    Document

    // The database name. If not supplied the database name will attempt to be
    // extracted from the uri
    Database string `json:"database,omitempty" bson:"database,omitempty"`

    // The collection name where the document is stored.
    Collection string `json:"collection,omitempty" bson:"collection,omitempty"`
}
