package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func convertToJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	// convert tipe data di Golang menjadi JSON.
	convertToJson("Hello")
	convertToJson(1)
	convertToJson(true)
	convertToJson([]string{"Hello", "World"})
}
