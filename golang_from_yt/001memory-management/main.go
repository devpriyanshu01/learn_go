package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	a := new(Person)
	fmt.Println("Person - ", *a)

	b := make(map[string]string)
	fmt.Println("b :- ", b)

	c := Person{name : "name"}
	fmt.Println("printing c==>", c)

	p1 := newPerson("krsna krsna")
	fmt.Println("printing p1 +> ", *p1)

}

func newPerson(name string) *Person {
	p := Person{name : name}
	p.age = 23
	return &p
}
