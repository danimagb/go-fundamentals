package main

import (
	"fmt"
)

func main() {

	//Slices are based on arrays
	//Slices are passed to functions by reference

	//creates a slice with initial lenght of 10 and a capacity of 100(creates an array with 100)
	//myCourses := make([]string, 10, 100)

	//Creates a slice with initial len 6 and capacity of 6
	//myIntegers := []int{1, 2, 3, 4, 5, 6}

	//Creates a subslice of myIntegers from index 2(inclusive) to index 5(exclusive)
	//sliceOfSlice1 := myIntegers[2:5]

	//Creates a subslice of myIntegers from index 0(inclusive) to index 5(exclusive)
	//sliceOfSlic2 := myIntegers[:5]

	//Creates a subslice of myIntegers from index 0(inclusive) to array len(exclusive)
	//sliceOfSlice3 := myIntegers[2:]

	// When capacity is reached the array is doubled sized e.g. -> 4 -> 8 ->16 ->32
	mySlice := make([]int, 1, 4)

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", cap(mySlice))
	}

	//Appending two slices of the same type(we're appending newSlice integers to mySlice)
	newSlice := []int{100, 200, 300}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)
}
