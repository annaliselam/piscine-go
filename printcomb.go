package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := 0; a < 8; a++ {
		for b := 1; b < 9; b++ {
			for c := 2; c < 10; c++ {
				if a < b && b < c {

					// numbers rune
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))

					// comma rune
					z01.PrintRune(rune(44))

					// space rune
					z01.PrintRune(32)
				}
			}
		}
	}
}
