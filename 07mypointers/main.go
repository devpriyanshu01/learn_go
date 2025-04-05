package main

import "fmt"

func main() {
	fmt.Println("Welcome to Learn Pointers...")

	var ptr *int 
	fmt.Println("pointer is : ", ptr)

	randomNo := 23
	var ptr2 *int = &randomNo
	fmt.Println("Value stored in ptr2 is: ", ptr2)
	fmt.Println("Value stored in ptr2 is: ", *ptr2)

	//Operations on pointers
	*ptr2 = *ptr2 * 2
	fmt.Println("New Value in the address pointed by ptr2 is: ", *ptr2)
}
