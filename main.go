// Tic-tac-toe inspired by 'A Tour of Go'.
package main

import (
	"fmt"
	"os"

	game "tictactoe/game"
)

// $ clear && go run main.go
func main() {
	fmt.Println("Hey! This is 3x3 Tic-tac-toe for 2 friends :)")

	err := game.PlayGame()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.\n", err)
		os.Exit(1)
	}
}
