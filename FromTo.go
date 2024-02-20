package main

import (
	"fmt"
	"strconv"

	//"github.com/01-edu/z01"
)

func FromTo(from int, to int) string {
	store := ""
	zero := "0"
	comma := ", "
	if from > 0 && to > 0 && from < 100 && to < 100 {
		if from > to {
			for i := from; i >= to; i-- {
				if i < 10 {
					store += zero
					store += strconv.Itoa(i)
					if i == to {
						continue
					}
					store += comma

				} else {
					store += strconv.Itoa(i)
					store += comma
				}
			}

		} else if from < to {
			for i := from; i <= to; i++ {
				if i < 10 {
					store += zero
					store += strconv.Itoa(i)
					if i == to {
						continue
					}
					store += comma

				} else {
					if to != i {
						store += strconv.Itoa(i)
						store += comma
						
					}else{
						store += strconv.Itoa(i)
						continue
					}
				}
			}
		} else if to == from {
			if from < 10 || to < 10 {
				store += zero
				store += strconv.Itoa(from)
				
			} else {
				store += strconv.Itoa(to)
				// new line
				
			}
			
		} else {
			fmt.Print("Invalid")
		}
	} else {
		fmt.Print("Invalid")
	}
	new := store 
	new += "\n"
	return new
}

func main() {
	fmt.Print(FromTo(5, 5))
	fmt.Print(FromTo(1, 5))
	fmt.Print(FromTo(99, 5))
	fmt.Print(FromTo(99, 1))
}
