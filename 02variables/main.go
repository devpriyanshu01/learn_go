package main

import "fmt"

const Token string = "ajsoeurasldjf"  //public as started with capital letter.

func main(){
	var username string = "raman"
	fmt.Println(username)
	fmt.Printf("Type of username is: %T \n", username)

	var isLazy bool = true
	fmt.Println(isLazy)
	fmt.Printf("Type of username is: %T \n", isLazy)

	var marks int = 32
	fmt.Println(marks)
	fmt.Printf("Type of username is: %T \n", marks)

	var decimal float64 = 22.23423424937429742
	fmt.Println(decimal)
	fmt.Printf("Type of username is: %T \n", decimal)

	//default values
	var defaultValue int
	fmt.Println(defaultValue)
	fmt.Printf("Type of username is: %T \n", defaultValue)

	var defaultValue_float float32
	fmt.Println(defaultValue_float)
	fmt.Printf("Type of username is: %T \n", defaultValue_float)

	var defaultValue_bool bool
	fmt.Println(defaultValue_bool)
	fmt.Printf("Type of username is: %T \n", defaultValue_bool)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(Token)
	fmt.Printf("Type of Token is : %T \n", Token)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")

	fmt.Println("Implicit variable declarations")
	//implicit type
	var website = "priyanshudev.in"
	fmt.Println(website)

	var mobile_no = 92342932784
	fmt.Println(mobile_no)

	
	//no var style
	fmt.Println("No var keyword declarations")
	pincode := 829323
	fmt.Println(pincode)

	name := "Krsna"
	fmt.Println(name)

}