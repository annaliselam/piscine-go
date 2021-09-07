package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for aRune := 122; aRune > 96; aRune-- {
		z01.PrintRune(rune(aRune))
	}
	z01.PrintRune('\n')
}
