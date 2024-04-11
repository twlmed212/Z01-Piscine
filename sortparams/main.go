package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	getArgs := os.Args[1:]

	for i := 0; i < len(getArgs); i++ {
		for j := i + 1; j < len(getArgs); j++ {
			if getArgs[i] > getArgs[j] {
				getArgs[i], getArgs[j] = getArgs[j], getArgs[i]
			}
		}
	}

	for _, getRune := range getArgs {
		for _, PrintRune := range getRune {
			z01.PrintRune(PrintRune)
		}
		z01.PrintRune('\n')
	}
}
