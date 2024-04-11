package main

import (
	"fmt"
	"os"
)

func main() {
	getArgs := os.Args[1:]
	count := 0
	for _, checkChar := range getArgs {
		if count == 0 && (checkChar == "01" || checkChar == "galaxy" || checkChar == "galaxy 01") {
			fmt.Println("Alert!!!")
			count++
		}
	}
}
