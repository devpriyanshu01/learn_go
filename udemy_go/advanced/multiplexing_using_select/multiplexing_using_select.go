package main

import (
	"fmt"
	_"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:	//ok get false when channel is closed.
			if !ok {
				fmt.Println("Channel closed")
				// clean up activities
				return
			}
			fmt.Println("Received:", msg)
		}
	}
}

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 10
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Received:", msg)
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Timeout.")
// 	}
// }


// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func(){
// 		time.Sleep(time.Second)
// 		ch1 <- 10
// 	}()

// 	go func(){
// 		time.Sleep(time.Second)
// 		ch2 <- 20
// 	}()
	
// 	time.Sleep(time.Second*2)

// 	for range 2 {
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("received from ch1:", msg)
// 		case msg := <-ch2:
// 			fmt.Println("received from ch2:", msg)
// 		default:
// 			fmt.Println("No channels ready...")
// 		}
// 	}
// 	fmt.Println("end of program...")
	
// }
