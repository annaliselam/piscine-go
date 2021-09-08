package piscine

func DivMod(a int, b int, div *int, mod *int) {
	var div *int = a / b
	var mod *int = a % b

	print(&div)
	print(&mod)
}