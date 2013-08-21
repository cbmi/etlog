package main

import (
    "etlog/encoding"
)

type Script struct {
    // URI which denotes where the data lives, where came from (streamed) or
    // where can be accessed in the future. The latter value is preferred since
    // it theoretically enables future access.
    Uri string `json:"uri,omitempty" bson:"uri,omitempty"`

    // The version of the script. This could be a timestamp, Git commit SHA,
    // tag version, etc. This useful when an issue is found with the script
    // and it can be assessed whether a transaction is affected.
    Version string `json:"version,omitempty" bson:"version,omitempty"`

    // The primary programming language the script is written in. If not
    // defined, this will be attempted to be inferred from filename specified
    // in uri.
    Language string `json:"language,omitempty" bson:"language,omitempty"`

    // The actual code as text or statement used during this transaction.
    // This is most useful for scripts that perform in-place operations that
    // never expose the data itself. For example, a SQL statement that selects
    // data from one table and inserts it into a new table.
    Code string `json:"code,omitempty" bson:"code,omitempty"`

	// Extra data from decoded JSON. This enables clients to store additional
	// metadata about the store.
    Extra map[string]interface{} `json:"-" bson",inline,omitempty"`
}

// Satisfies the `json.Marshaler` interface
func (s *Script) MarshalJSON() ([]byte, error) {
    return encoding.MarshalJSON(s)
}

// Satisfies the `json.Unmarshaler` interface
func (s *Script) UnmarshalJSON(b []byte) error {
    return encoding.UnmarshalJSON(b, s)
}
