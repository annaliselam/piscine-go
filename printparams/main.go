package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ab := os.Args

	a := ab[0]

	for _, char := range a {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}
