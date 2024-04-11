package piscine

func Capitalize(s string) string {
	var myBelovedString string
	count := 0
	for _, getChar := range s {
		sliceString := false
		if (getChar >= '!' && getChar <= '/') || getChar == ' ' {
			sliceString = true
			myBelovedString += string(getChar)
			count = 0

		}
		if !sliceString {
			if count == 0 && getChar >= 'a' && getChar <= 'z' {
				temp := getChar - 32
				myBelovedString += string(temp)
				count = 1

			} else {
				myBelovedString += string(getChar)
				count++
			}

		}

	}

	return myBelovedString
}

func BasicAtoi(s string) int {
	var myInt int
	for _, getChar := range s {
		if getChar >= '0' && getChar <= '9' {
			myInt *= 10
			// fmt.Println("before *", myInt)
			myInt += int(getChar - '0')
			// fmt.Println("after +=", myInt)

		}
	}
	return myInt
}

// func BasiItoa(s int) string {

// 	var myInt string
// 	for _, getChar := range s {
// 		if getChar >= '0' && getChar <= '9' {
// 			// fmt.Println("before *", myInt)
// 			myInt += string(getChar + '0')
// 			// fmt.Println("after +=", myInt)

// 		}
// 	}
// 	return myInt
// }

func ConcatSlice(slice1, slice2 []int) []int {

	for _, getElements := range slice2 {
		slice1 = append(slice1, getElements)
	}
	return slice1
}
