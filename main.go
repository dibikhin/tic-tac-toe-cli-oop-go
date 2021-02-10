// Tic Tac Toe inspired by 'A Tour of Go'
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type board [][]string

var logo = board{
	{"X", " ", "X"},
	{"O", "X", "O"},
	{"X", " ", "O"},
}

var empty = board{
	{"_", "_", "_"},
	{"_", "_", "_"},
	{"_", "_", "_"},
}

func main() {
	fmt.Println("Hey! This is PvP Tic Tac Toe :)")
	printBoard(logo)

	// Setting up
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	player1 := scanner.Text()
	fmt.Println()

	var player2 string
	if player1 == "x" {
		player2 = "o"
	} else {
		player2 = "x"
	}
	fmt.Println("Player 1 is:", strings.ToUpper(player1))
	fmt.Println("Player 2 is:", strings.ToUpper(player2))

	// Game loop
	for {
		fmt.Println()
		fmt.Println("Press 1 to 9 to mark, then press ENTER (e.g. 5 is center). Board:")

		printBoard(empty)
		fmt.Print("Player 1, your turn: ")

		board := make(board, len(empty))
		copy(board, empty)

		scanner.Scan()
		p1turn := scanner.Text()
		if p1turn != "" {
			board[1][1] = "X"
			printBoard(board)
		}
		fmt.Print("Player 2, your turn: ")

		scanner.Scan()
		p2turn := scanner.Text()
		if p2turn != "" {
			board[2][2] = "O"
			printBoard(board)
		}
	}
}

func printBoard(b board) {
	fmt.Println()
	for _, r := range b {
		fmt.Printf("%s\n", strings.Join(r, " "))
	}
	fmt.Println()
}
