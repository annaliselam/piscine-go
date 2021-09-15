package piscine

func Capitalize(s string) string {
	runeString := []rune(s)

	for i := 0; i < len(s); i++ {
		if runeString[i] == 32 || runeString[i] == '+' {
			if runeString[i+1] >= 'a' && runeString[i+1] <= 'z' {
				runeString[i+1] = runeString[i+1] - 32
			}
		}
	}

	return string(runeString)
}
