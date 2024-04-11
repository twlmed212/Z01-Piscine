package piscine

import (
	//"strconv"

	//"strconv"

	"github.com/01-edu/z01"
)

func main() {
	PrintNbr(123)
	// PrintNbr(0)
	// PrintNbr(123)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	} else if n == 0 {
		z01.PrintRune('0')
		return
	}
	counter := '0'
	for i := 1; i <= n%10; i++ {
		counter++
	}
	for i := -1; i >= n%10; i-- {
		counter++
	}
	z01.PrintRune(counter)
	z01.PrintRune('\n')
}
