package piscine

func AtoiBase(s string, base string) int {
	res := 0
	if isBaseValid(base) {
		for _, char := range s {
			index := getIndex(byte(char), base)
			res = res*len(base) + index
		}
	}
	return res
}
func getIndex(s byte, base string) int {
	for i, char := range base {
		if char == rune(s) {
			return i
		}
	}
	return -1
}

func isBaseValid(base string) bool {
	if len(base) < 2 {
		return false
	}
	myMap := map[rune]bool{}
	for _, char := range base {
		if myMap[char] || char == '+' || char == '-' {
			return false
		}
		myMap[char] = true
	}
	return true
}

/*

func AtoiBase(s string, base string) int {
	result := 0
	if baseIsValid(base) {
		for _, check := range s {
			index := getIndex(byte(check), base)
			result = result*len(base) + index
		}
	}
	return result
}

func getIndex(char byte, base string) int {
	for i, base := range base {
		if char == byte(base) {
			return i
		}
	}
	return -1
}

func baseIsValid(base string) bool {
	if len(base) < 2 {
		return false
	}
	baseChecker := map[rune]bool{}
	for _, char := range base {
		if baseChecker[char] || char == '-' || char == '+' {
			return false
		}
		baseChecker[char] = true
	}
	return true
}
*/
