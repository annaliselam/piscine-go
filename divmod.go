package piscine

func DivMod(a int, b int, div *int, mod *int) {
	var *div = a / b
	var *mod = a % b

	print(&div)
	print(&divmod)
}