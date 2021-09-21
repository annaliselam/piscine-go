package piscine

// Write a function ForEach that, for an int slice, applies a function on each elements of that slice.

func ForEach(f func(int), a []int) {

	for _, b := range a {
		a = f(b)
	}

	return a
}
