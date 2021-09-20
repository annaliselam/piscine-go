package main

import (
	"os"

	"github.com/01-edu/z01"
)

// use os.args to count the number of arguments
// check the number of arguments, then return an error message

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(numArgs int) bool {
	if (numArgs % 2) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:]
	var numArgs int

	for i := 1; i < len(args); i++ {
		numArgs += i
	}

	if isEven(numArgs) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
