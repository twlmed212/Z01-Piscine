package piscine

import "strings"

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}

	myslice := strings.Fields(str)
	myString := ""
	//remove Spaces
	for i := len(myslice) - 1; i >= 0; i-- {
		if i != 0 {
			myString += myslice[i] + " "
		} else {
			myString += myslice[i]

		}
	}
	return myString + "\n"
}
