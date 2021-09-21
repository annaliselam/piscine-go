package piscine

// Write a function IsSorted that returns true, if the slice of int is sorted, and that returns false otherwise.

// The function passed in parameter returns a positive int if a (the first argument) is superior to b (the second argument), it returns 0 if they are equal and it returns a negative int otherwise.

// To do your testing you have to write your own f function.

func IsSorted(f func(a, b int) int, a []int) bool {
	sortedSlice := true
	for i := 0; i <= len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				sortedSlice = false
				break
			}
		}
	}
	if sortedSlice {
		return true
	} else {
		return false
	}
}

func IsBigger(a int, b int) int {
	if a > b {
		return 1
	} else if b > a {
		return -1
	} else {
		return 0
	}
}
