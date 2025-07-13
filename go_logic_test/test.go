package main

import "fmt"

func main() {

	a := '}'

	getAsciiValue(a)
   --TODO
}

func getAsciiValue(char rune) {
	fmt.Printf("Ascii of %v is -- %v\n", string(char), char)
}
