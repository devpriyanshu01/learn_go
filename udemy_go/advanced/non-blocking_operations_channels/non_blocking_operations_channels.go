package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---------- Non - Blocking Operations on Channels -----------")

	//NON-BLOCKING RECEIVE OPERATION
	// ch := make(chan int)
	// select {
	// case <- ch:
	// 	fmt.Println("received:", <- ch)
	// default :
	// 	fmt.Println("channel data not available")
	// }

	//NON-BLOCKING SEND OPERATION
	// ch := make(chan int)
	// select {
	// case ch <- 10:
	// 	fmt.Println("data sent to channel")
	// default :
	// 	fmt.Println("couldn't sent to channel")
	// }

	//NON-BLOCKING OPERATIONS IN REAL TIME SYSTEMS
	data := make(chan int)
	quit := make(chan bool)

	go func(){
		for{
			select {
			case value := <- data:
				fmt.Println("data received:", value)
			case <- quit :
				fmt.Println("stopping...")
				return
			default:
				fmt.Println("waiting for data...")
				time.Sleep(500*time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- (10 * i)
	}
	quit <- true
}
