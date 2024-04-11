package piscine

func Index(s string, toFind string) int {
	lenS := len(s)
	lenFind := len(toFind)
	for i := range s {
		if i+lenFind < lenS && s[i:i+lenFind] == toFind {
			return i
		}
	}
	return -1
}
