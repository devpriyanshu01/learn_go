package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func loginWithContext(ctx context.Context, message string) {
	requestIdVal := ctx.Value("requestId")
	log.Printf("Request ID: %v - %v", requestIdVal, message)
}

func main() {
	rootContext := context.Background()
	// ctx, cancel := context.WithTimeout(rootContext, 2*time.Second) //timer of context starts here.
	// defer cancel()

	ctx, cancel := context.WithCancel(rootContext)
	
	go func(){
		time.Sleep(2*time.Second)
		cancel()
	}()


	ctx = context.WithValue(ctx, "requestId", "akdfjowjfasdfffsls")

	// go doWork(ctx)
	go loginWithContext(ctx, "hare krsna")
	time.Sleep(4 * time.Second)

	reqId := ctx.Value("requestId")
	if reqId != nil {
		fmt.Println("request Id:", reqId)
	} else {
		fmt.Println("no request Id found.")
	}
}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()

// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result with context.TODO():", result)

// 	ctx = context.Background()
// 	ctx,_ = context.WithTimeout(ctx, 1*time.Second)
// 	// defer cancel()

// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("Result from timeout context:", result)

// 	time.Sleep(3 * time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println("Result after timeout:", result)
// }

// ------- DIFFERENCE BETWEEN context.TODO() & context.Background() ---------
// func main() {
// 	contextTodo := context.TODO()	//create new context
// 	ctx := context.WithValue(contextTodo, "title", "Understand Context")
// 	fmt.Println("ctx:", ctx)
// 	fmt.Println(ctx.Value("title"))

// 	contextBg := context.Background()	//create new context
// 	ctx2 := context.WithValue(contextBg, "city", "Chandigarh")
// 	fmt.Println(ctx2)
// 	fmt.Println(ctx2.Value("city"))
// }
