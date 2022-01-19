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

	// Switch has implicit breaks, we can use fallthrough to enter next matches
	switch "docker" {
	case "docker":
		fmt.Println("linux")
	case "kubernetes":
		fmt.Println("docker")
		fallthrough
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("no match")
	}

	//Good practice, match multiple values in a single case instead of using 'fallthrough'
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got an even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got an odd number -", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
