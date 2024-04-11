package piscine

import "github.com/01-edu/z01"

func SaveAndMiss(s string, num int) string {

	save := 0
	miss := 0
	for _, check := range s {
		if save < 3 {
			z01.PrintRune(check)
			save++
			miss++

		} else if miss <= 3 {
			if miss == 1 {
				save = 0
			}
			miss--

		}
	}
	return ""
}
