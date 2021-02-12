// Tic Tac Toe inspired by 'A Tour of Go'
package main

import (
	"fmt"
	game "tictactoe/game"
)

// $ clear && go run main.go
func main() {
	fmt.Println("Hey! This is Tic Tac Toe for 2 friends :)")

	game.PrintLogo()
	game.Setup()
	for game.Loop() {
	}
}
