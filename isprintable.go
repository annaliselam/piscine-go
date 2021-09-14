package piscine

func IsPrintable(s string) bool {
	runeString := []rune(s)

	for _, letters := range runeString {
		if letters < 32 || letters > 126 {
			return false
		}
		if letters > 32 && letters < 127 {
		}

	}
	return true
}
