package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for aRune := 48; aRune < 58; aRune++ {
		z01.PrintRune(rune(aRune))
	}
	z01.PrintRune('\n')
}
