package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName":"Joe","LastName":"Thunder","Age":24}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
}
