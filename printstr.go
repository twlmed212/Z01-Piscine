package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	lenght := len(s)
	for x := 0; x < lenght; x++ {
		value := s[x]
		z01.PrintRune(rune(value))
	}
}
