package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-------- EPOCH --------")

	//00:00:00 UTC on Jan 1, 1970

	now := time.Now()	//current time
	unixTime := now.Unix()	//epoch time
	fmt.Println("Current Unix Time:", unixTime)	//print epoch time

	t := time.Unix(unixTime, 0)	//convert epoch time to unix format
	fmt.Println(t)
	fmt.Println("Time:", t.Format("2006-01-02"))

}
