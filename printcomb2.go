package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	/*
		//00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$
		  ab cd, ab cd, ab cd, ab cd, ab cd, ab cd, ab cd, ab cd, ab cd, ab cd, ab cd,
	*/
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					if (a*10)+b < (c*10)+d {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)
						if !(a == '9' && b == '8' && c == '9' && d == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}

	z01.PrintRune('\n')
}
