package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-------- Let's learn how to handle time in Golang ---------")

	fmt.Println("Time Now -", time.Now())	//current time

	//create a custom time
	createTime := time.Date(2025, time.April, 21, 22, 00, 00, 00, time.Local)
	fmt.Println("Created Time -", createTime)

	//print a time in YYYY-MM-DD format
	parsedTime, _ := time.Parse("2006-01-02", "2025-12-15")	//Mon Jan 2 15:04:05 MST 2006
	fmt.Println("Parsed Time -", parsedTime)

	//date with time
	parsedTime2,err := time.Parse("2007-01-02 15-04-05", "1998-12-15 15-30-01")
	if(err != nil){
		fmt.Println("err - ", err)
	}
	fmt.Println("Parsed Date Time -", parsedTime2)

	//format time
	t := time.Now()
	fmt.Println("Formatted Time :-", t.Format("Mon 2006-01-02 15:04:05"))

	//set time to 1 day later
	t = time.Now()
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("After 1 day ->", oneDayLater)
	fmt.Println(oneDayLater.Weekday())	//get the date

	loc, _ := time.LoadLocation("Asia/Kolkata")
	timeNew := time.Date(2025, time.April, 26, 13, 03, 30, 200, time.UTC)
	// timeNew := time.Now()

	//convert timeNew wiz UTC to local Indian Time.
	timeInLocal := timeNew.In(loc)

	fmt.Println("Time in UTC =>", timeNew)
	fmt.Println("Time in Local =>", timeInLocal)

	//perfom time rounding
	roundedTime := timeNew.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)
	fmt.Println("Rounded Time UTC =>", roundedTime)
	fmt.Println("Rounded Time Local =>", roundedTimeLocal)

	//Let's current time New York
	locNY, _ := time.LoadLocation("America/New_York")

	timeInNY := time.Now().In(locNY)
	fmt.Println("Current time in New York =>", timeInNY)

	//calculate the difference b/n the time.
	time1 := time.Date(2025, time.April, 26, 14, 50, 30, 0, time.Local)
	time2 := time.Date(2025, time.April, 27, 14, 50, 30, 0, time.Local)

	difference := time2.Sub(time1)
	fmt.Println("Difference in time is =>", difference)

	//comparing times
	fmt.Println("Is time2 after time1? =>", time2.After(time1))

}
