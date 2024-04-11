package main

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	getArgs := os.Args[1]
	originalSlice := strings.Fields(getArgs)
	remixSlice := strings.Fields(getArgs)
	//Append in Slice

	// remix the slice
	for i := 0; i < len(originalSlice)-1; i++ {

		if i >= len(originalSlice)-2 {
			originalSlice[len(originalSlice)-1] = remixSlice[0]
		}
		originalSlice[i] = remixSlice[i+1]

	}

	// Print Slice
	for i, print := range originalSlice {
		for _, reprint := range print {
			z01.PrintRune(reprint)
		}
		if i > len(originalSlice)-2 {

		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')

}
