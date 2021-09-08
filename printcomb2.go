package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					if c > a || d > b && c == a {
						if a == '9' && b == '8' && c == '9' && d == '9' {
							z01.PrintRune(rune(a))
							z01.PrintRune(rune(b))
							// space rune
							z01.PrintRune(32)
							// numbers rune
							z01.PrintRune(rune(c))
							z01.PrintRune(rune(d))
							z01.PrintRune('\n')
						} else {
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
						}
					}
				}
			}
		}
	}
}
