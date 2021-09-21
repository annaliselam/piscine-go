package piscine

func Map(f func(int) bool, a []int) []bool {
	
	returnSlice := make([]bool, len(a))

	for i, v := range a {
		returnSlice[i] = f(v)	
	}


	return returnSlice

}
