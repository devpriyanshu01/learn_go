package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("------- FILE PATH -------")

	//join paths
	joinedPath := filepath.Join("home/rab-5/", "Downloads", "filePath.txt")
	fmt.Println("joined path:", joinedPath)

	//clean paths
	cleanedPath := filepath.Clean("./data/../data/text.txt")
	fmt.Println("cleaned path:", cleanedPath)

	//split path into directory & file.
	dir, file := filepath.Split("home/rab-5/Downloads/filePath.txt")
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)

	//base in a file_path
	fmt.Println("Base?:", filepath.Base("home/rab-5/Downloads/filePath"))

	//absolute & relative paths.
	absolutePath := "/Desktop/Learnings/golang/udemy_go/intermediate/file_paths"
	relativePath := "../../udemy.go/random_numbers"

	fmt.Println("Is filepath absolute? :", filepath.IsAbs(relativePath))
	fmt.Println("Is filepath absolute? :", filepath.IsAbs(absolutePath))

	//calculate relative path from given base path & target path.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		fmt.Println("error calculating relative path:", err)
		return
	}
	fmt.Println("relative path:", rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		fmt.Println("error calculating relative path:", err)
		return
	}
	fmt.Println("relative path:", rel)

	//convert relative path to absolute path.
	relativePath = "../../udemy.go/random_numbers"
	convertedAbsolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("error converting relative to absolute path:", err)
		return
	}
	fmt.Println("converted absolute path:", convertedAbsolutePath)

}
