package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Init

var logo grid
var board grid

var scanner *bufio.Scanner

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
	scanner = bufio.NewScanner(os.Stdin)
}

// Game

func PrintLogo() {
	_print(logo)
}

func Setup() {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	mark1 := read(scanner)
	player1, player2 = arrange(mark1)

	fmt.Println()
	fmt.Println("Player 1 is:", player1)
	fmt.Println("Player 2 is:", player2)

	board.print()
}

func Loop() bool {
	ok := move(1, player1)
	if !ok {
		return false
	}
	ok = move(2, player2)
	return ok
}

func move(n int, player string) bool {
	fmt.Printf("Player %v (%v), your turn: ", n, player)
	turn := read(scanner)

	var c cell
	// Input loop
	for {
		if !isKey(turn) {
			board.print()
			fmt.Printf("Player %v (%v), your turn: ", n, player)
			turn = read(scanner)

			continue
		}
		c = toCell(turn)
		if board.isFilled(c) {
			board.print()
			fmt.Printf("Player %v (%v), your turn: ", n, player)
			turn = read(scanner)

			continue
		}
		break
	}
	// The turn is key and the cell is empty
	board.setCell(c, player)
	board.print()

	// Finished?
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

// Other

func read(bs *bufio.Scanner) string {
	bs.Scan()
	return strings.TrimSpace(bs.Text())
}

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
