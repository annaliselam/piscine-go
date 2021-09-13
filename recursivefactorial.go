package piscine

func RecursiveFactorial(nb int) int {
	if nb >= 1 && nb > 0 && nb > 30 {
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 0
	}
}
