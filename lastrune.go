package piscine

func LastRune(s string) rune {
	storeValue := []rune(s)
	return storeValue[len(storeValue)-1]
}
