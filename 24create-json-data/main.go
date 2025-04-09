package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int    `json:"price"`
	Platform string
	Password string   `json:"-"` //don't include when consuming
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("-------------- Creating JSON Data in Golang ---------------")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Golang Bootcamp", 5000, "www.youtube.com", "12341234", []string{"golang complete playlsit - by hitesh"}},
		{"Python Bootcamp", 5000, "www.youtube.com", "12341234", []string{"Python complete playlsit - by hitesh"}},
		{"Nodejs Bootcamp", 5000, "www.learncodeonline.in", "12341234", []string{"nodejs complete playlsit - by hitesh"}},
	}

	// lcoCourses := course{
	// 	"Golang Bootcamp", 5000, "www.youtube.com", "232342", []string{"golang cohort by Hitesh"},
	// }

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}
