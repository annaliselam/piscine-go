package piscine

func NRune(s string, n int) rune {
	if n < 0 {
		return 0
	} 
	
	if n > len(s) {
		return 0
	}
	
	runeString := []rune(s)
	nChar := runeString[n-1]


	return rune(nChar)
}
