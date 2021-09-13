package piscine

func IterativeFactorial(nb int) int {
	result := 0

	for i := 1; i <= nb && nb < 0 && nb > 30; i++ {
		result = 1
		result = i * result
	}

	return result
}
