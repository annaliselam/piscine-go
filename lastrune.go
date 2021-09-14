package piscine

func LastRune(s string) rune {
	runeString := []rune(s)
	lastChar := runeString[len(s)-1]

	return rune(lastChar)
}
