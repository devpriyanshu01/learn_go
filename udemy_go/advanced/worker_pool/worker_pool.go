package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID int
	numTickets int
	cost int
}

//simulate processing of ticket requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int){
	for req := range requests {
		fmt.Printf("processing %d ticket(s) of personID %d with total cost %d\n", req.numTickets, req.personID, req.cost)
		//simulate processing time
		time.Sleep(time.Second)
		results<- req.personID
	}
}

func main(){
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	//start ticket processor/workers
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	//send ticket requests
	for i := range numRequests{
		ticketRequests<- ticketRequest{personID : i+1, numTickets : (i+1)*2, cost :(i+1)*price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("ticket of the personID %d processed successfully!\n", <-ticketResults)
	}
}

// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("worker %d processing task %d\n", id, task)
// 		//simulate work
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }
// //This is a design pattern for executing more no. of tasks with less no. of workers.
// func main() {
// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	//create workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	//send values to tasks channel
// 	for i := range numJobs{
// 		tasks <- i
// 	}

// 	close(tasks)

// 	//collect the results
// 	for range numJobs{
// 		res := <-results
// 		fmt.Println("results:", res)
// 	}
// }
