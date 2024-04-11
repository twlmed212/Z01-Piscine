package piscine

func UltimateDivMod(a *int, b *int) {
	store := *a
	*a = store / *b
	*b = store % *b
}
