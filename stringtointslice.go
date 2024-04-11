package piscine

func StringToIntSlice(str string) []int {
	var finalResult []int
	if str == "" {
		return finalResult
	}
	for _, get := range str {
		finalResult = append(finalResult, int(get))
	}
	return finalResult
}
