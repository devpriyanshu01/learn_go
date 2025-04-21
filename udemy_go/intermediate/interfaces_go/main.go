package main

import (
	"fmt"
	"math"
	"reflect"
)

// define an interface
type geometry interface {
	area() float64
	perimeter() float64
}

// define a rectangle struct
type Rectangle struct {
	length, width float64
}

// define a circle struct
type Circle struct {
	radius float64
}

// area method for rectangle struct
func (r Rectangle) area() float64 {
	return r.length * r.width
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

// area method for circle struct
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

//Advantages of using Interface.
//if we've not used Geometry interface then we have to call
//rect.area() rect.perimeter() c.area() c.perimeter().
//But now we're just passing rect & c to measure fn.

func main() {
	rect := Rectangle{length: 4, width: 3}
	c := Circle{radius: 5}

	measure(rect)
	measure(c)

	//any type & any no. of type
	Printer("raman", 23, 345.23423, false)

	//call typeChecker
	typeChecker(23.234242423324)
}

// type checker function
func typeChecker(variable interface{}) {
	fmt.Println("Type is: ", reflect.ValueOf(variable).Kind())
}

// accepting any no. of values and any type of values
func Printer(i ...interface{}) {
	for _, value := range i {
		fmt.Println(value)
	}
}
