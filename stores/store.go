package stores

import (
    "etlog/encoding"
    "labix.org/v2/mgo/bson"
)

// Base which contains general information about the store
type Store struct {
    // Represents the Mongo ID
    Id bson.ObjectId `json:"-" bson:"_id,omitempty"`

	// URI which denotes where the data lives, where came from (streamed) or
	// where can be accessed in the future. The latter value is preferred
	// since it theoretically enables future access. The completeness and
	// value of this is store dependent.
    Uri string `json:"uri,omitempty" bson:"uri,omitempty"`

	// The store type as a string. This is used when messages are parsed and
	// for downstream processing. If not defined, the type will attempted to
	// be inferred by keys supplied.
	Type string `json:"type,omitempty" bson:"type,omitempty"`

	// The value or array of values processed. For sources this would
	// typically be the pre-transformed data while for targets this would be
	// post-processed data. Supplying the value here is typically unnecessary
	// since the value could be accessed using the other information supplied
	// in store data. However, for systems that treat this as an audit log or
	// if the target system does not perform an versioning of data of it's
	// own, this could act as primitive store of values.
    Value interface{} `json:"value,omitempty" bson:"value,omitempty"`

	// Extra data from decoded JSON. This enables clients to store additional
	// metadata about the store.
	Extra map[string]interface{} `json:"-" bson:",inline,omitempty"`
}

// Satisfies the `json.Marshaler` interface
func (s *Store) MarshalJSON() ([]byte, error) {
    return encoding.MarshalJSON(s)
}

// Satisfies the `json.Unmarshaler` interface
func (s *Store) UnmarshalJSON(b []byte) error {
    return encoding.UnmarshalJSON(b, s)
}
