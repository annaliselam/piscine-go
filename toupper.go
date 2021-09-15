package piscine

func ToUpper(s string) string {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		if runeString[i] >= 97 && runeString[i] <= 122 {
			runeString[i] = runeString[i] - 32
		}
	}
	return string(runeString)
}
