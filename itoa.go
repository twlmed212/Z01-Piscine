package piscine

func Itoa(n int) string {
	myString := ""
	sign := false
	if n == 0 {
		return "0"
	}
	if n < 0 {
		n = -n
		sign = true
	}
	for n > 0 {
		temp := n % 10
		n /= 10
		myString = string(temp+'0') + myString
	}
	if sign {
		myString = "-" + myString
	}
	return myString

}
