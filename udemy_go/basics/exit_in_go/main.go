package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Let's learn Exit in Golang")

	fmt.Println("Starting Main Fucntion")

	defer fmt.Println("Defer Statement") //This won't get executed.

	os.Exit(1)

	fmt.Println("After Exit Statement") //This also won't be executed.
}
