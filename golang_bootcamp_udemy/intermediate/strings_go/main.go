package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	greet := "Hello \nGo"
	message1 := `Hello \nGo`
	message2 := "Hello\tGo"
	message3 := "Hello \rGo!"

	fmt.Println(greet)		//Hello
							//Go
	fmt.Println(message1)	//Hello \nGo
	fmt.Println(message2)	//Hello   Go
	fmt.Println(message3)	//Go!lo

	fmt.Println("------- Working with Length ---------")
	fmt.Println(len(greet))	//return the byte of the string variable.
	fmt.Println(len(message1))
	fmt.Println(len(message2))

	fmt.Printf("%c\n", message1[0])

	fmt.Println("-------------- LEXICAL GRAPHICAL COMPARISON -----------")

	fmt.Println("apple" > "Apple") //97 > 65 = true
	fmt.Println("app" > "apple") //false

	fmt.Println("----------------- PLACEHOLDERS OR FORMAT ENCODERS -------------")

	str1 := "Hello"
	for i, char := range str1 {
		fmt.Printf("%d \t %c\n", i, char)
		// fmt.Printf("%x \n", char)	//hexadecimal values
		// fmt.Printf("%v \n", char)
	}

	fmt.Println("Rune Count: ", utf8.RuneCountInString("Hello"))


	fmt.Println("-----======== Runes ------------")

	var ch rune = 'a'
	jch := '日'

	fmt.Println("a - ", ch)
	fmt.Println("jch - ", jch)

	cstr := string(ch)
	fmt.Println("cstr ", cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	fmt.Println("----------- JAPANESE STRING ------------")

	japaneseChar := "月曜日"
	fmt.Println("Japanese String :- ", japaneseChar)

	for i, char := range japaneseChar {
		fmt.Printf("%d  %c\n", i , char)
	}

	japaneseHello := "こんにちは"
	fmt.Println("Japanese Hello - ", japaneseHello)
	
	for _, value := range japaneseHello {
		fmt.Printf("%c \n", value)
	}

}
