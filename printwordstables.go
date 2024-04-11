package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	// {Hello} {how} {are} {you?}
	// var reString string
	for _, getWord := range a {
		for _, getChar := range getWord {
			z01.PrintRune(getChar)
		}
		z01.PrintRune('\n')
	}
}
