package piscine

func RecursiveFactorial(nb int) int {
	if nb >= 1 && nb < 30 {
		nb = nb * RecursiveFactorial(nb-1)
	} else {
		nb = 0
	}
	return nb
}
