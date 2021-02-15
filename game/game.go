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
	if !board.hasEmpty() {
		fmt.Println("Draw!")
		return false
	}
	fmt.Printf("Player %v (%v), your turn: ", n, player)
	turn := read(scanner)

	var row, col int
	for {
		if !isKey(turn) {
			board.print()
			fmt.Printf("Player %v (%v), your turn: ", n, player)
			turn = read(scanner)

			continue
		}
		row, col = pos(turn)
		if board.isFilled(row, col) {
			board.print()
			fmt.Printf("Player %v (%v), your turn: ", n, player)
			turn = read(scanner)

			continue
		}
		break
	}
	board[row][col] = player

	if board.isWinner(player) {
		board.print()
		fmt.Printf("Player %v (%v) won!\n", n, player)
		return false
	}
	board.print()

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

func pos(key string) (int, int) {
	m := map[string]struct {
		row, col int
	}{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
	pos := m[key] // TODO: detect and propagate errors
	return pos.row, pos.col
}
