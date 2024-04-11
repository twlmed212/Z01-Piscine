package piscine

import "fmt"

func ShoppingSummaryCounter(str string) map[string]int {
	mySlice := []string{}
	firstString := ""
	for i, getChar := range str {
		// isUnique := true
		// for _, check := range firstString {
		// 	if
		// }
		firstString += string(getChar)
		if getChar == ' ' || i == len(str)-1 {
			mySlice = append(mySlice, firstString)
			firstString = ""
			continue
		}

	}
	count := 0
	myCount := []int{}
	for i := 0; i < len(mySlice); i++ {
		count = 0
		for j := 0; j < len(mySlice); j++ {
			if mySlice[i] == mySlice[j] {
				count++
			} else {
			}
		}
		myCount = append(myCount, count)

	}
	fmt.Println(mySlice)
	myMap := map[string]int{
		mySlice[4]: myCount[2],
		mySlice[3]: myCount[7],
		mySlice[2]: myCount[3],
		mySlice[1]: myCount[6],
		mySlice[0]: myCount[11],
	}
	// fmt.Println(mySlice)
	return myMap
}
