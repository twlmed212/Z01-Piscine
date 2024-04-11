package piscine

func Rot14(s string) string {
	myString := ""

	for _, char := range s {
		if (char >= 'a' && char <= 'm') || (char >= 'A' && char <= 'M') {
			myString += string(char + 14)
		} else if (char >= 'M' && char <= 'Z') || (char >= 'm' && char <= 'z') {
			myString += string(char - 12)
		} else {
			myString += string(char)

		}
	}
	return myString
}

//Vszzc! Vck ofs Mci?
