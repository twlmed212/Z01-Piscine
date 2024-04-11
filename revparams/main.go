package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	getArgs := os.Args
	var saveArray string
	var saveSlice []string
	for i, get := range getArgs {
		if i > 0 {
			for _, results := range get {
				saveArray += string(results)
			}
			saveSlice = append(saveSlice, saveArray)
			saveArray = ""
		}
	}
	for i := len(saveSlice) - 1; i >= 0; i-- {
		getRun := []rune(saveSlice[i])
		for j := 0; j < len(getRun); j++ {
			z01.PrintRune(getRun[j])
		}
		// fmt.Print(saveSlice[i])
		z01.PrintRune('\n')
	}
}
