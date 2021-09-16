package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ab := os.Args

	// a := ab[1]

	for i := 1; i < len(os.Args); i++ {
		b := ab[i]
		for _, char := range b {
			z01.PrintRune(rune(char))
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
