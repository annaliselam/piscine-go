package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombTwo() {
	for a := 48; a <= 57; a++ {
		for b := 48; b <= 57; b++ {
			for c := 48; c <= 57; c++ {
				for d := 48; d <= 57; d++ {
					// numbers rune
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					// space rune
					z01.PrintRune(32)
					// numbers rune
					z01.PrintRune(rune(c))
					z01.PrintRune(rune(d))

					if a == 57 && b == 57 && c == 57 && d == 57 {
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
}
