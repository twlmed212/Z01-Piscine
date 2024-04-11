package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	lengthOfArg := os.Args[1:]
	count := 0
	for range lengthOfArg {
		count += 1
	}
	// fmt.Print(count)
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(count) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
