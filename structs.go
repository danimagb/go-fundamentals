package main

import (
	"fmt"
)

type courseMeta struct {
	Author string
	Level  string
	Rating float64
}

func main() {
	//Structs are custom data types
	//Go doesn't have Objects, Classes and inheritance

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	dockerDeepDive := courseMeta{
		Author: "Some Person",
		Level:  "Intermediate",
		Rating: 5,
	}

	dockerDeepDivePointer := &courseMeta{
		Author: "Some Person 2",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(dockerDeepDive)
	fmt.Println("\nDocker deep dive author is:", dockerDeepDive.Author)
	fmt.Println("\nDocker deep dive rating is:", dockerDeepDive.Rating)

	//Modifying by value, nothing happens because an entirely new copy of the struct is created and set to the function
	modifyValue(dockerDeepDive)
	fmt.Println("\nDocker deep dive author is:", dockerDeepDive.Author)

	//Modifying by pointer, it actually changes the Author because we are passing a pointer to the reference
	modifyPointer(&dockerDeepDive)
	fmt.Println("\nDocker deep dive author is:", dockerDeepDive.Author)

	fmt.Println("\n")
	fmt.Println((*dockerDeepDivePointer))
	fmt.Println("\nDocker deep dive author is:", (*dockerDeepDivePointer).Author)
	fmt.Println("\nDocker deep dive rating is:", (*dockerDeepDivePointer).Rating)

	//Anonymous structure, useful when you want to create a one-time usable structure
	element := struct {
		name string
		from string
	}{
		name: "Daniel",
		from: "Portugal",
	}

	fmt.Println(element)

}

func modifyValue(c courseMeta) {
	c.Author = "Modified " + c.Author
}

func modifyPointer(c *courseMeta) {
	c.Author = "Modified " + c.Author
}
