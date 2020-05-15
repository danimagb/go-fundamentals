package main

import (
	"fmt"
)

func main() {
	firstRank := "39"
	secondRank := "614"

	// Simple If
	if firstRank < secondRank {
		fmt.Println("\n First course is doing better than the second course")
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear... your first course must be doing abysmally!!")
	} else {
		fmt.Println("\nBoth courses are the same rank")
	}

	// If with statements initialization (variables are defined within the if scope and are garbage collected at the end)
	if rankOne, rank2 := 39, 614; rankOne < rank2 {
		fmt.Println("\n First course is doing better than the second course")
	} else if rankOne > rank2 {
		fmt.Println("\nOh dear... your first course must be doing abysmally!!")
	} else {
		fmt.Println("\nBoth courses are the same rank")
	}
}
