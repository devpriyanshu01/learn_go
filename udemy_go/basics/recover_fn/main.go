package main

import "fmt"

func main() {
	fmt.Println("Welcome to Learn Recover in Golang")

	process()
	fmt.Println("After Process is done!")
}

func process() {
	fmt.Println("process begins...")

	defer func(){
		r := recover()
		if r != nil {
			fmt.Println("Printing r :")
			fmt.Println(r)
			fmt.Println("Recovered from Error!")
		}
	}()

	panic("found error and closing the process")
}
