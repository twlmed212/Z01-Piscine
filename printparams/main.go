package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	getArgs := []string(os.Args)

	for i := 1; i < len(getArgs); i++ {
		results := getArgs[i]
		for j := 0; j < len(results); j++ {
			z01.PrintRune(rune(results[j]))
		}
		z01.PrintRune('\n')

	}
}
