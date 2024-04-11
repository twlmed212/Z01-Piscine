package piscine

func ConcatParams(mySlice []string) string {
	myString := ""
	for i := range mySlice {
		myString += mySlice[i]
		myString += "\n"
	}
	return myString
}

/*
func ConcatParams(args []string) string {
	getString := ""
	for i, v := range args {
		if i == len(args)-1 {
			getString += v
		} else {
			getString += v + "\n"
		}
	}
	return getString
}
*/
