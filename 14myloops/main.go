package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang")

	var days = []string{"Monday", "Sunday", "Tuesday", "Thursday", "Saturday"}

	fmt.Println("Days are -------------->>", days)

	// basic for loop 
	for i := 0; i < len(days) ; i++ {
		// fmt.Printf("%v ,", days[d])
		fmt.Println(days[i])
	}

	fmt.Println("------------------- LINE BREAK -------------------")

	//First Varient of foreach loop
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("---------------- LINE BREAK ------------------")
	//for each loop
	for index, value := range days {
		fmt.Printf("Day at index %v is %v \n", index, value)
	}

	fmt.Println("=================== While Loop in Go ====================")

	initValue := 0;

	for initValue < 10 {

		// if initValue == 5{
		// 	break
		// }

		if initValue == 8 {   //This will go to lebel Vaikuntha & will remaining
			goto Vaikuntha    //iteration.
		}

		if initValue == 2 {
			initValue++
			continue
		}

		fmt.Println(initValue)
		initValue++
	}

	Vaikuntha :
		fmt.Println("==============  Welcome to Vaikuntha, Now you are free from material bondage  =================")

}
