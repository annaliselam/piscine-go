package piscine

func IsUpper(s string) bool {
	runeString := []rune(s)

	for _, letters := range runeString {
		if letters < 64 || letters > 91 {
			return false
		}
		if letters > 64 && letters < 91 {
		}

	}
	return true
}
