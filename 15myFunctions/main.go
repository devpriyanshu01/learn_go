package main

import "fmt"

func main() {
	fmt.Println("============= Welcome to Fn's in Golang ==============")

	Greet()

	sum := mySum(23,342)
	fmt.Println("Sum is --->", sum)

	fmt.Println("---------- Working with ProAdder fn -------------")
	result := proAdder(23,23452,2342,2323)
	fmt.Println("Result of proAdder is ==> ", result)
	
}

func mySum(param1 int, param2 int) int {
	fmt.Println("---------- Inside mySum fn. -----------")
	return param1 + param2
}

func proAdder(params ...int) int {	//params is a Slice here, which can accept 
	ans := 0						//any number of Values.

	for _, value := range params {
		ans += value
	}

	return ans
}

func Greet() {
	fmt.Println("_________HK from Greet fn.____________")
}
