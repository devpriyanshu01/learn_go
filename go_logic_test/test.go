package main

import (
	"fmt"
	"reflect"
)

type student struct {
	id         int
	First_name string
	last_name  string
}

func main() {
	s1 := student{
		id:         100,
		First_name: "raman",
		last_name:  "kumar",
	}
	var value interface{} = "Roshan"
	sv := reflect.ValueOf(&s1).Elem()
	if sv.Field(1).CanSet() {
		fmt.Println("inside canset")
		sv.Field(1).SetString(value.(string))
	}

	fmt.Println("s1", s1)
}
