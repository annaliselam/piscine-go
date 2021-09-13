package piscine

func IterativePower(nb int, power int) int {
	result := 1

	if nb < 0 {
		result = 0
	} else if power == 0 {
		result = 1
	} else {
		for i := 1; i < nb; i++ {
			result *= nb
			power -= 1
		}
	}

	return result
}
