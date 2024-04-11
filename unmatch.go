package piscine

func Unmatch(a []int) int {
	count := 0
	for _, getNumbers := range a {
		for _, checkNumbers := range a {
			if getNumbers == checkNumbers {
				count++
			}
		}
		if count == 1 || (count%2 == 1) {
			return getNumbers
		}
		count = 0
	}
	return -1
}
