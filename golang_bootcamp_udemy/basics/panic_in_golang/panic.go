package main

import "fmt"

func main() {
	fmt.Println("Learning Panic in Golang")

	//panic(interface{})

	checkInput(-1)
}

func checkInput(input int) {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	if input < 0 {
		panic("Can't accept a non-positive integer")
	}
	fmt.Println("Input is --->> ", input)

}
