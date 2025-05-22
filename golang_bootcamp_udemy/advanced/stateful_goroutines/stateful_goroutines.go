package main

import (
	"fmt"
	"time"
)

type statefulWorker struct {
	count int
	ch    chan int
}

func StartRoutine(sw *statefulWorker) {
	go func() {
		for {
			select {
			case receiver := <-sw.ch:
				sw.count += receiver
				fmt.Println("Updated Value of Count:", sw.count)
			default:
			}
		}
	}()
}


func main() {
	workerInstance := &statefulWorker{
		count: 0,
		ch : make(chan int),
	}
	StartRoutine(workerInstance)

	for i := range 5 {
		workerInstance.ch <- i
		time.Sleep(time.Millisecond*200)
	}

}
