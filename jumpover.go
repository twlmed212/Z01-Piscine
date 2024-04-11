package piscine

func JumpOver(str string) string {
	myBeLovedString := ""
	if len(str) <= 2 {
		return "\n"
	}

	for i := range str {
		i = i + 1
		if i%3 == 0 {
			myBeLovedString += string(str[i-1])
		}
	}

	return myBeLovedString + "\n"
}
