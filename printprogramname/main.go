package main

// Write a program that prints the name of the program

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ab := os.Args

	a := ab[0][2:]

	for _, char := range a {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}
