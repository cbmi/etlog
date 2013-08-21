package stores

import (
	"encoding/json"
	"github.com/cojac/assert"
	"testing"
)

// Note: Although this tests JSON marshalling, the BSON marshaller is used internally
func TestStoreJSONMarshal(t *testing.T) {
	s := &Store{Uri: "http://127.0.0.1"}
	b, _ := json.Marshal(s)
	assert.Equal(t, `{"uri":"http://127.0.0.1"}`, string(b))

	s = &Store{Uri: "http://127.0.0.1", Value: []string{"one", "two"}}
	b, _ = json.Marshal(s)
	assert.Equal(t, `{"uri":"http://127.0.0.1","value":["one","two"]}`, string(b))

	s = &Store{Uri: "http://127.0.0.1", Value: map[string]int{"one": 1, "two": 2}}
	b, _ = json.Marshal(s)
	assert.Equal(t, `{"uri":"http://127.0.0.1","value":{"one":1,"two":2}}`, string(b))

	s = &Store{Uri: "http://127.0.0.1", Extra: map[string]interface{}{"one": 1, "two": 2}}
	b, _ = json.Marshal(s)
	assert.Equal(t, `{"one":1,"two":2,"uri":"http://127.0.0.1"}`, string(b))
}
