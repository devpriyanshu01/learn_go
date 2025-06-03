package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("---------- BUFIO PACKAGE ----------")

	// reader := bufio.NewReader(strings.NewReader("Hello, Bufio Package!. You are awesome!!!\n"))

	// //Reading byte slice
	// data := make([]byte, 20)	//allocate space for storing the read data.
	// n, err := reader.Read(data)	//read & store it in data var
	// if err != nil {
	// 	fmt.Println("Error Reading: =>", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes, data:- %s \n", n, data[:])	//print the data

	// //Two functions:
	// //1. reader.Read() - reads till specified no. of byte.
	// //2. reader.ReadString() - reads till a specific char/rune.

	// storer, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Read String output =>", storer)

	//Writer
	writer := bufio.NewWriter(os.Stdout) //target

	//1. write byte slice
	data := []byte("Hello, Bufio Package!!!\n") //source
	bytesWritten, err := writer.Write(data)     //return no. of bytes written & writes to a buffer.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("No. of Bytes written =>", bytesWritten)

	//Flush the buffer to ensure all the data is written to os.Stdout(target).
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	//2. Writing String
	//creating source
	str := "This is the string that needs to be written to os.Stdout..."
	bytesWritten, err = writer.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("No. of Bytes written when writing string to os.Stdout =>", bytesWritten)

	//Flush the buffer to os.Stdout
	err = writer.Flush()
	fmt.Println("\n")
	if err != nil {
		fmt.Println(err)
		return
	}

}
