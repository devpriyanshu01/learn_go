package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

// stack struct
type Stack[T any] struct {
	elements []T
}

// push elements to stack
func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

// pop element from stack
func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	//store last element to return
	element := s.elements[len(s.elements)-1]
	//remove last element
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

// check if Stack is empty
func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	fmt.Println("----------------- Generics in Golang -----------------")

	// x, y := 2, 3
	// x, y = swap(x, y)
	// fmt.Println("Swapped Value:-", x, y)

	// s1, s2 := "John", "Jane"
	// s1, s2 = swap(s1, s2)
	// fmt.Println("Swapped Value:- ", s1, s2)

	//Construct Stack Struct.
	intStack := Stack[int]{
		elements: []int{1, 2, 3},
	}
	fmt.Println("intStack =>", intStack)

	//push elements to above stack.
	intStack.push(4)
	intStack.push(5)
	intStack.push(6)
	fmt.Println("After pushing elements =>", intStack)

	//pop elements from above stack.
	intStack.pop()
	fmt.Println("After popping last element => ", intStack)

	//check if stack is now empty
	if intStack.isEmpty() {
		fmt.Println("Stack is Empty!")
	} else {
		fmt.Println("Stack is not Empty!")
	}

	//Creating a empty String Stack
	stringStack := Stack[string]{}
	//check if stack is empty
	if stringStack.isEmpty() {
		fmt.Println("stringStack is empty")
	} else {
		fmt.Println("stringStack is not empty")
	}

	//push element
	stringStack.push("India")
	stringStack.push("Is")
	stringStack.push("Great")
	fmt.Println("Printing our pushing elements =>", stringStack)

	//pop element
	stringStack.pop()
	fmt.Println("Printing after popping =>", stringStack)
}
