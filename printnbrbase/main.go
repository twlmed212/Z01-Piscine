package main

import "fmt"

func isValid(base string) bool {
	UniqueBase := ""
	for _, check := range base {
		isUnique := true
		for _, CheckBase := range UniqueBase {
			if check == CheckBase {
				isUnique = false
			}
		}
		if isUnique {
			UniqueBase += string(check)
		}
	}
	return len(UniqueBase) == len(base)
}

func atoi(s string) int {
	nmbr := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			nmbr *= 10
			nmbr += int(char - '0')
		}
	}
	return nmbr
}
func printNbrBase(s string, base string) string {
	if isValid(base) {
		// 125%16 =      13
		// 125/16 = 7
		// 7%16 =        7
		intNumber := atoi(s)
		myString := ""
		for intNumber > 0 {
			mod := intNumber % len(base)
			myString = string(base[mod]) + myString
			intNumber /= len(base)
		}
		if s[0] == '-' {
			myString = "-" + myString
		}
		return myString

	}
	return "NV"

}
func main() {
	fmt.Println(printNbrBase("125", "0123456789ABCDEF"))
	fmt.Println(printNbrBase("125", "0123456789"))
	fmt.Println(printNbrBase("125", "01"))
	fmt.Println(printNbrBase("-125", "choumi"))
	fmt.Println(printNbrBase("125", "01234567"))
	fmt.Println(printNbrBase("125", "012234567"))
}
