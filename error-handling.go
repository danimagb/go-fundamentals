package main

import (
	"fmt"
	"os"
)

func main() {
	// Good Practice to allways return an error on every function
	// e.g. func (target string) (rspTime float64 err error)
	// We should check if err is != nil
	if _, err := os.Open("c:\\temp\\tt.txt"); err != nil {
		fmt.Println("Error returned was:", err)
	}

}
