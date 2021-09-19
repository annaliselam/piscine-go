package piscine

// Write a function that takes an int min and an int max as parameters. That function returns a slice of ints with all the values between min and max.
// Min is included, and max is excluded.
// If min is superior or equal to max, a nil slice is returned.
// make is not allowed for this exercise.

func AppendRange(min, max int) []int {
	var answer []int
	var size int
	
	if min > max {
		size = 0
		
	} else {
		size = max - min
	}
	
	for i := 0; i < size; i++ {
			answer = append(answer, ((max-min)) + 1)
			for i := range answer {
				answer[i] = min + i
			}
		}

	return answer
}
