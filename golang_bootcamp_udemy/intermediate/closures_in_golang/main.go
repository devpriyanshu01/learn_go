package main

import "fmt"

func main() {
	fmt.Println("Let's learn Closures in Golang")

	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	fmt.Println("----------- Working on Sequence2 -------------")

	sequence2 := adder()
	fmt.Println(sequence2())
	fmt.Println(sequence2())

	fmt.Println("------------ CLOSURES SECOND EXAMPLE ------------")

	subtracter := func() func(int) int {
		countDown := 99

		return func(x int) int {
			countDown -= x
			return countDown
		}
	}()

	//using subtractor closure
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))

}

func adder() func() int {
	i := 0
	fmt.Println("Value of i is: ", i)

	return func() int {
		i++
		fmt.Println("Added 1 to value of i")
		return i
	}
}
