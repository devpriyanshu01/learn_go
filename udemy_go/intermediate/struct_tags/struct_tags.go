package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name,omitempty" db:"firstName" xml:"firstname"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"-"`	//will omit this field anyway.
}

func main() {
	// person1 := Person{FirstName: "Stephen", Age: 40}
	person2 := Person{FirstName: "Stephen", Age: 40}
	//marshell to json
	person1Json, err := json.Marshal(person2)
	if err != nil {
		log.Fatalln("error marshelling:", err)
		return
	}
	fmt.Println("person1 as json:", string(person1Json))
}
