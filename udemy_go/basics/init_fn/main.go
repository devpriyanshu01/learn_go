package main

import "fmt"

func init() {
	fmt.Println("Initializing package1")
}

func init() {
	fmt.Println("Initializing package2")
}
func init() {
	fmt.Println("Initializing package3")
}

func main() {
	fmt.Println("Starting of Main Fn.")
}
