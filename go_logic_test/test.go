package main

import "fmt"

func main() {

	a := '}'

	getAsciiValue(a)
}

func getAsciiValue(char rune) {
	fmt.Printf("Ascii of %v is -- %v\n", string(char), char)
}
