package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Your are welcomed to learn Slices in Go")

	var fruitList = []string {"Grapes", "Mango", "Papaya"}

	//see the type of fruitList
	fmt.Printf("The type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Guava", "Gooseberry")
	//print new fruitList
	fmt.Println("Updated FruitList is: ", fruitList)
	fmt.Println("Some fruits from FruitList are: ", fruitList[:4])

	//High-Scores Examples
	highScores := make([]int, 4)

	highScores[0] = 232
	highScores[1] = 235
	highScores[2] = 236
	highScores[3] = 222

	// highScores[4] = 332 //This will throw error
	//But when append is used, it will not throw error because
	//memory allocation happens again.

	highScores = append(highScores, 978,977)

	fmt.Println("Print HighScores :", highScores)
	fmt.Printf("Type %T\n", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted highscores : ", highScores)
	fmt.Println("is highScore sorted --> ", sort.IntsAreSorted(highScores))

	fmt.Println("-----------------ANOTHER LECTURE-------------------")

	//How to remove a value from Slices based on an index.
	//initialize a slice
	var courses = []string {"C#", "Java", "JS", "NodeJS", "Docker", "Postgres"}

	fmt.Println("Printing Courses --> ", courses)

	//added few more values to courses
	//exclude NodeJS
	var index int = 3
	courses = append(courses[:index], courses[index+1:]...) //why ...? Learn in future
	fmt.Println("Printing updated Courses slice --> ", courses)
}
