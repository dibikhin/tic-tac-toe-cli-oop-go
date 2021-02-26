// Package Game implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet)
// Players choose their mark, put them, then game checks the winner or draw
package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mark = string // to avoid conversions

type reader func() string

// Init

var logo grid
var board grid

var player1 mark
var player2 mark

var scanner *bufio.Scanner

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
	scanner = bufio.NewScanner(os.Stdin)
}

// IO

// Read gets user's input and returns it as a text.
// It's a default impl of the `reader` Strategy. It's used for testing to prevent mocking.
func Read() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Game

func PrintLogo() {
	fmt.Println()
	fmt.Println(logo)
	fmt.Println()
}

// Setup helps users to choose mark.
// The `read` param is a strategy to prevent mocking
func Setup(read reader) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	mark1 := read()
	player1, player2 = arrange(mark1)

	fmt.Println()
	fmt.Println("Player 1 is:", player1)
	fmt.Println("Player 2 is:", player2)

	board.print()
}

// Loop function prompts players to take turns.
// The `read` param is a strategy to prevent mocking
// The `grid` is returned for testing
func Loop(read reader) (grid, bool) {
	ok := turn(1, player1, read)
	if !ok {
		return board, false
	}
	ok = turn(2, player2, read)
	return board, ok
}

func turn(n int, player mark, read reader) bool {
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
func input(read reader, player mark, n int) cell {
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

func arrange(s string) (mark, mark) {
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

func prompt(player mark, n int) {
	fmt.Printf("Player %v (%v), your turn: ", n, player)
}
