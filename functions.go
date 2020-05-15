package main

import (
	"fmt"
)

func main() {

	bestFinish := bestLigaNosFinishes(1, 1, 2, 3, 1, 2, 3)

	fmt.Println(bestFinish)
}

func bestLigaNosFinishes(finishes ...int) int {

	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
