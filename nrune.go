package piscine

func NRune(s string, n int) rune {
	if n < 0 || n >= len(s) {
		return 0
	} else {
	}
	runeString := []rune(s)
	nChar := runeString[n]

	return rune(nChar)
}
