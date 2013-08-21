package services

import (
	"io"
	"io/ioutil"
	"log"
    "encoding/json"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleMessage(r io.Reader) {
    var t map[string]interface{}
	b, _ := ioutil.ReadAll(r)
    err := json.Unmarshal(b, &t)
	handleError(err)
	insertDoc(t)
}
