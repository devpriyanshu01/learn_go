package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("------- Command Line Arguments --------")

	fmt.Println("Command:", os.Args[0])

	// fmt.Println("Argument1", os.Args[1])

	// args := os.Args
	// for _, value := range args {
	// 	fmt.Println("value:", value)
	// }

	//define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 20, "age of user")
	flag.BoolVar(&male, "male", false, "gender of user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("male:", male)
}
