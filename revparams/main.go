package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ab := os.Args

	for i := len(os.Args) - 1; i > 0; i-- {
		for _, x := range ab[i] {
			z01.PrintRune(rune(x))
		}
		z01.PrintRune('\n')
	}
}
