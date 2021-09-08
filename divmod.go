package piscine

func DivMod(a int, b int, div *int, mod *int) {
	var division int = a / b
	var modulo int = a % b

	division = *div
	modulo = *mod
}
