package encoding

import (
	"github.com/bruth/assert"
	"testing"
)

type Message struct {
	Text  string                 `json:"text" bson:"text"`
	Extra map[string]interface{} `json:"-" bson:",inline"`
}

func TestUnmarshalJSON(t *testing.T) {
	m := Message{}

	b := []byte(`{
        "text": "Hello World",
        "time": "2013-08-21T22:10:34"
    }`)

	UnmarshalJSON(b, &m)

	assert.Equal(t, m.Text, "Hello World")
	assert.Equal(t, m.Extra, map[string]interface{}{"time": "2013-08-21T22:10:34"})
}

func TestMarshalJSON(t *testing.T) {
	m := Message{
		Text: "Hello World",
		Extra: map[string]interface{}{
			"time": "2013-08-21T22:10:34",
		},
	}

	b, _ := MarshalJSON(&m)
	assert.Equal(t, string(b), `{"text":"Hello World","time":"2013-08-21T22:10:34"}`)
}
