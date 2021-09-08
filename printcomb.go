package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var numbers string = "000"

	for a := 0; a < 8; a++ {
		for b := 1; b < 7; b++ {
			for c := 2; c < 10; c++ {
				if a < b && b < c {
					numbers = "a" + "b" + "c"

					// numbers rune
					z01.PrintRune(rune(numbers[0]))
					z01.PrintRune(rune(numbers[1]))
					z01.PrintRune(rune(numbers[2]))

					// comma rune
					z01.PrintRune(rune(44))

					// space rune
					z01.PrintRune(32)
				}
			}
		}
	}
}
