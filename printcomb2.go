package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombTwo() {
	for a := 48; a <= 57; a++ {
		for b := 48; b <= 57; b++ {
					// numbers rune
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					if a == 57 && b == 57 {
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
