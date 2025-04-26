package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("--------- Let's learn regular expressions in golang ---------")

	//Escape Sequences :
	fmt.Println("He said, \"I am Great!\"")
	//				Or
	fmt.Println(`He said, "I am Great!`)

	//compile a regex pattern to validate email addresses.
	re := regexp.MustCompile(`[A-Za-z0-9_.%+-]+@[A-Za-z0-9]+\.[a-zA-Z]{2,}`)

	email1 := "raman00268@gmail.com"
	email2 := "invalidemail"

	fmt.Println("email1 valid??", re.MatchString(email1))
	fmt.Println("email2 valid??", re.MatchString(email2))

	//compile a regex pattern to validate date.
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date1 := "2024-12-21"
	fmt.Println(re.MatchString(date1))	//true
	subMatches := re.FindStringSubmatch(date1)
	fmt.Println(subMatches)	//[2024-12-21 2024 12 21]
	fmt.Println("subMatches[1] -", subMatches[1])//	subMatches[1] - 2024
	fmt.Println("subMatches[2] -", subMatches[2])// subMatches[2] - 12
	fmt.Println("subMatches[3] -", subMatches[3])// subMatches[3] - 21
	
	//replace source string
	str := "Hello World!"

	re = regexp.MustCompile(`[aeiou]`)
	fmt.Println("replaced string:", re.ReplaceAllString(str, "*"))

	//i - case insensitive
	//m - multi line model
	//s - dot matches all

	re = regexp.MustCompile(`(?i)go`)
	//test string
	text := "GOlang is great"

	//match
	fmt.Println("Match -", re.MatchString(text))
}
