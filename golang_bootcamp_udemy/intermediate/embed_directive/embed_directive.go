package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var fileContent string

//go:embed basics
var basicsFolder embed.FS


func main() {
	fmt.Println("-------- Embed Directive ---------")

	fmt.Println("embedded content:", fileContent)
	content, err := basicsFolder.ReadFile("basics/hello.txt")
	
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("basic folder content:", string(content))

	//fs.WalkDir
	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("err")
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
