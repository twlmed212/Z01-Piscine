package piscine

func IsLower(s string) bool {
	Store := []rune(s)
	counting := 0

	for i := range Store {
		counting = i
	}
	for i := 0; i <= counting; i++ {
		if Store[i] < 'a' || Store[i] > 'z' {
			return false
		}
	}

	return true
}
