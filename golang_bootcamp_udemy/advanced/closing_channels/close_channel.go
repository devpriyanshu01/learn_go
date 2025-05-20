package main

import "fmt"

func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}
}

//===== RANGE OVER A CLOSED CHANNEL
// func main(){
// 	ch := make(chan int)
// 	go func(){
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()
// 	for val := range ch {
// 		fmt.Println("received:", val)
// 	}
// }

//======== RECEIVING FROM A CLOSED CHANNEL
// func main(){
// 	ch := make(chan int)
// 	// ch <- 10
// 	close(ch)
// 	val, ok := <- ch 
// 	if !ok {
// 		fmt.Println("channel is closed.")
// 			// break
// 	}else{
// 		fmt.Println("value received:", val)
// 	}
// }

//==== SIMPLE CLOSING CHANNEL EXAMPLE
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()
// 	for val := range ch {
// 		fmt.Println("received:", val)
// 	}
// }
