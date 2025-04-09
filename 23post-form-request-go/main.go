package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("------------------- Post Form Request in Golang ---------------")

	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const postUrl = "http://localhost:3000/postform"

	//creating a test formData
	data := url.Values{}
	data.Add("firstName", "Raman")
	data.Add("LastName", "Kumar")
	data.Add("email", "raman008@gmail.com")

	response, err := http.PostForm(postUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occured -", err)
		panic(err)
	}
	fmt.Println("Readable Response is- \n\n ", string(content))

}

