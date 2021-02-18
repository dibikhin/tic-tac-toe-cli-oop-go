// Package game implements 3x3 Tic Tac Toe for 2 friends (cannot play with computer yet)
// Players choose their mark, put them, then game checks the winner or draw
package game

import (
	"fmt"
	"strconv"
	"strings"
)

type reader func() string

// Init

var logo grid
var board grid

var player1 string
var player2 string

func init() {
	logo = grid{
		{"X", " ", "X"},
		{"O", "X", "O"},
		{"X", " ", "O"},
	}
	board = grid{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
}

// Game

func PrintLogo() {
	fmt.Println()
	fmt.Println(logo)
	fmt.Println()
}

// Setup helps users to choose mark
func Setup(read reader) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	mark1 := read()
	player1, player2 = arrange(mark1)

	fmt.Println()
	fmt.Println("Player 1 is:", player1)
	fmt.Println("Player 2 is:", player2)

	board.print()
}

// Loop function prompts players to take turns
func Loop(read reader) bool {
	ok := turn(1, player1, read)
	if !ok {
		return false
	}
	ok = turn(2, player2, read)
	return ok
}

func turn(n int, player string, read reader) bool {
	prompt(player, n)

	cell := input(read, player, n)
	board.setCell(cell, player)
	board.print()

	if board.isWinner(player) {
		fmt.Printf("Player %v (%v) won!\n", n, player)
		return false
	}
	if !board.hasEmpty() {
		fmt.Println("Draw!")
		return false
	}
	return true
}

// Input loop
func input(read reader, player string, n int) cell {
	var c cell
	for {
		turn := read()
		if !isKey(turn) {
			board.print()
			prompt(player, n)

			continue
		}
		c = toCell(turn)
		if board.isFilled(c) {
			board.print()
			prompt(player, n)

			continue
		}
		break
	}
	return c
}

// Other

func arrange(s string) (string, string) {
	if strings.ToLower(s) == "x" {
		return "X", "O"
	} else {
		return "O", "X"
	}
}

func isKey(s string) bool {
	k, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return k >= 1 && k <= 9
}

// IO

func prompt(player string, n int) {
	fmt.Printf("Player %v (%v), your turn: ", n, player)
}
