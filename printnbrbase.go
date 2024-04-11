package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nmbr int, base string) {
	calculatedBase := ""
	minus := false
	overFlow := false
	if isValid(base) {
		if nmbr < 0 {
			if nmbr <= -9223372036854775807 {
				overFlow = true
				nmbr = -9223372036854775807
			}
			minus = true
			nmbr = -nmbr
		} //

		for nmbr > 0 {
			if overFlow {
				temp := nmbr % len(base)
				temp = temp + 1
				nmbr /= len(base)
				calculatedBase = string(base[temp]) + calculatedBase
				overFlow = false
			} else {
				temp := nmbr % len(base)
				nmbr /= len(base)
				calculatedBase = string(base[temp]) + calculatedBase
			}
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
	if minus {
		calculatedBase = "-" + calculatedBase
	}
	for _, get := range calculatedBase {
		z01.PrintRune(get)
	}
}

func isValid(base string) bool {
	myBase := ""
	for _, getChar := range base {
		isDouble := false
		if getChar == '-' || getChar == '+' {
			return false
		}
		for _, Compare := range myBase {
			if Compare == getChar {
				isDouble = true
			}
		}
		if !isDouble {
			myBase += string(getChar)
		}
	}
	if len(base) != len(myBase) || len(base) < 2 {
		return false
	} else {
		return true
	}
}
