package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("main fn")

	orderby := "first_name:desc"
	res := strings.Split(orderby, ":")
	for i, value := range res {
		fmt.Println(i, " => ", value)
	}
}
