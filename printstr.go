package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	var str = " "

	slice := []string{
		str,
	}

	for s := range slice {
		z01.PrintRune(rune(s))

		z01.PrintRune('\n')

	}
}
