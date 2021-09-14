package piscine

func Index(s string, toFind string) int {
	// the original string is converted to a slice rune string
	// find is the length of the substring
	runeString := []rune(s)
	find := len([]rune(toFind))

	// if the length of the original string is less than the substring, return -1
	if len(runeString) < find {
		return -1
	}

	// for loop to search the original string
	for i := 0; i < len(runeString); i++ {
		// if the length of the split substring is less than the substring, return -1
		if len(runeString[i:]) < len([]rune(toFind)) {
			return -1
		} else {
			// if the substring matches the original string, it will return the index
			// need to have another look at this
			if s[i:i+find] == toFind {
				return i
			}
		}
	}
	return -1
}
