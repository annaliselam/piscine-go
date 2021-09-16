package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ab := os.Args

	for i := 0; i < len(os.Args); i++ {
		a := ab[i]
		for _, char := range a {
			z01.PrintRune(rune(char))
		}
	}

	z01.PrintRune('\n')
}
