package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	numberOfPlayers := 1
	cardsPerPlayer := 3
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d:", numberOfPlayers)
		for j := 0; j < cardsPerPlayer; j++ {
			cardIndex := i*cardsPerPlayer + j
			if j != cardsPerPlayer-1 {
				fmt.Printf(" %d,", deck[cardIndex])
			} else {
				fmt.Printf(" %d", deck[cardIndex])
			}
		}
		numberOfPlayers++
		fmt.Printf("\n")
	}
}

// func DealAPackOfCards(deck []int) {
// 	numberOfPlayers := 1
// 	players := 3
// 	// lastElement := deck[len(deck)-1]
// 	// numberOfPlayers := lastElement/4
// 	count := 0
// 	for range "four" {
// 		fmt.Printf("Player %d:", numberOfPlayers)
// 		for i := range deck[1:4] {
// 			fmt.Print(" ", i*players+, ", ")
// 			if count == 3 {
// 				count = 0
// 			}
// 			count++
// 		}
// 		numberOfPlayers++
// 		z01.PrintRune('\n')

// 	}
// 	// fmt.Print(lastElement)
// }
