package piscine

func Swap(a *int, b *int) {
	swapA := *a
	swapB := *b

	*a = swapB
	*b = swapA

}
