package main

import (
	"os"

	"github.com/01-edu/z01"
	// "strconv"
	// "github.com/01-edu/z01"
)

func main() {
	getArgs := os.Args[1:]
	var runeSlice []string
	for _, runeArgs := range getArgs {
		runeSlice = append(runeSlice, runeArgs)
	}
	for i := 0; i < len(runeSlice); i++ {
		// Convert, _ := strconv.Atoi(runeSlice[i])
	}
	for i := 0; i < len(runeSlice); i++ {
		for j := 0; j < len(runeSlice[i]); j++ {
			// if runeSlice[i] == j {
			z01.PrintRune(rune(j))
			// }
		}
	}
	// fmt.Print(runeSlice)
}
