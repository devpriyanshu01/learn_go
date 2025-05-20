package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("----- NUMBER PARSING -----")

	//convert string to int
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	fmt.Println("num type =>", reflect.TypeOf(num))

	if err != nil {
		fmt.Println("error in pasrsing: =>", err)
	}
	fmt.Println("Parsed Value =>", num)
	fmt.Println("Parsed Value + 10 =>", num+10)
	fmt.Println("--------------------------------------------------------------")
	//convert string to int64 type
	num64, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("error in parsing int value: =>", err)
	}
	fmt.Println("Num64 =>", num64)
	fmt.Println("type of num64 =>", reflect.TypeOf(num64))
	fmt.Println("--------------------------------------------------------------")

	//convert string to float
	strFloat := "3.14"
	float64Value, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("error in parsing float string: =>", err)
	}
	fmt.Println("Float64 Value =>", float64Value)
	fmt.Println("float64 type =>", reflect.TypeOf(float64Value))
	fmt.Println("--------------------------------------------------------------")

	//binary to decimal conversion.
	binaryStr := "11111111"
	decimalValue, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error in parsing binary str =>", err)
	}
	fmt.Println("decimalValue =>", decimalValue)
	fmt.Println("Type of var decimalValue =>", reflect.TypeOf(decimalValue))
	fmt.Println("--------------------------------------------------------------")

	//conversion of string hexadecimal value to int64
	hexStr := "FF"
	hexInt, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Err in parsing hex value to int64 =>", err)
		return
	}
	fmt.Println("Value of var hexInt =>", hexInt)
	fmt.Println("Type of var hexInt", reflect.TypeOf(hexInt))

	fmt.Println("--------------------------------------------------------------")

	//trying to parse an invalid String to Integer.
	invalidIntStr := "abc1234"
	invalidParse, err := strconv.ParseInt(invalidIntStr, 10, 64)
	if err != nil {
		fmt.Println("Error occured in parsing =>", err)
		return
	}
	fmt.Println("invalidParse Value", invalidParse)
}
