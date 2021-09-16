package main

// Write a program that prints the name of the program

import (
	"os"
	// "github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for i := 0; i <= len(arguments); i++ {
		println(os.Args)
		// z01.PrintRune(rune())
	}
}
