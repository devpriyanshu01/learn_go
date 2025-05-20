package main

import "fmt"

func main() {
	//Arrays in Go
	var fruitList [4]string

	fmt.Println("Length of friuitList is: ", len(fruitList))

	fruitList[0] = "Grapes"
	fruitList[1] = "Banana"
	fruitList[3] = "Guava"

	fmt.Println("Fruit lists are: ", fruitList)
	fmt.Println("Fruit at index 2 is: ", fruitList[2])

	//Veggies List
	var veggieList = [3]string {"Bhindi", "Kathal", "Began"}
	fmt.Println("VeggieList is : ", veggieList)

}
