package main

import (
	"time"
    "labix.org/v2/mgo/bson"
    "etlog/stores"
    "etlog/encoding"
)

type Transaction struct {
    Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`

	// Timestamp of when the transaction occurred
	Timestamp time.Time `json:"timestamp,omitempty" bson:"timestamp,omitempty"`

	// The action that was performed on the target. This may not be
    // applicable or available depending on what operation the script is
    // performing. The value is generally target-specific; for example,
    // insert and delete for row-based operations, append for file-based
    // writing, pop for a Redis list, etc.
	Action string `json:"action,omitempty" bson:"action,omitempty"`

	// An object representing the script used that performed the ETL and
	// produced this transaction.
	Script Script `json:"script,omitempty" bson:"script,omitempty"`

	// An object or array of objects representing the sources of data being
	// used in the target output.
	Source stores.Store `json:"source,omitempty" bson:"source,omitempty"`

	// An object or array of objects representing the targets
	Target stores.Store `json:"target,omitempty" bson:"target,omitempty"`

	// Extra data from decoded JSON. This enables clients to store additional
	// metadata about the store.
	Extra map[string]interface{} `json:"-" bson:",inline,omitempty"`
}

// Satisfies the `json.Marshaler` interface
func (t *Transaction) MarshalJSON() ([]byte, error) {
    return encoding.MarshalJSON(t)
}

// Satisfies the `json.Unmarshaler` interface
func (t *Transaction) UnmarshalJSON(b []byte) error {
    return encoding.UnmarshalJSON(b, t)
}
