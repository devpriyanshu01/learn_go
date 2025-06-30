package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Inside Main Fn.")
	// go printNumbers()
	// go printLetters()

	// time.Sleep(time.Second*3)
	numThreads := 5
	runtime.GOMAXPROCS(numThreads)
	var wg sync.WaitGroup

	wg.Add(numThreads)
	for i := range numThreads {
		go heavyTask(i, &wg)
	}
	wg.Wait()
}

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %d is starting...\n", id)
	for range 1000_000 {

	}
	fmt.Println(time.Now())
	fmt.Printf("Task %d has finished.\n", id)
}

func printNumbers() {
	for i := range 5 {
		fmt.Println(time.Now())
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printLetters() {
	for _, rune := range "ABCDE" {
		fmt.Println(time.Now())
		fmt.Println(string(rune))
		time.Sleep(time.Millisecond * 500)
	}
}
