package main

import (
	"fmt"
	"strconv"
	"time"
)

// func main(){
// 	ch := make(chan int)

// 	go func(){
// 		time.Sleep(2*time.Second)
// 		ch <- 9	//blocking until the value is received in channel.
// 		fmt.Println("Value sent to channel...")
// 	}()
// 	value := <- ch	//blocking until value is sent from channel.
// 	fmt.Println("value received from channel:", value)
// 	fmt.Println("end of execution...")
// }

// func main() {
// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2*time.Second)
// 		done <- struct{}{}
// 	}()

// 	<-done
// 	fmt.Println("End Of Execution...")
// }

// func main() {
// 	//Synchronizing multiple goroutines and ensuring that all are complete.
// 	numGoroutines := 3
// 	ch := make(chan int, 3)

// 	for i := range numGoroutines {
// 		go func(id int){
// 			time.Sleep(2 * time.Second)
// 			ch <- id
// 		}(i)
// 	}
// 	fmt.Println("awaiting goroutines...")
// 	for range numGoroutines {
// 		fmt.Println("Received:", <- ch)
// 	}
// 	fmt.Println("All goroutines finished. End of execution...")
// }

//Example 4: Synchronizing data exhange.
func main() {
	data := make(chan string)

	go func(){
		for i := range 5 {
			data <- "hello " + strconv.Itoa(i)
			time.Sleep(time.Second)
		}
		close(data)	//closes the channel when empty.
	}()

	for value := range data {
		fmt.Println("Received:", value, time.Now())
	}// Loops over only on active channel, creates receiver each time and stops creating receiver (looping) once the channel is closed
}