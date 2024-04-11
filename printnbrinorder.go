package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var StoreValue int // 1
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var StoreNumbers []int
	if n > 0 {
		for n > 0 {
			StoreValue = n % 10
			n /= 10
			StoreNumbers = append(StoreNumbers, StoreValue)
		}
	}
	for i := 0; i < len(StoreNumbers); i++ {
		for j := i + 1; j < len(StoreNumbers); j++ {
			if StoreNumbers[i] > StoreNumbers[j] {
				StoreNumbers[i], StoreNumbers[j] = StoreNumbers[j], StoreNumbers[i]
			}
		}
	}
	for i := 0; i < len(StoreNumbers); i++ {
		z01.PrintRune(rune(StoreNumbers[i] + '0'))
	}
}
