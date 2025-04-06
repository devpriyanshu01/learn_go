package main

import "fmt"

func main() {
	marks := 75

	if marks > 60 {
		fmt.Println("1st Division")
	} else if marks >= 50 && marks < 60 {
		fmt.Println("2nd Division")
	} else {
		fmt.Println("3rd Division or fail")
	}

	fmt.Println("----------------- Weird Syntax Below ------------------")

	//weird syntax - mostly used when handling web requests.
	if num := 7; num > 10 {
		fmt.Println("num is greater than 10")
	}else {
		fmt.Println("num is less than 10")
	}

}
