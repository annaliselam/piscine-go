package piscine

func ToLower(s string) string {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		if runeString[i] >= 65 && runeString[i] <= 90 {
			runeString[i] = runeString[i] + 32
		}
	}
	return string(runeString)
}
