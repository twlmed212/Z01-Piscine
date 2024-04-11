package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	if len(s) >= n {
		v := []rune(s)
		return v[n-1]
	}
	return 0
}
