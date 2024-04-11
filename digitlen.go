package piscine

func DigitLen(n, base int) int {

	if base >= 2 && base <= 36 {
		count := 0
		if n < 0 {
			n = -n
		}
		for n > 0 {
			n /= base
			count++
		}
		return count
	}
	return -1
}
