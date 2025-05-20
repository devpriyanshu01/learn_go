package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("----- Environment Variables in Golang -----")

	user := os.Getenv("USER")
	home := os.Getenv("HOME")
	fmt.Println("User:", user)
	fmt.Println("Home:", home)

	//set a new env variable
	os.Setenv("FRUIT", "Guava")
	fmt.Println(os.Getenv("FRUIT"))

	//Get all the environment variables.
	allEnvVars := os.Environ()
	fmt.Println("All env's:\n", allEnvVars)

	fmt.Println("--------------------------------------------------------")
	for _, oneEnv := range allEnvVars {
		keyValue := strings.Split(oneEnv, "=")
		fmt.Println(keyValue[0])
	}
}
