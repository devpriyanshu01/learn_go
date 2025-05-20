package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://courses.chaicode.com/learn/my-enrollments?show=all&type=100"

func main() {
	fmt.Println("Welcome to handling URLs in Golang")

	//parse the url
	result, _ := url.Parse(myurl)
	fmt.Println(result)

	//print
	fmt.Println(result.Scheme)	//gives https
	fmt.Println(result.Path)	//gives the routing - /learn/my-enrollments
	fmt.Println(result.Host)	//gives - /learn/my-enrollments
	fmt.Println(result.RawPath)	//giving blank
	fmt.Println(result.Port())	//blank
	fmt.Println(result.RawQuery)	//show=all&type=100

	qParams := result.Query()
	fmt.Printf("%T\n", qParams)
	// fmt.Println(qParams)
	for index, value := range qParams {
		fmt.Printf("%v ==> %v\n", index, value)
	}

	fmt.Println("-------------- NEXT SECTION ---------------")

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host : "priyanshudev.in",
		Path : "learn/golang",
		RawPath: "user=raman",
	}

	fmt.Println("New url is - ", partsOfUrl.String())

}
