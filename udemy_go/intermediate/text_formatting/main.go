package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	fmt.Println("-------- Let's learn text formatting -----------")

	// tmpl := template.New("example")

	// tmpl, err := template.New("example").Parse("Welcome, {{.name}} How are you doing?\n")
	// if err != nil {
	// 	panic(err)
	// }

	//Above 4 lines can be minimized to 1 line using template.Must() as it uses error
	//handling internally. So we don't we do - if err != nil

	// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}} How are you doing?\n"))

	// //Define data for welcome message string template
	// data := map[string]interface{}{
	// 	"name" : "Raman Kumar",
	// }

	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Define named templates for different types of 
	templates := map[string]string {
		"welcome" : "Welcome, {{.name}}! We're glad you joined.",
		"notification" : "{{.name}} - You've new notifcation: {{.notification}}",
		"error" : "Oops! An error occured: {{.errorMessage}}",
	}
}
