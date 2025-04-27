package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("-------- Random Numbers ---------")

	// fmt.Println("Random No. =>", rand.Intn(100))	//[0,100)

	//Generate Random No. between 6 to 10
	fmt.Println((rand.Intn(6) + 5))

	//Seeding
	val := rand.New(rand.NewSource(40))
	// val := rand.New(rand.NewSource(time.Now().Unix()))

	fmt.Println("Seeded Random No. =>", val.Intn(6)) //same value always

	fmt.Println("Random Float =>", rand.Float64())
	//above generate random float64 values.

	//Create Dice Game:
	fmt.Println("DICE GAME")

	for {
		//show menu
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Println("Enter your Choice.")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid Choice, enter 1 or 2")
			continue
		}
		if choice == 2 {
			fmt.Println("Exiting the Game!")
			break
		}
		dice1 := rand.Intn(6)+1
		dice2 := rand.Intn(6)+1

		fmt.Printf("You've got %d & %d \n", dice1, dice2)
		fmt.Printf("Total : %d\n", dice1+dice2)


	}

}
