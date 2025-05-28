package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{Name: "Raman Kr", Email: "raman@gmail.com"}
	fmt.Println("user", user)
	//marshal to json
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println("raw data", jsonData)
	fmt.Println(string(jsonData))

	//Unmarshal Json
	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatalln("err un-marshalling", err)
	}
	fmt.Println("user1:", user1)

	fmt.Println("--------------------------------------------------------------------")
	jsonObject := `{"Name" : "Krsna", "Email" : "krsna1@gmail.com"}`
	reader := strings.NewReader(jsonObject)
	decoder := json.NewDecoder(reader)

	var user2 User
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Encoded json string", buf.String())
}
