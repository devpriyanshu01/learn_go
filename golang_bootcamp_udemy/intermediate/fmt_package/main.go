package main

import (
	"fmt"
)

func main() {
	//PRINTING ACTIONS IN GO
	fmt.Println("PRINTING ACTIONS IN GO")

	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println(12, 456)

	fmt.Print("Hello") //no newline
	fmt.Print("World")
	fmt.Print(12, 456)

	fmt.Print("\n")

	name := "John"
	age := 252

	fmt.Printf("Name - %s, Age: %d\n", name, age)
	fmt.Printf("Age in Binary - %b, Age in Hexa-decimal : %X\n", age, age)

	//Foramtting Functions
	fmt.Println(".............. Formatting Functions ..............")

	s := fmt.Sprint("Hello", "World!", 1234, 1234325)
	//formats but no space in between & no new line after the end.
	fmt.Print(s)
	fmt.Print("\n")

	s2 := fmt.Sprintln("Hello", "World!", 1234, 3453)
	//formats the values, gives spaces between each world & adds newline at end.
	fmt.Print(s2)

	sf := fmt.Sprintf("Name : %s, Age : %d", "Raman", 20)
	fmt.Print(sf, "\n")

	fmt.Println("---------- Scanning Functions ------------")

	var userName string
	var userAge int

	fmt.Println("Enter name & age")
	//accept input/scan
	// fmt.Scan(&userName, &userAge)
	// fmt.Scanln(&userName, &userAge)
	// fmt.Scanf("%s %d", &userName, &userAge)

	//print the result
	fmt.Println("Name :- ", userName, "\nAge :- ", userAge)

	//Error Formatting Functions
	err := checkAge(122)
	if err != nil {
		fmt.Println("Error - ", err)
	} else {
		fmt.Println("Licensed To Drive")
	}

}
func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is too young to drive", age)
	}
	return nil
}
