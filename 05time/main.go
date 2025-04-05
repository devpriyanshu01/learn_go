package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to study time package in Go!")

	currentTime := time.Now()

	//let's format this currentTime var
	formatedTime := currentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println(formatedTime)
	//or you may directly print after formatting it.
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	//Let's try to do the reverse of above i.e. create a date
	birthDay := time.Date(2008, time.December, 15, 15, 32, 0, 0, time.Local)
	fmt.Println("Your birthday is :", birthDay)

}
