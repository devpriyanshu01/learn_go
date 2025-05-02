package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("-------- Directory --------")

	//create new directory
	// err := os.Mkdir("new_dir", 0700)
	// checkError(err)

	//remove directory
	// err := os.RemoveAll("create_dir")
	// checkError(err)

	//write to a file
	//to be written successfully - parent directory of the file should exists.
	err2 := os.WriteFile("new_dir/test.txt", []byte("Hare Krsna!!!"), 0700) 
	checkError(err2)

	//create nested directory
	err3 := os.MkdirAll("grand_dir/parent_dir/child1/test_file.txt", 0700)
	checkError(err3)
	err4 := os.MkdirAll("grand_dir/parent_dir/child2/test_file.txt", 0700)
	checkError(err4)
	err5 := os.MkdirAll("grand_dir/parent_dir/child3/test_file.txt", 0700)
	checkError(err5)

	//read directory
	result, err6 := os.ReadDir("grand_dir/parent_dir")
	checkError(err6)

	for _, entry := range result {
		fmt.Println(entry)
	}

	//change directory & look the look it's contents.
	//current working directory
	currDir, err7 := os.Getwd()
	checkError(err7)

	fmt.Println("current working directory:", currDir)

	//change directory now
	err8 := os.Chdir("./grand_dir/")
	checkError(err8)
	currDir, err9 := os.Getwd()
	checkError(err9)
	fmt.Println("current directory:", currDir)

	//read the content of the changed dir.
	result, err10 := os.ReadDir("./parent_dir")
	checkError(err10)
	for _, entry := range result {
		fmt.Println(entry)
	}

	//filepath.Walk & filepath.WalkDir
	fmt.Println("-------- WalkDir Learning --------")
	filePath := "./parent_dir"
	filepath.WalkDir(filePath, func (path string, d os.DirEntry, err error) error{
		if err != nil {
			fmt.Println("Error_Walkdir:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	

}

func checkError(err error){
	if err != nil {
		fmt.Println("error occured:", err)
		return
	}	
}
