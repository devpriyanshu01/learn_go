package main

import "fmt"

// Keypoints for Go Functions
// 	1. multiple return type should be enclosed in () & separated by comma(,) .
// 	2. Arguments passed to fn's are passed as copy, not as reference.

func main() {
	fmt.Println("Inside Main fn.")
	
	sum := add(2,3)
	fmt.Println("sum -->", sum)

	//Anonymous Fn.
	func(){
		fmt.Println("Anonymous Fn.")
	}()

	greet := func(){
		fmt.Println("Good Morning Go")
	}
	//greet is now fn. So, it can be called.
	greet()

	operation := add	//func add is assigned to a variable operation.
	result := operation(5,10)	//Since operation is now a Fn. So it can be called.
	fmt.Println("Result is --> ", result)

	//we're passing a function as an argument
	result2 := applyOperation(18, 11, add)
	fmt.Println("result2 --> ", result2)

	//Fucntion returning a function
	multiplyBy2 := createMultiplier(33)
	fmt.Println("2 * 33 = ", multiplyBy2(2))
	//			OR
	ans := multiplyBy2(2)	
	fmt.Println("ans -- ", ans)

}

//Function that takes a function as an argument.
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

//Function that returns a function.
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func add(a, b int) int {
	return a+b
}
