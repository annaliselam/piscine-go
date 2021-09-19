package piscine

func SplitWhiteSpaces(str string) []string {
	ln := 0
	spaces := false
	for c := range str {
		if spaces && c != 0 && (str[c-1] == '\n' ||
			str[c-1] == '\t' || str[c-1] == ' ') &&
			str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			ln++
		}
		if str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			spaces = true
		}
	}
	ln++

	count := 0
	ans := make([]string, ln)
	ok := true
	for _, c := range str {
		if c == '\n' || c == '\t' || c == ' ' {
			if !ok {
				count++
			}
			ok = true
			continue
		}
		ans[count] = ans[count] + string(c)
		ok = false
	}
	return ans
}
