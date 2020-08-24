package main

import (
	"encoding/json"
	"fmt"
)

// The identifiers have to be capitalized so that they can be exported to the standard library functions
// JSON tags help the decoder to map the JSON data to a type in our Go code
type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []Person{p1, p2}
	fmt.Println(people)

	xb, err := json.Marshal(people)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(string(xb))

	var data []Person
	err = json.Unmarshal(xb, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(data)
}
