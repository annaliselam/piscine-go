package piscine

func StrLen(s string) int {
	// Write a function that counts the runes of a string and that returns that count.
	aStringChangeable := []rune(s)
	length := len(aStringChangeable)
	return length
}
