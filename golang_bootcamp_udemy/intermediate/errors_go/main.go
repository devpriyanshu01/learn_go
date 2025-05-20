package main

import (
	"errors"
	"fmt"
)

func sqrt(x float32)(float32, error){
	if x < 0 {
		return 0, errors.New("Math Error: square root of negative no.")
	}
	//compute the square root
	return 1, nil
}

//func to check if data is empty?
func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error : Empty Data")
	}
	return nil
}

func main() {
	fmt.Println("---------------- Learning Errors in Golang ---------------")

	result1, err1 := sqrt(32)
	if err1 != nil {
		fmt.Println(err1)
	}else {
		fmt.Println("Result1 is =>", result1)
	}

	result2, err2 := sqrt(-32)
	if err2 != nil {
		fmt.Println(err2)
	}else {
		fmt.Println("Result1 is =>", result2)
	}

	//check data if it's empty
	// data := []byte {}
	// err := process(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if err := process(data) ; err != nil {}
	// fmt.Println("Data processed successfully")

	fmt.Println("--------- Custom Error Message -----------")
	err := eprocess()
	fmt.Println("Error: ", err)

	fmt.Println("----------- Read Data --------------")
	res := readData()
	if res!= nil {
		fmt.Println(res)
		return
	}
	fmt.Println("Read Data Successful!!!")
}

//Let's now create a custom error message.
type myError struct {
	message string
}
//method for this struct
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

//Above is implementing the built-in error interface
//This built-in error interface looks like as below.
// type error interface{
// 	Error() string
// }

func eprocess() error {
	return &myError{"Custom Error Message"}
}

func readConfig() error{
	return errors.New("Config Error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}