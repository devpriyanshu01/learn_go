package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("------- Write to a File -------")

	file, err := os.Create("output_file.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	//write data to file
	data := []byte("Hello Sir! ")
	// fmt.Println(data)
	bytesWritten, err := file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}else{
		fmt.Println("Write to file successful - Bytes Written:", bytesWritten)
	}

	//Writing string to a file.
	fileString, err := os.Create("write_string.txt")
	if err != nil {
		fmt.Println("error creating file:", fileString)
		return
	}
	defer fileString.Close()

	//write string to this file
	bytesWritten, err = fileString.WriteString("Hello Golang!\n\n")
	if err != nil {
		fmt.Println("error writing to file:", err)
		return
	}else {
		fmt.Println("Write successful - Bytes written", bytesWritten)
	}


}
