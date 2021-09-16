package main

// Write a program that prints the name of the program

import (
	"os"
	"github.com/01-edu/z01"
)

func main() { 
	arguments := os.Args
	argName := len(arguments)

	for i := 1; i < len(arguments); i++ {
		z01.PrintRune(rune(argName))
	}
	
	z01.PrintRune('\n')
}
