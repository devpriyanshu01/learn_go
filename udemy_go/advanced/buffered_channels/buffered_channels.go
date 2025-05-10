package main

import (
	"fmt"
	"time"
)

func main() {
	//====== DEMONSTRATION OF BLOCK ON SEND WHEN BUFFER IS FULL =======.
	ch := make(chan int, 2)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
		ch <- 2
	}()
	fmt.Println("blocking because channel is empty...")
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	fmt.Println("end of execution...")
}

// ====== DEMONSTRATION OF BLOCK ON SEND WHEN BUFFER IS FULL =======
// func main() {
// 	ch := make(chan int, 2) //a channel is created.

// 	// ch <- 1
// 	// ch <- 2

// 	//no error as this channel can store upto 2 int values.
// 	//If you receive the value, buffer gets empty.
// 	// fmt.Println("Received:", <-ch)
// 	// fmt.Println("Received:", <-ch)

// 	//Blocking mechanism of buffered channels.
// 	ch <- 1
// 	ch <- 2
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println("Received:", <-ch)
// 	}()
// 	fmt.Println("Blocking begins")
// 	ch <- 3	//blocks the execution as buffer is full.
// 	fmt.Println("Received:",<-ch)
// 	fmt.Println("Received:", <-ch)

// 	fmt.Println("end of program...")
// }
