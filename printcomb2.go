package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := 48; a <= 57; a++ {
		for b := 48; b <= 57; b++ {
			for c := 48; c <= 57; c++ {
				for d := 49; d <= 57; d++ {
					if c > a || d > b && c == a {
						if a == 9 && b == 8 && c == 9 && d == 9 {
							// numbers rune
							z01.PrintRune(rune(a))
							z01.PrintRune(rune(b))
							// space rune
							z01.PrintRune(32)
							// numbers rune
							z01.PrintRune(rune(c))
							z01.PrintRune(rune(d))
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
}
