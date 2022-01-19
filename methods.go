package main

import (
	"fmt"
)

type team struct {
	name         string
	titles       int
	abbreviation string
}

type data int

func main() {

	// Methods contains a receiver argument in it, the method can access the properties of the receiver.
	// The receiver can be a struct or non-struct type
	// Similar to extension methods in c#

	bestTeam := team{"Futebol Clube do Porto", 31, "FCP"}

	//Method with struct type receiver
	bestTeam.show()

	//Method with non-struct type receiver
	value1 := data(20)
	value2 := data(10)
	result := value1.multiply(value2)
	fmt.Println("Result: ", result)

	//Method with pointer receiver
	bestTeamPointer := &bestTeam
	bestTeamPointer.incrementTitles()
	bestTeam.show()
}

func (t team) show() {
	fmt.Println("Name: ", t.name)
	fmt.Println("Abbreviation: ", t.abbreviation)
	fmt.Println("Titles: ", t.titles)
}

func (i data) multiply(multiplier data) data {
	return i * multiplier
}

func (t *team) incrementTitles() {
	(*t).titles += 1
}
