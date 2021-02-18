// Tic Tac Toe inspired by 'A Tour of Go'
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	game "tictactoe/game"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

// $ clear && go run main.go
func main() {
	fmt.Println("Hey! This is 3x3 Tic Tac Toe for 2 friends :)")

	game.PrintLogo()
	game.Setup(read)
	for game.Loop(read) {
	}
}

// IO

func read() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
