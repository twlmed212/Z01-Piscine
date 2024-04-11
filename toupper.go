package piscine

func ToUpper(s string) string {
	runeIt := []rune(s)

	for i := 0; i < len(s); i++ {
		if runeIt[i] >= 'a' && runeIt[i] <= 'z' {
			for j := 'A'; j <= 'Z'; j++ {
				if runeIt[i]-32 == j {
					runeIt[i] = j
				}
			}
		}
	}
	cnvRunToString := string(runeIt)
	return cnvRunToString
}
