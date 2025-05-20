package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("----- BASE64 CODING -----")

	data := []byte("He~~lo, Base64 encoding")

	//encode to Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded :", encoded)

	//decode back
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(decoded))

	//URL Safe Encoding: avoid '/' & '+'
	urlSafeEncoding := base64.URLEncoding.EncodeToString(data)
	fmt.Println("url safe encoding:", urlSafeEncoding)

}
