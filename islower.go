package piscine

func IsLower(s string) bool {
	runeString := []rune(s)

	for _, letters := range runeString {
		if letters < 97 || letters > 122 {
			return false
		}
		if letters > 96 && letters < 123 {
		}

	}
	return true
}
