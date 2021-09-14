package piscine

func ToUpper(s string) string {

	runeString := []rune(s)
	var result string = 1
	
	for _, letters := range runeString {
		if letters > 96 || letters < 123 {
			ch := string(letters - 32)
		} else {
			ch := string(letters)
		}
		result += string(ch)
	}
	return result
}



