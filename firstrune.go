package piscine

// Write a function that returns the first rune of a string.

func FirstRune(s string) rune {
	runeString := []rune(s)
	firstChar := runeString[0]

	return rune(firstChar)
}
