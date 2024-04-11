package piscine

func Atoi(nmbr string) int {
	number := 0
	minus := 1
	count := 0
	for i, v := range nmbr {
		if i == 0 && v == '-' {
			minus = -1
		} else if v == '+' {
			minus = 1
			count++
		} else if v >= '0' && v <= '9' && count <= 1 {
			number *= 10
			number += int(v - '0')
		} else if v == ' ' {
			return 0
		} else {
			return 0
		}
	}
	return number * minus
}
func BasicAtoi2(nb string) int {
	nmbr := 0
	for _, v := range nb {
		if v >= '0' && v <= '9' {
			nmbr *= 10
			nmbr += int(v - '0')
		} else {
			return 0
		}
	}

	return nmbr
}
