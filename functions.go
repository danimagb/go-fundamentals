package main

import (
	"fmt"
	"strings"
)

func main() {

	coach := "sergio conceição"
	team := "fc.Porto"

	fmt.Println(convert(coach, team))
	fmt.Println(bestLigaNosFinishes(1, 1, 2, 3, 1, 2, 3))

	// Anonymous functions
	funcToExecute := func(c string) {
		fmt.Println(c)
	}
	funcToExecute(coach)

	//Self executed Anonymous function
	func(c string) {
		fmt.Println(c)
	}(coach)

	//Passing Anonymous functions as an argument into other function
	convertionFunc := func(coach, team string) string {
		return "MR. " + strings.Title(coach) + " from " + strings.ToUpper(team)
	}
	ConvertWith(convertionFunc, coach, team)

	//Returning a function from another function
	convertionfunc2 := GetConvertionFunction()
	fmt.Println(convertionfunc2(coach, team))
}

func convert(s1, s2 string) (str1, str2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)

	return s1, s2
}

//variadic function, when we don't know how many parameters will the function will receive.
//The input parameter is stored as a Slice of hints
//Same as 'params' keyword in c#
func bestLigaNosFinishes(finishes ...int) int {

	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}

func ConvertWith(i func(s, p string) string, s1, s2 string) {
	fmt.Println(i(s1, s2))
}

//Function which returns a function
func GetConvertionFunction() func(i, j string) string {
	myFunc := func(i, j string) string {
		return "Miss. " + strings.Title(i) + " from " + strings.ToUpper(j)
	}
	return myFunc
}
