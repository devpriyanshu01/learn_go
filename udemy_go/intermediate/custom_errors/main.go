package main

import "fmt"

func main() {
	fmt.Println("------------- Handling Custom Erros in Golang -------------")

	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation Completed Successfully!")
}

type customError struct {
	code    int
	message string
}

//Error returns the error message, implementing Error() method of error interface.
func (e *customError) Error() string {
	return fmt.Sprintf("Error: %d %s", e.code, e.message)
}

//Function that returns a Custom Error
func doSomething() error {
	return &customError{
		code : 501,
		message: "Something went Wrong!",
	}
}


