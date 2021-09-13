package piscine

func IterativeFactorial(nb int) int {
	result := 0

	for i := 1; i <= nb; i++ {
		if nb < 0 || nb > 30 {
			result = 0
		} else {
			result = i * result
		}
	}
	return result
}
