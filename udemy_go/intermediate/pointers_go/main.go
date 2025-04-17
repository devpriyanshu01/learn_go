package main

import "fmt"

func main() {
	var num int = 10

	var ptr *int = &num

	fmt.Println("Value of num: ", num)
	fmt.Println("Value of pointer variable: ", ptr)
	fmt.Println("Address of variable num: ", &num)
	fmt.Println("Value stored at the address pointer by ptr variable: ", *ptr)

	//modify the value of num by passing pointer as arg.
	modifyValue(ptr)
	fmt.Println("Modified Value, -", num)
}

func modifyValue(ptr *int) {
	*ptr = *ptr * 2
}
