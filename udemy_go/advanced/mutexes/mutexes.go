package main

import (
	"fmt"
	"sync"
)

// Example 2: Understanding Mutex more.
//How mutex protect the block of code.
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int

	numGoroutines := 2500

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}

	wg.Add(numGoroutines)
	for range numGoroutines {
		go increment()
	}
	wg.Wait()
	fmt.Println("Updated Value of count variable is:", count)
}

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
//So what we see is a program that creates a counter
//that multiple go routines can safely increment without
//interfering with each other.
func main() {
	var wg sync.WaitGroup
	counter := &counter{}
	numGoroutines := 10
	// wg.Add(numGoroutines)

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// counter.count++
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Final counter value:%d\n", counter.getValue())
}
