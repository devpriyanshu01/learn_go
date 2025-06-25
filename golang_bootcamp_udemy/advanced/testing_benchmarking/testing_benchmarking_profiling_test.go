package main

import (
	"math/rand"
	"testing"
)

//create a fn. that return slice of ints of given size.
func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		gv := rand.Intn(100)
		slice[i] = gv
	}
	return slice
}

//fn. to calculate the sum of the above slice generated.
func calculateSum(slice []int) int {
	ans := 0
	for _, v := range slice {
		ans = ans + v
	}
	return ans
}

//write a test for generateSlice fn.
func TestGenerateSlice(t *testing.T){
	size := 1000
	slice := generateSlice(1000)
	if len(slice) != size {
		t.Errorf("Expected Slice Size = %d, Received = %d", size, len(slice))
	}
}

//write Benchmarking fn for generateSlice fn.
func BenchmarkGenerateSlice(b *testing.B) {
	for range b.N {
		generateSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := generateSlice(1000)
	b.ResetTimer()	//resetTimer() is used so that once the set up are done then only count the time after resetting the time.
	for range b.N {
		calculateSum(slice)
	}
}

 
// func add(a int, b int) int {
// 	return a + b
// }

// func BenchmarkAddSmallInput(b *testing.B) {
// 	for range b.N {
// 		add(2, 3)
// 	}
// }

// func BenchmarkAddMediumInput(b *testing.B) {
// 	for range b.N {
// 		add(200, 300)
// 	}
// }

// func BenchmarkAddLargeInput(b *testing.B) {
// 	for range b.N {
// 		add(20000, 30000)
// 	}
// }

// sub-tests
//The need for sub-tests is important because it can tell which specific test case
//has caused abnormality. With below approach (table driven test), which testcase
//actually caused the abnormality can't be detected directly.
// func TestAddSubtests(t *testing.T) {
// 	testCases := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 1},
// 	}

// 	for _, test := range testCases {
// 		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
// 			result := add(test.a, test.b)
// 			if result != test.expected {
// 				t.Errorf("result = %d; Want = %d", result, test.expected)
// 			}
// 		})
// 	}
// }

// func TestAddTableDriven(t *testing.T) {
// 	testCases := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 1},
// 	}

// 	for _, test := range testCases {
// 		ans := add(test.a, test.b)
// 		if ans != test.expected {
// 			t.Errorf("Add(%d, %d) = %d, Want = %d", test.a, test.b, ans, test.expected)
// 		}
// 	}
// }

// func TestAdd(t *testing.T) {
// 	result := add(2, 5)
// 	expected := 9
// 	if result != expected {
// 		t.Errorf("result = %d, expected = %d", result, expected)
// 	}
// }
