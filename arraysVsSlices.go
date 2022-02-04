package main

import (
	"fmt"
	"sort"
)

func main() {

	//Arrays have a fixed limit, they are not resizable
	// Arrays ar of value type and not of reference type
	arr := [4]string{"docker", "kuberneted", "terraform"}

	//Declaring an array with the ellipsis, the array lenght is determined by the number of initialized elements
	// E.g arr := [...]string{"docker", "kuberneted"} -> len(arr) = 2

	fmt.Println("\nTraversing array")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//Multi dimensional arrays
	array2d := [3][2]string{
		{"C#", "C"},
		{"GO", "Python"},
		{"C++", "TypeScript"},
	}

	fmt.Println("\nTraversing 2d array")

	for i := 0; i < len(array2d); i++ {
		for j := 0; j < len(array2d[0]); j++ {
			fmt.Println(array2d[i][j])
		}
	}

	fmt.Println("\nTraversing 2d array with a for range loop")

	for _, row := range array2d {
		for _, val := range row {
			fmt.Println(val)
		}
	}

	//Slices are based on arrays, but they are resizable
	//Slices are passed to functions by reference

	//creates a slice with initial lenght of 10 and a capacity of 100(creates an array with 100)
	//myCourses := make([]string, 10, 100)
	//myCourses[0] = "Docker"
	//myCourses[1] = "Kubernetes"

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

	fmt.Println("\nIncreesing the capacity of a slice")
	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d", cap(mySlice))
	}

	//Appending two slices of the same type(we're appending newSlice integers to mySlice)
	fmt.Println("\nAppending two slices")
	newSlice := []int{100, 200, 300}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

	// Creating multi-dimensional slice
	s1 := [][]int{
		{12, 34},
		{56, 47},
		{29, 40},
		{46, 78},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 1 : ", s1)

	// Creating multi-dimensional slice
	s2 := [][]string{
		{"Docker", "K8S"},
		{"Terraform", "AWS"},
		{"GCP", "Azure"},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 2 : ", s2)

	//Sorting slices in ascending order
	scl1 := []int{400, 600, 100, 300, 500, 200, 900}

	fmt.Println("Is slice sorted? ", sort.IntsAreSorted(scl1))
	sort.Ints(scl1)
	fmt.Println("Is slice sorted? ", sort.IntsAreSorted(scl1))

}
