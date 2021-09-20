package piscine

func BasicAtoi (s string) int {
	result := 0
	for _, num := range s {
		convert := int(num) - 48
		result = (result * 10) + convert

	}
	return result
}
