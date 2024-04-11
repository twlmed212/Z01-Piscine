package piscine

func Join(strs []string, sep string) string {
	var myString string
	for _, getChar := range strs {
		if getChar != strs[len(strs)-1] {
			myString += string(getChar) + sep
		} else {
			myString += string(getChar)
		}
	}

	return myString
}
