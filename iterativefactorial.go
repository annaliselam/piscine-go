package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb <= 0 || nb > 30 {
		result = 0
		return result
	} else {
		result := 1
		for i := 1; i <= nb; i++ {
			result *= i
			return result
		}
	}
	return result
}
