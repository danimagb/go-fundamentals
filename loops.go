package main

import (
	"fmt"
)

func main() {

	//for { // Infinite loop
	//
	//	}

	courseList := []string{"docker", "kubernetes"}
	for _, course := range courseList {
		fmt.Println(course)
	}

	for i := 0; i < 10; i++ {

	}

}
