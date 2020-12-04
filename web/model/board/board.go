package board

import (
	"fmt"
)

// GameBoard is a type to hold all the steps information and functions to act on it
type GameBoard [9]string

// Print the board
func (gb *GameBoard) Print() {
	fmt.Println("\n-----------")
	for i, v := range gb {
		if (i+1)%3 == 0 {
			fmt.Printf(" %v ", v)
			fmt.Println("\n-----------")
		} else {
			fmt.Printf(" %v |", v)
		}
	}
	fmt.Println()
}