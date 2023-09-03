package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Client struct {
	Customer  Customer
	Hobbies   []string
	Addresses []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJsonArray(t *testing.T) {
	client := Client{
		Customer: Customer{
			FirstName: "Johnny",
			LastName:  "Silverhand",
			Age:       30,
		},
		Hobbies: []string{"Coding", "Cooking"},
		Addresses: []Address{
			{
				Street:     "Street 1",
				Country:    "Indonesia",
				PostalCode: "8017",
			},
			{
				Street:     "Street 2",
				Country:    "Indonesia",
				PostalCode: "8017",
			},
		},
	}

	bytes, _ := json.Marshal(client)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"Customer":{"FirstName":"Johnny","LastName":"Silverhand","Age":30},"Hobbies":["Coding","Cooking"],"Addresses":[{"Street":"Street 1","Country":"Indonesia","PostalCode":"8017"},{"Street":"Street 2","Country":"Indonesia","PostalCode":"8017"}]}`
	jsonBytes := []byte(jsonString)

	client := &Client{}
	err := json.Unmarshal(jsonBytes, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(client)
	fmt.Println(client.Customer.FirstName)
	fmt.Println(client.Customer.LastName)
	fmt.Println(client.Customer.Age)
	fmt.Println(client.Hobbies)
	fmt.Println(client.Addresses)

}
