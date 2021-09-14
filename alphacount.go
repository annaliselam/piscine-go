package piscine

func AlphaCount(s string) int {
	var counter int

	runeString := []rune(s)

	// if i < 65 && i < 90 && i < 97 && i < 122 {
	for _, letters := range runeString {
		if (letters > 65 && letters < 90) || (letters > 97 && letters < 122) {
			counter++
		}
	}
	return counter
}
