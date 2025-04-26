package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------- Time formatting -----------")

	// Mon Jan 2 15:04:05 MST 2006

	//time formatting 1
	layout := "2006-01-02T15:04:05Z07:00"
	str := "2025-04-26T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("error =>", err)
	}
	fmt.Println("Formatted Time =>", t)

	//time formatting 2
	layout2 := "Jan 02, 2006 3:04:05 PM"
	str2 := "Apr 26, 2025 4:51:30 PM"
	
	t2, err := time.Parse(layout2, str2)
	if err != nil {
		fmt.Println("Error =>", err)
	}
	fmt.Println("Formatted time2 =>", t2)
}
