package piscine

func IterativeFactorial(nb int) int {
	result := 0
	if nb < 0 || nb > 30 {
	} else {
		for i := 1; i <= nb; i++ {
			result = i * result
		}
	}
	return result
}
