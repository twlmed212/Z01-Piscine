package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	firstArgs := os.Args[1]
	secondArgs := os.Args[2]

	uniqueString := ""
	for _, check := range firstArgs {
		isUnique := true
		for _, checkString := range uniqueString {
			if check == checkString {
				isUnique = false
			}
		}
		if isUnique {
			uniqueString += string(check)
		}
	}

	for _, check := range secondArgs {
		isUnique := true
		for _, checkString := range uniqueString {
			if check == checkString {
				isUnique = false
			}
		}
		if isUnique {
			uniqueString += string(check)
		}
	}
	for _, getChar := range uniqueString {
		z01.PrintRune(getChar)
	}
	z01.PrintRune('\n')

	// if len(os.Args) != 3 {
	// 	z01.PrintRune('\n')
	// 	return
	// } else {
	// 	fString := os.Args[1]
	// 	sString := os.Args[2]
	// 	// zpadinton paqefwtdjetyiytjneytjoeyjnejeyj |
	// 	myString := ""
	// 	for _, getChar := range fString {
	// 		isUnique := true
	// 		for _, toCompare := range myString {
	// 			if toCompare == getChar {
	// 				isUnique = false
	// 			}
	// 		}
	// 		if isUnique {
	// 			myString += string(getChar)
	// 		}
	// 	}
	// 	for _, getChar := range sString {
	// 		isUnique := true
	// 		for _, toCompare := range myString {
	// 			if toCompare == getChar {
	// 				isUnique = false
	// 			}
	// 		}
	// 		if isUnique {
	// 			myString += string(getChar)
	// 		}
	// 	}
	// 	for _, char := range myString {
	// 		z01.PrintRune(char)
	// 	}
	// 	z01.PrintRune('\n')

	// }

}
