package piscine

func BasicJoin(elems []string) string {
	emptyString := ""

	for _, char := range elems {
		emptyString += char
	}
	return emptyString
}
