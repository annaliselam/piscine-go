package piscine

func IsNumeric(s string) bool {
	runeString := []rune(s)

	for _, letters := range runeString {
		if (letters < 48) || (letters > 57) {
			return false
		}
		if letters > 47 && letters < 58 {
		}

	}
	return true
}
