package main

import (
	"fmt"
)

func main() {
	fmt.Println("main fn")

	// orderby := "first_name:desc"
	// res := strings.Split(orderby, ":")
	// for i, value := range res {
	// 	fmt.Println(i, " => ", value)
	// }

	var test interface{}
	test = "alsdf"
	var i string
	i = test.(string)
	fmt.Println("i", i)
}
