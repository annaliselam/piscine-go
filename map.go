package piscine

func Map(f func(int) bool, a []int) []bool {
	for _, elem := range a {
		f(elem)
	}
	answer := make([]bool, len(a))

	return answer
}
