package piscine

func MakeRange(min, max int) []int {
	var size int

	if min >= max {
		size = 0
	} else if min < 0 || max < 0 {
		size = 0
	} else {
		size = max - min
	}

	answer := make([]int, size)

	for i := 0; i < size; i++ {
		answer[i] = (max - min) + 1
		for i := range answer {
			answer[i] = min + i
		}
	}

	return answer
}
