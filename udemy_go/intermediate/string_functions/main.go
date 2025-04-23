package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("--------------- Welcome to learn string functions in Golang --------------")

	//String Functions Discussed:
	// 1. strconv.Itoa()
	// 2. strings.Split()
	// 3. strings.Contains()
	// 4. strings.Join()
	// 5. strings.Replace()
	// 6. strings.TrimSpace()
	// 7. strings.HasPreffix()
	// 8. strings.HasSuffix()
	// 9. strings.ToLower()
	// 10. strings.ToUpper()
	// 11. strings.Repeat()
	// 12. strings.Count()

	// 1. strconv.Itoa()
	num := 23
	numInInt := strconv.Itoa(num)
	fmt.Println("Printing converted num var", numInInt)
	fmt.Println("Printing type of ", reflect.TypeOf(num))

	// 2. strings.Split()
	toSplit := "Raman, Aman, Shaman, Daman"
	fmt.Println("Type of toSplit -", reflect.TypeOf(toSplit))
	splitted := strings.Split(toSplit, ",") //return a slice
	fmt.Println("Splitted - ", splitted)
	fmt.Println("Type of splitted -", reflect.TypeOf(splitted))

	// 3. strings.Contains()
	subStr := "he"
	parentStr := "hello world"
	fmt.Println(strings.Contains(parentStr, subStr)) //true
	fmt.Println(strings.Contains(parentStr, "He"))   //false

	// 4. strings.Join()
	randomSlice := []string{"Raman", "Radha", "Shyama", "Sudama"}
	joined := strings.Join(randomSlice, "---")
	fmt.Println("Joined =>", joined)

	// 5. strings.Replace()
	randomString := "Hare Krsna Hare Krsna"
	replacedString := strings.Replace(randomString, "Krsna", "Ram", -1)
	fmt.Println("Replaced String =>", replacedString)

	// 6. strings.TrimSpace()
	withSpaceString := "     This is just a randome   string  .  "
	spaceTrimmedString := strings.TrimSpace(withSpaceString)
	fmt.Println("Space Trimmed String =>", spaceTrimmedString)

	// 7. strings.HasPreffix()
	randomString2 := "Hare Krsna Hare Krsna"
	fmt.Println("Has Preffix =>", strings.HasPrefix(randomString2, "Hare"))

	// 8. strings.HasSuffix()
	randomString3 := "Hare Krsna Hare Krsna"
	fmt.Println("Has Preffix =>", strings.HasSuffix(randomString3, "Krsna"))

	//Let's see some regular expression
	randomString4 := "Hello, 1234 say it like 23 22sss2"

	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(randomString4, -1)
	fmt.Println("Matches =>", matches)

	//String Builder - efficient compare to normal string operations.
	var builder strings.Builder

	//create string using string builder
	builder.WriteString("Govindam ")
	builder.WriteString("Aadipurusam Tamham ")
	builder.WriteString("Bhajami.")

	//convert the string builder to string
	result := builder.String()
	fmt.Println("String-Builder converted to String =>", result)

	//Let's also use string-builder for Runes.
	builder.WriteRune('=')
	builder.WriteString("HareKrsnaHareKrsna")

	result = builder.String()
	fmt.Println("Final Result =>", result)

	//reset the builder
	builder.Reset()	//builder is now empty
	builder.WriteString("String after restting the builder")
	//builder to string conversion
	result = builder.String()
	fmt.Println("After restting the builder =>", result)



}
