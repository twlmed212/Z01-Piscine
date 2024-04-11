package piscine

func Compare(a, b string) int {
	storeValuesA := []rune(a)
	storeValuesB := []rune(b)
	if storeValuesA[0] == storeValuesB[0] && len(storeValuesA) == len(storeValuesB) {
		return 0
	} else if storeValuesA[0] < storeValuesB[0] {
		return -1
	} else if storeValuesA[0] == storeValuesB[0] && len(storeValuesB) <= len(storeValuesA) {
		return 1
	} else if storeValuesA[0] < storeValuesB[0] && len(storeValuesA) >= len(storeValuesB) {
		return -1
	} else if storeValuesA[0] > storeValuesB[0] && len(storeValuesA) >= len(storeValuesA) {
		return 1
	}
	return 23
}
