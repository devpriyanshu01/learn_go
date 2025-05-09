package main

import (
	"fmt"
	"time"
)
// Goroutines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished/ready to return any value
// Goroutines do not stop the program flow and are non blocking

func main() {
	fmt.Println("statement 1")
	go sayHello()
	go printLetters()
	go printNumbers()
	fmt.Println("statement 2")
	//sleep
	time.Sleep(1 * time.Second)
	fmt.Println("statement 3")
}

func sayHello() {
	fmt.Println("good afternoon.")
	time.Sleep(2 * time.Second)
}
func printNumbers() {
	for i := 1; i < 5 ; i++ {
		fmt.Println("i =", i)
		time.Sleep(100 * time.Millisecond)
	}
}
func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("letter:", string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}
