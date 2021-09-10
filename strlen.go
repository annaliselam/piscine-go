package piscine

import "github.com/01-edu/z01"

func StrLen(s string) int {
	// Write a function that counts the runes of a string and that returns that count.

	aStringChangeable := []byte(s)

	length := len(aStringChangeable)
	z01.PrintRune(rune(length))

	/* for i := 0; i < len(aStringChangeable); i++ {
		z01.PrintRune(rune(aStringChangeable[i]))
	} */
}
