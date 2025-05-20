package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLPerson xml.Name `xml:"person"`
	Name      string   `xml:"name,omitempty"`
	Age       int      `xml:"age,omitempty"`
	City      string   `xml:"city,omitempty"`
	Email     string   `xml:"email,omitempty"`
	Address Address `xml:"address,omitempty"`
}
type Address struct {
	City string	`xml:"city,omitempty"`
	State string `xml:"state,omitempty"`
}
func main() {
	person1 := Person{
		Name:  "Prakash Choudhary",
		Age:   24,
		City:  "Jaipur",
		Email: "prakashchoudhary@gmail.com",
		Address: Address{
			City: "Jaipur",
			State : "Rajasthan",
		},
	}
	//marshel to xml
	xmlPerson1, err := xml.Marshal(person1)
	if err != nil {
		fmt.Println("error marshelling:", err)
		return
	}
	fmt.Println("xmlPerson1 w/o Indent:", string(xmlPerson1))

	fmt.Println("---------------------------------------------------------------")
	xmlPerson1Indented,err := xml.MarshalIndent(person1,"","  ")
	if err != nil {
		fmt.Println("error marshelling:", err)
		return
	}
	fmt.Println("xmlPerson1 Indented:", string(xmlPerson1Indented))
	fmt.Println("---------------------------------------------------------------")

	//unmarshelling
	xmlData1 := `<person><name>Prakash Choudhary</name><age>25</age><address><city>jaipur</city><state>rajasthan</state></address></person>`
	var xmlData1Unmarshalled Person
	err = xml.Unmarshal([]byte(xmlData1),&xmlData1Unmarshalled)
	if err != nil {
		log.Fatalln("error unmarshelling:", err)
		return
	}
	fmt.Println("unmarshalled data:", xmlData1Unmarshalled)
	fmt.Println("--------------------------------------------------------------")
	book1 := Book{
		ISBN: "9342-934-5-792-345",
		Title: "Stress Management",
		Author: "HG RadheShyam Pr",
		Pseudo: "ajljefj2of",
		PseduoAttr: "lsajflasjdf",
	}
	
	//marshell
	book1Xml, err := xml.MarshalIndent(book1, "", "  ")
	if err != nil {
		log.Fatalln("error marshalling:", book1Xml)
		return
	}
	fmt.Println("book1 as xml:", string(book1Xml))
}
//declaring attribute in struct tags.
type Book struct {
	XMLName	xml.Name `xml:"book"`
	ISBN	string	`xml:"isbn,attr"`
	Title	string `xml:"title,attr"`
	Author	string	`xml:"author,attr"`	//attr means - consider as attribute.
	Pseudo	string 	`xml:"pseduo"`
	PseduoAttr string `xml:"pseudoattr,attr"`
}
