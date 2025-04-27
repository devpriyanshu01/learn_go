package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("----------- URL PARSING ---------")

	// [scheme://][username@info]host[:port][/path][?query][#fragment]
	//go package - net/url

	rawUrl := "https://example.com:8080/path?query=param#fragment"

	//parsing means - to breakdown something into it's components.
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error encountered while parsing =>", err)
	}
	fmt.Println("scheme =>", parsedUrl.Scheme)
	fmt.Println("host =>", parsedUrl.Host)
	fmt.Println("path =>", parsedUrl.Path)
	fmt.Println("query =>", parsedUrl.RawQuery)
	fmt.Println("fragment =>", parsedUrl.Fragment)
	fmt.Println("port =>", parsedUrl.Port())
	
	fmt.Println("------------- Example of Query Paramenter ---------------")
	rawUrl1 := "https://example.com:8080/path?name=raman&age=22#fragment=nil"
	parsedUrl1, err := url.Parse(rawUrl1)
	if err != nil {
		fmt.Println("Error =>", err)
		return
	}
	// queryParams := parsedUrl1.RawQuery //returns query_param as string
	queryParams := parsedUrl1.Query()	//return query_param as a map.
	fmt.Println("QueryParams =>", queryParams)
	fmt.Println("Name =>", queryParams.Get("name"))
	fmt.Println("Age =>", queryParams.Get("age"))

	//building url
	fmt.Println("---------- Building a Url -----------")
	
	baseUrl := &url.URL{
		Scheme: "https",
		Host: "example.com",
		Path: "/path",
	}
	query := baseUrl.Query()	//empty map
	query.Set("name", "raman")	//add a value to map
	query.Set("age", "10")
	baseUrl.RawQuery = query.Encode()	//attach query_params to base url

	fmt.Println("Base Url =>", baseUrl.String())

	//creating query param values to attach to a url
	values := url.Values{}
	values.Add("username", "raman_rocks")
	values.Add("email", "ramanrocks00@gmail.com")
	values.Add("name", "Raman Bhargav")
	values.Add("age", "23")

	//encode the values
	encodedForUrl := values.Encode()
	fmt.Println("Encoded For Url =>", encodedForUrl)
	baseUrl2 := "https://example.com/search"
	fullUrl := baseUrl2 + "?" + encodedForUrl
	fmt.Println("Full-Url =>", fullUrl)
}
