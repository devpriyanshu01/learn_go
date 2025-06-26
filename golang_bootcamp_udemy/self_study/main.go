package main

import (
	"bytes"
	"fmt"
	"os"
)

// UNDERSTANDING io.Writer
func main() {
	data := []byte("Rome was not built in a day.")
	fmt.Println("len of data", len(data))

	var buf bytes.Buffer
	n, err := buf.Write(data)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("bytes read", n)

	bytesRead, err := os.Stdout.Write(data)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("bytes read to stdout", bytesRead)
}

//UNDERSTANDING io.Reader
// func main() {
// 	stringReader := strings.NewReader("Hare krsna hare krsna krsna krsna hare hare, hare ram hare ram ram ram hare hare.")

// 	buffer := make([]byte, 10)

// 	for {
// 		n, err := stringReader.Read(buffer)
// 		if err == io.EOF {
// 			fmt.Println("END OF FILE.")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("Error reading file from the source.")
// 			break
// 		}
// 		fmt.Println("bytes read:", n, "data:", string(buffer))
// 	}
// }
