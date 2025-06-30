package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter1 := &Counter{}

	numGoroutines := 10
	var wg sync.WaitGroup
	
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 10000 {
				counter1.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("count value:", counter1.getCount())
}
