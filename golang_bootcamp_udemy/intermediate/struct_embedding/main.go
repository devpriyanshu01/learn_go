package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person //struct embedding
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("My name is %s and I am %d year old.\n", p.name, p.age)
}

//method over-riding
func (e Employee) introduce() {
	fmt.Printf("I am employee with empId %s & my salary is - %.2f. \n",e.empId, e.salary )
}
func main() {
	emp := Employee{
		person: person{name: "raman", age: 23},
		empId:  "91734",
		salary: 917323.23,
	}
	//So here as it is evident, employee can directly access name and age fields 
	//as if they were part of Employee and that is what field promotion is all about.

	fmt.Println("Printing emp object :", emp)
	fmt.Println("Name:", emp.name)
	fmt.Println("Age:", emp.age)
	fmt.Println("Emp Id:", emp.empId)
	fmt.Println("Salary:", emp.salary)

	//Since, we've embedded person struct in Employee struct, we can 
	//use introduce method for Employee object as well.
	emp.introduce()

	//Let's now work with person struct.
	p1 := person{name: "Arjun", age : 55}
	fmt.Println("Printing p1 object of person struct --->", p1)
	p1.introduce()
}


