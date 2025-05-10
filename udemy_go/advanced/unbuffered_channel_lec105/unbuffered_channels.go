package main

import (
	"fmt"
	"time"
)

func main() {
	//Unbuffered channels always needs an immediate receiver.
	//Channels and goroutine go hand and hand.
	//Receiver awaits for all the goroutines to finish.
	//Unbuffered Channels block on receive if there is no corresponding send is ready.
	ch := make(chan int)

	go func() {
		time.Sleep(2*time.Second)
		ch <- 1
	}()

	go func(){
		ch <- 2
		time.Sleep(4*time.Second)
	}()

	receiver := <- ch
	fmt.Println("received:", receiver)

	receiver = <- ch
	fmt.Println("received:", receiver)

	//Unbuffered Channels block on send if there is no corresponding receiver ready.
	//
	go func(){
		time.Sleep(5 * time.Second)
		receiver1 := <- ch
		fmt.Println("5 second goroutine finished.", receiver1)
	}()
	
	ch <- 3
}
