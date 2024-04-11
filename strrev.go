package piscine

func StrRev(s string) string {
	store := ""
	for i := len(s) - 1; i >= 0; i-- {
		store += string(s[i])
	}
	return store
}
