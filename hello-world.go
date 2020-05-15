package main

import (
	"fmt"
	"os"
)

var (
	name   = "Daniel"
	course = "GoLang"
	module = 2.1
)

func main() {

	name = os.Getenv("USERNAME")
	// ptr := &module

	//fmt.Println("Hello from", runtime.GOOS)

	// fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	// fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	// fmt.Println("NameTwo is", nameTwo, "and is of type", reflect.TypeOf(nameTwo))

	// fmt.Println("Memory address of *module* variable is ", ptr, "and the value of *module* is ", *ptr)

	fmt.Println("\nHi", name, "you're currently watching course", course)

	changeCourse(&course)

	fmt.Println("\nHi", name, "you are now watching course", course)
}

func changeCourse(course *string) string {

	*course = "Native docker clustering"

	fmt.Println("\nTrying to change course to", *course)

	return *course
}
