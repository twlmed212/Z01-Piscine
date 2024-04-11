package piscine

func CountAlpha(s string) int {
	count := 0
	for _, getChar := range s {
		if getChar >= 'a' && getChar <= 'z' || getChar >= 'A' && getChar <= 'Z' {
			count++
		}
	}
	return count
}
