package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

//struct embedding
type Shape struct {
	Rectangle
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

//Method with pointer receiver
func (r *Rectangle) Scale(scale float64) {
	r.length *= scale
	r.width *= scale
}

type MyInt int

//Methods can exist for normal data types as well, it's
//not necessary to have ony for composite data type i.e. structs
func (num MyInt) isPositive() bool{
	if num > 0 {
		return true
	}else {
		return false
	}
}

//Anonymous intance
func (MyInt) welcomeMsg() string{
	return "Good morning sir"
}

func main() {
	rect := Rectangle{length: 10, width: 9}		//initialize a rectangle object.
	area := rect.Area()		//calculate area on rect object
	fmt.Println("Area :", area)
	rect.Scale(2)	//scale len & width by 2 by calling Scale - pointer receiver
	areaNew := rect.Area()	//get new area by doing object.Area()
	fmt.Println("New Area :", areaNew)

	num := MyInt(5)
	num2 := MyInt(-2)
	fmt.Println("num is: ", num.isPositive())
	fmt.Println("num2 is:", num2.isPositive())
	fmt.Println("num2 is:", num2.isPositive())

	fmt.Println("num2 msg:", num2.welcomeMsg())

	shape1 := Shape{Rectangle: Rectangle{length: 4, width: 3}}
	//struct embedding promoted method embedding to outer struct
	fmt.Println("Embedded Struct Area:", shape1.Area())
	//above is more convnient than below.
	fmt.Println("Embedded Struct Area2:", shape1.Rectangle.Area())

	
}
