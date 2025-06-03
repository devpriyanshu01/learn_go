package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("----- Reading Files in Golang ------")

	content, err := os.ReadFile("output.txt")
	// if err != nil {
	// 	fmt.Println("error reading file:", err)
	// 	return
	// }else {
	// 	fmt.Println("file content:", content)
	// }

	//Open a file.
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}else {
		fmt.Println("File opened successfully!")
	}
	
	defer func() {
		fmt.Println("Closing the file")
		file.Close()
	}()

	// //read contents of the open file at once.
	// data := make([]byte, 1024)	//buffer to read the file
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("error reading file:", err)
	// 	return
	// }
	
	// fmt.Println("content of the file:", string(data))

	//Read a file line by line.
	scanner := bufio.NewScanner(file)

	//read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("err reading file line by line:", err)
		return
	}

}
