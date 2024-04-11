package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\b' || s[i] == '\f' || s[i] == '\r' {
			return false
		}
	}

	return true
}
