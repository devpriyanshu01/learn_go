package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome Sir!!!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the Pizza.")

	//comma ok || comma err syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your ", input, "rating, sir.")
}
