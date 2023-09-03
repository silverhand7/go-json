package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Jordan",
		LastName:  "Henderson",
		Age:       31,
	}

	encoder.Encode(customer)
}
