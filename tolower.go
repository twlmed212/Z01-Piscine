package piscine

func ToLower(s string) string {
	runeIt := []rune(s)

	for i := 0; i < len(s); i++ {
		if runeIt[i] >= 'A' && runeIt[i] <= 'Z' {
			for j := 'a'; j <= 'z'; j++ {
				if runeIt[i]+32 == j {
					runeIt[i] = j
				}
			}
		}
	}
	return string(runeIt)
}
