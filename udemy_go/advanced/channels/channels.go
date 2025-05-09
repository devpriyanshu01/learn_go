package main

import (
	"fmt"
	"time"
)

func main() {
	//channel creation
	greeting := make(chan string)

	greetString := "Hare Krsna"

	go func() {
		greeting <- greetString	//channels are blocking.
		greeting <- "Hare Krsna"

		for _, alpha := range "abcde" {
			greeting <- "alphabet - " + string(alpha)
		}
	}()

	go func ()  {
		receiver := <- greeting	//receivers are block till they receive any value. 
		fmt.Println("received:", receiver)	//after receiving the value they're non blocking.
		receiver = <- greeting
		fmt.Println("received:", receiver)

		for _ = range 5 {	//we need to have equal no. of receiver.
			receiver = <- greeting
			fmt.Println("received:", receiver)
		}
	}()
	// receiver := <-greeting
	// fmt.Println("received:", receiver)
	fmt.Println("waiting...\n...\n...\n...")
	time.Sleep(time.Second * 2)
	fmt.Println("finished.")
}
