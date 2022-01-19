package main

import (
	"fmt"
)

func main() {

	// Defer statements delay the execution of the function/method until the nearby functions return.
	// We can have multiple defer statements and they are executed in LIFO order
	// Defer statements are generally used to ensure that the files are closed when their need is over, or to close the channel, or to catch the panics in the program.

	fmt.Println("Starting")
	multiply(2, 4) // 8

	defer multiply(8, 2) // 16

	fmt.Println("Next to be executed is the last defer multiply")

	// Multiple defer statements
	// Executes in LIFO order
	defer fmt.Println("End")
	defer multiply(9, 4) // 36
	defer multiply(8, 9) // 72

}

func multiply(a1, a2 int) {
	res := a1 * a2
	fmt.Println("Result: ", res)
}
