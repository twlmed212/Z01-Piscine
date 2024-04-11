package piscine

func Swap(a *int, b *int) {
	aSwap := *a
	bSwap := *b
	*a = bSwap
	*b = aSwap
}
