package piscine

func IsAlpha(s string) bool {
	runeString := []rune(s)

	for _, letters := range runeString {
		if (letters < 48) || (letters > 57 && letters < 65) || (letters > 90 && letters < 97) || (letters > 122) {
			return false
		}
		if (letters > 47 && letters < 58) || (letters > 64 && letters < 91) || (letters > 96 && letters < 123) {
		}

	}
	return true
}
