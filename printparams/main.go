package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ab := os.Args

	for i := 1; i < len(os.Args); i++ {
		for _, x := range ab[i] {
			z01.PrintRune(rune(x))
		}
		z01.PrintRune('\n')
	}
}
