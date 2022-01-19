package main

import (
	"fmt"
	"time"
)

func main() {

	//for { // Infinite loop
	//
	//	}

	fmt.Println("## For-range loop to iterate over an array ##")
	courseList := []string{"docker", "kubernetes"}
	for i, course := range courseList {
		fmt.Println(course)
		fmt.Println(i)
	}

	fmt.Println("## For loop with pre and post statements ##")
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

}
