package encoding

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
)

// This encodes it as bson as an intermediate step to take advantage of the
// `inline` tag to support arbirtary fields. This can be used as an alternate
// to `json.Marshal`.
// See: http://godoc.org/labix.org/v2/mgo/bson#Marshal
func MarshalJSON(v interface{}) ([]byte, error) {
	var j interface{}
	b, _ := bson.Marshal(v)
	bson.Unmarshal(b, &j)
	return json.Marshal(&j)
}

// This decodes it as json into an intermediate map, encodes it into BSON,
// the decodes it back into the store. This take advantages of the `inline`
// tag to support arbirtary fields. This can be used as an alternate to
// `json.Unmarshal`
// See: http://godoc.org/labix.org/v2/mgo/bson#Marshal
func UnmarshalJSON(b []byte, v interface{}) error {
	var j map[string]interface{}
	json.Unmarshal(b, &j)
	b, _ = bson.Marshal(&j)
	return bson.Unmarshal(b, v)
}
