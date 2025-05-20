package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Let's learn Get Requests in Go")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const postUrl = "http://localhost:3000/echo"

	//fake json body
	requestBody := strings.NewReader(`
		{
			"courseName" : "Let's learn Golang",
			"price" : 0,
			"tutor" : "Hitesh",
			"platform" : "YouTube"
		}
	`)

	//do a post request
	response, err := http.Post(postUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response is :- \n \n", string(content))
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/"

	res, err := http.Get(myUrl)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Response Status Code ---> ", res.StatusCode)
	fmt.Println("Response Content Length ---> ", res.ContentLength)

	content, _ := io.ReadAll(res.Body)	//content is in byte format. So,
	// fmt.Println(string(content))		//we need to convert it to string.

	//Alternate way to convert this bytes which is in content variable to string using String Builder.
	var resString strings.Builder
	byteCount, _ := resString.Write(content)
	fmt.Println("ByteCount is- ", byteCount)
	//How we get the actual content?? See Below
	fmt.Println("Readable Response - ", resString.String())


}
