package piscine

func Split(s, sep string) []string {

	mySlice := []string{}
	myString := ""

	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			mySlice = append(mySlice, myString)
			myString = ""
			i = i + len(sep)
		}
		if i > len(s)-len(sep)-2 {
			myString += string(s[i:])
			mySlice = append(mySlice, myString)
			myString = ""

		}
		myString += string(s[i])

	}
	return mySlice
}

/*
for i := 0; i < len(s)-len(sep); i++ {

		if s[i:i+len(sep)] == sep {
			mySlice = append(mySlice, myString)
			myString = ""
			i = i + len(sep)
		}
		if i > len(s)-len(sep)-2 {
			myString += string(s[i:])
			mySlice = append(mySlice, myString)
		}
		myString += string(s[i])
	}*/
