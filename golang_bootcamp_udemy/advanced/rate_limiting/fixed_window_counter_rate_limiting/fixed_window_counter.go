package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	resetTime time.Time
	duration  time.Duration
}

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		limit:    limit,
		duration: duration,
	}
	return rl
}

func (rl *RateLimiter) Allowed() bool {
	now := time.Now()

	rl.mu.Lock()
	defer rl.mu.Unlock()

	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.duration)
		rl.count = 0
	}

	if rl.count <= rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {
	var wg sync.WaitGroup

	rateLimit := NewRateLimiter(5, time.Second)
	for range 10 {
		wg.Add(1)
		go func(){
			if rateLimit.Allowed() {
				fmt.Println("request allowed.")
			}else {
				fmt.Println("request not allowd")
			}
			// time.Sleep(time.Millisecond*200)
			wg.Done()
		}()
	}
	wg.Wait()
}
