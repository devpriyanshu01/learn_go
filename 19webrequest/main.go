package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://staging-kitchen.priyanshudev.in"

func main() {
	fmt.Println("LCO web request")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of error is => %T", err)
	fmt.Printf("Type of error is => %T", res)

	defer res.Body.Close()
	//It's caller responsibility to close the connection.

	dataBytes, err := io.ReadAll(res.Body)
	
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println(content)

}
