package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Grid

type grid [3][3]string

func (b grid) isFilled(row, col int) bool {
	v := b[row][col]
	return v != "_"
}

func (b grid) hasEmpty() bool {
	for _, r := range b {
		for _, v := range r {
			if v == "_" {
				return true
			}
		}
	}
	return false
}

func (b grid) isWinner(p string) bool {
	// Something better needed, too naive

	// Horizontal
	x0 := b[0][0] == p && b[0][1] == p && b[0][2] == p
	x1 := b[1][0] == p && b[1][1] == p && b[1][2] == p
	x2 := b[2][0] == p && b[2][1] == p && b[2][2] == p

	// Vertical
	x3 := b[0][0] == p && b[1][0] == p && b[2][0] == p
	x4 := b[0][1] == p && b[1][1] == p && b[2][1] == p
	x5 := b[0][2] == p && b[1][2] == p && b[2][2] == p

	// Diagonal
	x6 := b[0][0] == p && b[1][1] == p && b[2][2] == p
	x7 := b[0][2] == p && b[1][1] == p && b[2][0] == p

	return x0 || x1 || x2 || x3 || x4 || x5 || x6 || x7
}

func (b grid) print() {
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")

	_print(b)
}

func _print(b grid) {
	fmt.Println()
	for _, r := range b {
		fmt.Printf("%s\n", strings.Join(r[:], " "))
	}
	fmt.Println()
}

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

func arrange(m string) (string, string) {
	if strings.ToLower(m) == "x" {
		return "X", "O"
	} else {
		return "O", "X"
	}
}

func isKey(s string) bool {
	for _, v := range strings.Split("123456789", "") {
		if s == v {
			return true
		}
	}
	return false
}

func pos(key string) (int, int) {
	m := map[string]struct {
		row, col int
	}{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
	pos := m[key] // TODO: detect and propagate error
	return pos.row, pos.col
}
