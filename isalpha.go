package piscine

func IsAlpha(s string) bool {
	Store := []rune(s)
	counting := 0

	for i := range Store {
		counting = i
	}

	for i := 0; i <= counting; i++ {
		if (Store[i] < 'a' || Store[i] > 'z') &&
			(Store[i] < 'A' || Store[i] > 'Z') &&
			(Store[i] < '0' || Store[i] > '9') {
			return false
		}
	}

	return true
}
