package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	varS := os.Args[0]
	for i, s := range varS { ///tmp/go-build3952465748/b001/exe/main /  // ./main
		if i >= 2 {
			z01.PrintRune(s) // ./main
		}
	}
	z01.PrintRune('\n')
}
