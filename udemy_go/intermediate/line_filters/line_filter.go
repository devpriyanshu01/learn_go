package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("------- Line Filters -------")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}else{
		fmt.Println("file opened successfully")
	}

	//read file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "hare"){
			fmt.Println("filtered line:", line)	//filterd line on condition.
			modifiedLine := strings.ReplaceAll(line, "hare", "radhe")
			fmt.Println("modified line:", modifiedLine)
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("error reading file line by line:", err)
		return
	}
}
