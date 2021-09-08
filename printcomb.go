package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := 48; a < 56; a++ {
	}

	for b := 49; b < 57; b++ {
	}

	for c := 50; c < 49+48; c++ {
	}

	var numbers rune = 'a' + 'b' + 'c'

	z01.PrintRune(rune(numbers))

	// comma rune
	z01.PrintRune(rune(44))

	// space rune
	z01.PrintRune(32)
}

func main() {
	piscine.PrintComb()
}
