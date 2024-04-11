package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		theArgs := os.Args[1:]
		firstNumber := 0
		secondNumber := 0
		sign := false
		Numberfound := 0
		result := 0
		countSign := 0
		for _, string := range theArgs {
			for _, char := range string {
				if char == '+' || char == '*' || char == '-' || char == '/' || char == '%' {
					sign = true
					countSign++
					if char == '+' {
						if result != 0 {
							result = firstNumber + result
						} else {
							result = firstNumber + secondNumber
						}
					}
					if char == '-' {

						if result != 0 {
							result = result - firstNumber
						} else {
							result = firstNumber - secondNumber
						}
					}
					if char == '*' {
						if result != 0 {
							result = firstNumber * result
						} else {
							result = firstNumber * secondNumber
						}
					}
				} else if char >= '0' && char <= '9' {
					Numberfound++ //1

					if Numberfound == 1 {

						firstNumber = int(char - '0') //2
						if sign {
							Numberfound = 0
							continue
						}
					} else {

						secondNumber = int(char - '0') //3
						Numberfound = 0
					}
				} else {
					if char == ' ' {
						continue
					} else {

						fmt.Println("Error")
						return
					}

				}
			}
		}
		if countSign == 1 {
			fmt.Println("Error")
			return
		}
		fmt.Println(result)

	} else {
		fmt.Println("Error")
	}
}
