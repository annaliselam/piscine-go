package main

// Write a program that prints the name of the program

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	a := os.Args[0][2:]

	for _, char := range a {

		z01.PrintRune(rune(char))

	}
	
	}
