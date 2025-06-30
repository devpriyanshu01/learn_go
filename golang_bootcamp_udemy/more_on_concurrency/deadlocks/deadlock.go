package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("Goroutine 1, locked mu1")
		time.Sleep(time.Second)

		mu2.Lock()
		fmt.Println("Goroutine 1, locked mu2")
		time.Sleep(time.Second)

		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("Goroutine 1 finished.")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("Goroutine 2, locked mu2")
		time.Sleep(time.Second)

		mu2.Lock()
		fmt.Println("Goroutine 2, locked mu1")
		time.Sleep(time.Second)

		mu2.Unlock()
		mu1.Unlock()
		fmt.Println("Goroutine 2 finished.")
	}()
	wg.Wait()
}
