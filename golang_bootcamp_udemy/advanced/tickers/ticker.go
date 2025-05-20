package main

import (
	"fmt"
	"time"
)
//using timers to stop a ticker.
func main(){
	ticker3 := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	fmt.Println("hi there")
	defer ticker3.Stop()

	for {
		select {
		case tick := <-ticker3.C:
			fmt.Println("tick at:", tick)
		case <-stop:
			fmt.Println("ticker3 stopped.")
			return
		}
	}
}

// func periodicTasks() {
// 	fmt.Println("periodic task at:", time.Now())
// }
// ---------- Periodic Task Example ------------
// func main() {
// 	ticker2 := time.NewTicker(time.Second)
// 	defer ticker2.Stop()
// 	for {
// 		select {
// 		case <-ticker2.C:
// 			periodicTasks()
// 		}
// 	}
// }

// func main() {
// 	ticker1 := time.NewTicker(time.Second)
// 	defer ticker1.Stop()

// 	i := 2
// 	for range 10 {
// 		i = i * 2
// 		fmt.Println(i)
// 	}
// }
