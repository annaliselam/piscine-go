package piscine

func Capitalize(s string) string {
	str := []rune(s)
	for i, word := range str {
		if word >= 'A' && word <= 'Z' || word >= 'a' && word <= 'z' {
			if word >= 'A' && word <= 'Z' {
				str[i] = word + 32
			}
		}
	}
	for i, word := range str {
		if word >= 'a' && word <= 'z' {
			if i-1 < 0 {
				str[i] = word - 32
			} else if (str[i-1] >= 'a' && str[i-1] <= 'z') || (str[i-1] >= 'A' && str[i-1] <= 'Z') || (str[i-1] >= '0' && str[i-1] <= '9') {
			} else {
				str[i] = word - 32
			}
		}
	}
	return string(str)
}
