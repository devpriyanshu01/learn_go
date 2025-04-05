package main

import "fmt"

func main() {
	fmt.Println("-------------  Structs in Golang  -------------")

	//There is no concept of inheritance in Golang. Also, there nothing
	//like super or parent in Golang.

	newUser := User{"Raman", "raman@gamil.com", 12, true}
	fmt.Printf("Let's look at the new user:- \n %+v\n",newUser )

	//Also, we can print the fields separately.
	fmt.Printf("newUser name is %v and email is %v \n", newUser.Name, newUser.Email)

}

type User struct {
	Name   string
	Email  string
	age    int
	status bool
}
