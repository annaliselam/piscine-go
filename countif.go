package piscine

func CountIf(f func(string) bool, tab []string) int {
	var counter int
	for _, value := range tab {
		if f(value) {
			counter++
		}
	}
	return counter
}
