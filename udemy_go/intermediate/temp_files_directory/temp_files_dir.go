package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("------- Temp Files & Directories --------")

	//creating temp files
	tempFile, err := os.CreateTemp("", "tempFile")	//create tempfile
	checkErr(err)

	fmt.Println("Temp File created:", tempFile.Name())

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	//creating temp directory
	tempDir, err := os.MkdirTemp("", "go_temp_dir")
	checkErr(err)

	fmt.Println("temp dir created:", tempDir)

	defer os.Remove(tempDir)	//remove the temp dir to avoid accumulation.
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
