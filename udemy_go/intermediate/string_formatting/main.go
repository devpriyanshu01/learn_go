package main

import "fmt"

func main() {
	fmt.Println("------------ Let's learn about String Formatting in Golang ---------")

	num := 424	//00424
	fmt.Printf("%05d\n", num)

	str := "hello"
	fmt.Printf("|%10s|\n", str)	//leading spaces - |     hello|
	fmt.Printf("|%-10s|\n", str) //trailing spaces - |hello     |

	//string interpolation
	message1 := "Hello \nWorld!"
	message2 := `Hello \nWorld!`

	fmt.Println(message1) //will provide a newline because of \n
	fmt.Println(message2) //will keep is a line \n is considered a char not
						  //a format specifier.
}
