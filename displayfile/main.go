package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("File name missing\n")
		return
	} else if len(os.Args) > 2 {
		fmt.Print("Too many arguments\n")
		return
	} else {
		nameOfFile := os.Args[1:]
		openFile, _ := os.ReadFile(nameOfFile[0])
		fmt.Print(string(openFile))
	}
}
