package main

import (
	"fmt"
	"time"
)



func main(){
	timer := time.NewTimer(2 * time.Second)
	go func(){
		<-timer.C
		fmt.Println("delayed operation executed")
	}()
	fmt.Println("waiting...")
	time.Sleep(time.Second*3)
	fmt.Println("end of program")
}

// func longRunningOperation(){
// 	for i := range 20{
// 		fmt.Println("i:", i)
// 		time.Sleep(time.Second)
// 	}

// }

// func main(){
// 	timeout := time.After(2 * time.Second)
// 	done := make(chan bool)
// 	go func(){
// 		longRunningOperation()
// 		done <- true
// 	}()
// 	select {
// 	case <-timeout:
// 		fmt.Println("timer completed")
// 	case <-done:
// 		fmt.Println("goroutine finished.")
// 	}
// }

// func main() {
// 	timer1 := time.NewTimer(2 * time.Second)
	
// 	fmt.Println("waiting for the timer...")
// 	stopped := timer1.Stop()
// 	if stopped {
// 		fmt.Println("timer stopped")
// 	}else{
// 		fmt.Println("timer running")
// 	}
// 	timer1.Reset(time.Second*3)
// 	<-timer1.C //blocking in nature
// 	fmt.Println("timer  expired")
// }
