package main

import (
	"fmt"
	"reflect"
)

type Greeter struct{}

func (g Greeter) Greet(fname string) string {
	return "Hello " + fname
}

// WORKING WITH METHODS
func main() {
	g := Greeter{}

	t := reflect.TypeOf(g)
	fmt.Println("type of g", t)

	//find the no. of method for g
	for i := range t.NumMethod() {
		method := t.Method(i)
		fmt.Printf("Method %d --> %v \n", i, method.Name)
	}

	//use Greet method of g struct
	v := reflect.ValueOf(g)
	method := t.Method(0)
	
	m := v.MethodByName(method.Name)
	result := m.Call([]reflect.Value{reflect.ValueOf("Raman")})

	fmt.Println("Result:", result[0].String())


}

// ======== WORKING WITH STRUCTS and FIELDS
// type person struct {
// 	Name string
// 	age  int
// }

// func main() {
// 	p := person{Name: "Alice", age: 30}
// 	v := reflect.ValueOf(p)

// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}

// 	v1 := reflect.ValueOf(&p).Elem()

// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet() {
// 		nameField.SetString("Jane")
// 	} else {
// 		fmt.Println("Cannot set")
// 	}

// 	fmt.Println("Modified Person:", p)
// }

// func main() {

// 	x := 42
// 	v := reflect.ValueOf(x)
// 	t := v.Type()

// 	fmt.Println("Value:", v)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())
// 	fmt.Println("Is Int:", t.Kind() == reflect.Int)
// 	fmt.Println("Is String:", t.Kind() == reflect.String)
// 	fmt.Println("Is Zero:", v.IsZero())

// 	y := 10
// 	v1 := reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)
// 	fmt.Println("V2 Type:", v2.Type())

// 	fmt.Println("Original value:", v1.Int())

// 	v1.SetInt(18)
// 	fmt.Println("Modified value:", v1.Int())

// 	var itf interface{} = "Hello"
// 	v3 := reflect.ValueOf(itf)

// 	fmt.Println("V3 Type:", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println("String value:", v3.String())
// 	}
