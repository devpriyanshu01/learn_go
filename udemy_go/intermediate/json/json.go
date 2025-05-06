package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	Age       int    `json:"age"`
	EmailAddress string `json:"email,omitentry"`
	Address	Address	`json:"address,omitentry"`
}
type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}

func main() {
	person1 := Person{
		FirstName: "Raman Kumar",
	}
	//struct to json conversion
	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error Marshelling:", err)
	}
	fmt.Println("jsondata:", string(jsonData))

	//nested structs
	person2 := Person{
		FirstName: "John Doe",
		Age : 20,
		EmailAddress: "johndoe@gmail.com",
		Address: Address{
			City : "New York",
			State :"NY",
		},
	}
	//print person2 object
	fmt.Println("person2:",person2)
	//struct object to json.
	person2AsJson, err2 := json.Marshal(person2)
	if err2 != nil {
		fmt.Println("Error parsing to json:", err2)
		return
	}
	fmt.Println("parsed to json:", string(person2AsJson))
	fmt.Println("--------------------------------------------------------------------------")
	//create a json string
	jsonString := `{"name":"John Doe","emp_id":"2000","age":"20","address":{"city":"california","state":"CA"}}`
	//json string to struct object conversion
	var jenny Employee
	err3 := json.Unmarshal([]byte(jsonString),&jenny)
	if err3 != nil {
		fmt.Println("Error unmarshelling:", err3)
		return
	}
	fmt.Println("Unmarshalled Object:", jenny)
	fmt.Println("Jenny age:", jenny.Age)
	fmt.Println("Jenny city:", jenny.Address.City)

	fmt.Println("--------------------------------------------------------------------------------")
	//Marshelling slice of Address to Json 
	listOfCityState := []Address{
		{City: "New Delhi", State: "Delhi"},
		{City: "Mumbai", State: "Maharashtra"},
		{City: "Patna", State: "Bihar"},
		{City: "Ranchi", State: "Jharkhand"},
		{City: "Kolkata", State: "West Bengal"},
	}
	//marshelling
	jsonListOfCityState, err := json.Marshal(listOfCityState)
	if err != nil {
		fmt.Println("error marshelling:", err)
		return
	}
	fmt.Println(string(jsonListOfCityState))

	fmt.Println("---------------------------------------------------------------------------")
	//handling json's whose structure are unknown.
	randomJson := `{"address":{"city":"dhanbad","state":"jharkhand"},"age":"20","name":"Govind","emp_id":"J300"}`
	//unmarshalling 
	var randomData map[string]interface{}
	err = json.Unmarshal([]byte(randomJson),&randomData)
	if err != nil {
		log.Fatalln("error unmarshalling:", err)
		return
	}
	//print the unmarshalled data
	fmt.Println("unmarshalled data:", randomData)
	fmt.Println("randomData name:", randomData["name"])
	fmt.Println("randomData emp_id:", randomData["emp_id"])
}
type Employee struct {
	Name string `json:"name"`
	Empid string `json:"emp_id"`
	Age string  `json:"age"`
	Address Address `json:"address"`
}
