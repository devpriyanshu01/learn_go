package main

import (
	"fmt"
)

func main() {

	fmt.Println("................... INSIDE MAIN .......................")
	var mapSlice = []map[string]interface{} {
		{"id" : "6", "name" : "raman"},
		{"id" : 100, "name" : "raju"},
	}
	mapSlice[0]["email"] = "raman@gmail.com"
	fmt.Println("goSlice[0]", mapSlice[0])

}
