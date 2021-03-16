// Tic-tac-toe inspired by 'A Tour of Go'
package main

import (
	"fmt"

	game "tictactoe/game"
)

// $ clear && go run main.go
func main() {
	fmt.Println("Hey! This is 3x3 Tic-tac-toe for 2 friends :)")

	game.Setup(game.Read)
	_, ok := game.Loop(game.Read)
	for ok {
		_, ok = game.Loop(game.Read)
	}
}
