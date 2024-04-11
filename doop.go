package piscine

func Doop(args []string) int {
	results := 0
	firstValue := Atoi(args[0])
	sign := args[1]
	secondValue := Atoi(args[2])
	if sign == "+" {
		results = firstValue + secondValue
	}
	if sign == "-" {
		results = firstValue - secondValue
	}
	if sign == "/" {
		results = firstValue / secondValue
	}
	if sign == "%" {
		results = firstValue % secondValue
	}
	//fmt.Print(firstValue, sign, thirdValue)
	return results
}
