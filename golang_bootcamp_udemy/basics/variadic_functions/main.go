package main

import "fmt"

func main() {
	//...Ellipsis
	//func functionName(param1 type1, param2 type2, param3 ... type3) return_type {
	//body of fn
	// }
	//Here param1 & param2 are mandatory and rest are optional.
	//variadic parameters must be at the end.
	//passing slice as argument in variadic fn, will destructure it and then assign it.

	fmt.Println("Sum = ", calculateSum(2, 23, 43, 23))

	numbers := []int{1,3,5,7,9}

	text, sum := calculateSum2("Slice as arg. in Variadic Fn. ", numbers...)
	fmt.Println(text, sum)

}

func calculateSum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func calculateSum2(text string, nums ...int) (string, int) {
	total := 0
	for _, value := range nums {
		total += value
	}
	return text, total
}
