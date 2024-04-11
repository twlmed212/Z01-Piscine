package piscine

func AlphaCount(s string) int {
	cnt := 0
	for _, counting := range s {
		if counting >= 'a' && counting <= 'z' || counting >= 'A' && counting <= 'Z' {
			cnt++
		}
	}
	return cnt
}
