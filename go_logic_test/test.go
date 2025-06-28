package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Class int
}

func main() {
	fmt.Println("main fn")

	student1 := student{
		Name : "raman",
		Class : 2,
	}

	v := reflect.ValueOf(&student1).Elem()
	fmt.Println("v", v)
	t := v.Type()
	fmt.Println("type", t)
	fmt.Println("type at 0", v.Field(0).Type())

	// data := make(map[string]string)
	// data["first_name"] = "raman"
	// fmt.Println(len(data))
	// fmt.Println()

	// student1 := student{
	// 	Name : "raman",
	// 	Class : 2,
	// }

	// v1 := reflect.ValueOf(student1)
	// fmt.Println("v1", v1)
	// v2 := reflect.ValueOf(&student1).Elem()
	// fmt.Println("v2", v2)

	// if v1 == v2 {
	// 	fmt.Println("student1 == v2")
	// }else {
	// 	fmt.Println("student1 & v2 are different struct objects")
	// }

	// orderby := "first_name:desc"
	// res := strings.Split(orderby, ":")
	// for i, value := range res {
	// 	fmt.Println(i, " => ", value)
	// }

	// var test interface{}
	// test = "alsdf"
	// var i string
	// i = test.(string)
	// fmt.Println("i", i)
}
