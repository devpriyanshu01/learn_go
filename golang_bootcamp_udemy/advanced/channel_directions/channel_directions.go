package main

import "fmt"

//send only channel
// func main() {
// 	ch := make(chan  int)

// 	go func(ch chan<- int) {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}(ch)

// 	for range 5 {
// 		fmt.Println("Received:", <-ch)
// 	}
// }

//Receive only channel
func main() {
	ch := make(chan  int)

	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	receiveData(ch)
}
func receiveData(ch <- chan int){
	for range 5 {
		fmt.Println("Received:", <- ch)
	}
}
