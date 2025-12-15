package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string  `json:"name"`
	Age  int     `json:"age,omitempty"`
	City string  `json:"city"`
	Addr Address `json:"address,omitempty"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip,omitempty"`
}

func main() {
	person := Person{
		Name: "Alice",
		City: "New York",
		Addr: Address{
			Street: "123 Main St",
			City:   "New York",
		},
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	jsonData1 := `{"name":"Bob","age":30,"city":"Los Angeles","address":{"street":"456 Elm St","city":"Los Angeles","zip":"90001"}}`
	var person1 Person
	err = json.Unmarshal([]byte(jsonData1), &person1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", person1)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData1), &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}
