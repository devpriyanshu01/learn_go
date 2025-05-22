package main

import (
	"fmt"
	"sort"
)

func main() {
	intArray := []int{5, 2, 39, 22, 21}
	sort.Ints(intArray)
	fmt.Println("sorted array:", intArray)

	stringArray := []string{"Hare", "Krsna", "Ram", "Apple"}
	sort.Strings(stringArray)
	fmt.Println("sorted array:", stringArray)
}
