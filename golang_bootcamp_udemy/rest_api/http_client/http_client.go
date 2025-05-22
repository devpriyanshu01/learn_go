package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{} //create new http client

	res, err := client.Get("https://dummyjson.com/users/1")
	if err != nil {
		fmt.Println("error in client.get:", err)
		return
	}
	defer res.Body.Close()

	//read and print the response body
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err reading res.Body", err)
		return
	}
	fmt.Println("res is:", string(bytes))
}
