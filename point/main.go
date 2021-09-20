package main

import (
	"github.com/01-edu/z01"
)

type xy struct {
	name  string
	value int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	x := xy{}
	y := xy{}

	x.name = "x"
	x.value = 42

	y.name = "y"
	y.value = 21

	printStr("x = 42, y =21")
}

// test go run .
// output x = 42, y = 21
