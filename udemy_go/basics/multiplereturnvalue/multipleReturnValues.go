package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Multiple Return Values from a Fn. in Golang")

	q, r := divide(10,3)
	fmt.Printf("Quotient - %v, Remainder - %v\n", q,r)

	result, err := compare(3 , 3)
	if err != nil {
		fmt.Println("Error encountered : ", err)
	}else {
		fmt.Println("Result is : ", result)
	}

}
func divide(a, b int)(int,int){
	quotient := a/b
	remainder := a%b

	return quotient,remainder
}

func divide2(a, b int)(quotient int,remainder int){
	quotient = a/b
	remainder = a%b
	//Go compiler is smart enough to understand that quotient and remainder
	//needs to be returned.
	return
}

func compare(a, b int)(string, error) {
	if a > b {
		return "a is greater than b", nil
	}else if b > a {
		return "b is greater than a", nil
	}else {
		return "", errors.New("Can't decide which is greater")
	}
}
