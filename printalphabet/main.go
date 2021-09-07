package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for aRune := 97; aRune < 123; aRune++ {
		z01.PrintRune(rune(aRune))
	}
	z01.PrintRune('\n')
}
