package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Joe",
		LastName:  "Thunder",
		Age:       24,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
