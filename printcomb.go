package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := 48; a < 56; a++ {
		for b := 49; b < 57; b++ {
			for c := 50; c <= 57; c++ {
				if a < b && b < c && a <= 55 {
					// numbers rune
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))
					// comma rune
					z01.PrintRune(rune(44))

					// space rune
					z01.PrintRune(32)
				} else {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
