package piscine

func IsUpper(s string) bool {
	Store := []rune(s)
	counting := 0

	for i := range Store {
		counting = i
	}
	for i := 0; i <= counting; i++ {
		if Store[i] < 'A' || Store[i] > 'Z' {
			return false
		}
	}

	return true
}
