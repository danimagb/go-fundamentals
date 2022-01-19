package main

import (
	"fmt"
	"reflect"
	"runtime"
	"os"
	"strconv"
)

var (
	name   = "Daniel"
	course = "GoLang"
	module = 2
	clip = "3"
)

func main() {

	name = os.Getenv("USERNAME")
	ptr := &module


	fmt.Println("Hello from", runtime.GOOS)

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))

	fmt.Println("Memory address of *module* variable is ", ptr, "and the value of *module* is ", *ptr)

	fmt.Println("\nHi", name, "you're currently watching course", course)

	changeCourse(&course)

	fmt.Println("\nHi", name, "you are now watching course", course)

	//Type Conversion string to Int
	integerClip, err := strconv.Atoi(clip)
	if err == nil{
		total := integerClip + module
		fmt.Println("Clip plus module equals", total)
	}
}

func changeCourse(course *string) string {

	*course = "Native docker clustering"

	fmt.Println("\nTrying to change course to", *course)

	return *course
}
