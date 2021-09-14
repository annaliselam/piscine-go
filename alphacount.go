package piscine

func AlphaCount(s string) int {
	var counter int

	runeString := []rune(s)

	for _, letters := range runeString {
		if (letters > 64 && letters < 91) || (letters > 96 && letters < 123) {
			counter++
		}
	}
	return counter
}
