package main

import "fmt"

func main() {
	fmt.Println("------------ WELCOME TO LEARN DEFER IN GO ------------")

	//Defer in english means to postpone for later.
	//All defer codes are executed in LIFO - Last In First Out order.
	
	defer fmt.Println("three")
	defer fmt.Println("two")

	fmt.Println("one")
	defer Iterator()

	//Output will be :- 
	//one, 15, 14, 13, 12, 11, 10, two , three
}

func Iterator() {
	for i := 10; i <= 15; i++ {
		defer fmt.Println(i)
	}
}
