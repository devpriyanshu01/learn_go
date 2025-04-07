package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("\\\\\\\\ Working with files in Go ////////")

	content := "This needs to be added - LearnCodeOnline. "

	file, err := os.Create("./logFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./logFile.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))

}
