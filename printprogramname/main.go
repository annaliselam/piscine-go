package main

// Write a program that prints the name of the program

import (
	"os"
	
)

func main() {
	arguments := os.Args

	for i := 0; i <= len(arguments); i++ {
		
		println(os.Args)
	}
}
