package main

import "fmt"

func main() {
	ans := findFactorial(100)
	fmt.Println("factorial is : ", ans)
}

func findFactorial(num int) int {
	//base case
	if num == 1 {
		return 1
	} else if num == 0 {
		return 1
	}

	return num * findFactorial(num-1)
}
