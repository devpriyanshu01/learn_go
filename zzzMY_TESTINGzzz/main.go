package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("-----------  Welcome to my Testing Region  -------------")

	//In short below two lines is used to generate random nos.
	// source := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(source)

	randomNo := rand.Intn(100)
	fmt.Println("Generated random no - ", randomNo)

	// fmt.Println("Value of r -", *r)
}
