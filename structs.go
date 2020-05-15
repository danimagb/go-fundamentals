package main

import (
	"fmt"
)

func main() {
	//Structs are custom data types (OO programming)
	//Go doesn't have Objects, Classes and inheritance

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println("\nDocker deep dive author is:", DockerDeepDive.Author)
	fmt.Println("\nDocker deep dive rating is:", DockerDeepDive.Rating)
}
