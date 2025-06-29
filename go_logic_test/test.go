package main

import (
	"fmt"
	"reflect"
	"strings"
)

// type student struct {
// 	Name  string
// 	Class int
// }

type Teacher struct {
	Name string
	Subject string `json:"subject,omitempty"`
	Age 	int
}

func main() {

	teacher1 := Teacher {
		Name: "Raman",
		Subject: "Computer Science",
		Age : 21,
	}
	fmt.Println("teacher1", teacher1)

	teacherVal := reflect.ValueOf(&teacher1).Elem()
	teacherType := teacherVal.Type()

	fmt.Println("teacherType", teacherType)	//main.Teacher
	fmt.Println("teacherVal.Field(1) ------", teacherVal.Field(1).Type())
	// fmt.Println("teacherType.Field(i)", teacherType.Field(1))
	typeField := teacherType.Field(1)
	fmt.Println(typeField.Tag.Get("json"))
	str := typeField.Tag.Get("json")
	trimedString := strings.TrimSuffix(str, ",omitempty")
	fmt.Println("trimmed string", trimedString)

	// fmt.Println("main fn")

	// data := make(map[string]interface{})
	
	// data["id"] = 101
	// data["name"] = "raman"

	// id, ok := data["id"].(int)
	// if !ok {
	// 	fmt.Println("ERROR OCCURED")
	// }else {
	// 	fmt.Println("Value =", id)
	// 	fmt.Println("ok = ", ok)
	// }

	// student1 := student{
	// 	Name : "raman",
	// 	Class : 2,
	// }

	// v := reflect.ValueOf(&student1).Elem()
	// fmt.Println("v", v)
	// t := v.Type()
	// fmt.Println("type", t)
	// fmt.Println("type at 0", v.Field(0).Type())

	

	

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
