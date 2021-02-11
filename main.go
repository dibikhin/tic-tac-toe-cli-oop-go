// Tic Tac Toe inspired by 'A Tour of Go'
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grid [3][3]string

func (b grid) isEmpty(row, col int) bool {
	v := b[row][col]
	return v == "_"
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
	// Need something better, too naive

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

var logo = grid{
	{"X", " ", "X"},
	{"O", "X", "O"},
	{"X", " ", "O"},
}

var board = grid{
	{"_", "_", "_"},
	{"_", "_", "_"},
	{"_", "_", "_"},
}

// $ clear && go run main.go
func main() {
	fmt.Println("Hey! This is Tic Tac Toe with friend :)")
	printLogo()

	// Setting up
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	scanner := bufio.NewScanner(os.Stdin)
	player1 := getInput(scanner)

	var player2 string
	if strings.ToLower(player1) == "x" {
		player1 = "X"
		player2 = "O"
	} else {
		player1 = "O"
		player2 = "X"
	}
	fmt.Println()
	fmt.Println("Player 1 is:", player1)
	fmt.Println("Player 2 is:", player2)

	print(board)

	// Game loop
	for {
		var row, col int

		if !board.hasEmpty() {
			fmt.Println("Draw!")
			return
		}
		fmt.Printf("Player 1 (%v), your turn: ", player1)
		p1turn := getInput(scanner)
		for {
			if !isKey(p1turn) {
				print(board)
				fmt.Printf("Player 1 (%v), your turn: ", player1)
				p1turn = getInput(scanner)
				continue
			}
			row, col = pos(p1turn)
			if !board.isEmpty(row, col) {
				print(board)
				fmt.Printf("Player 1 (%v), your turn: ", player1)
				p1turn = getInput(scanner)
				continue
			}
			break
		}
		board[row][col] = player1
		if board.isWinner(player1) {
			print(board)
			fmt.Printf("Player 1 (%v) won!", player1)
			fmt.Println()
			return
		}
		print(board)

		if !board.hasEmpty() {
			fmt.Println("Draw!")
			return
		}
		fmt.Printf("Player 2 (%v), your turn: ", player2)
		p2turn := getInput(scanner)
		for {
			if !isKey(p2turn) {
				print(board)
				fmt.Printf("Player 2 (%v), your turn: ", player2)
				p2turn = getInput(scanner)
				continue
			}
			row, col = pos(p2turn)
			if !board.isEmpty(row, col) {
				print(board)
				fmt.Printf("Player 2 (%v), your turn: ", player2)
				p2turn = getInput(scanner)
				continue
			}
			break
		}
		board[row][col] = player2
		if board.isWinner(player2) {
			print(board)
			fmt.Printf("Player 2 (%v) won!", player2)
			fmt.Println()
			return
		}
		print(board)
	}
}

func getInput(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

func print(b grid) {
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")

	_print(b)
}

func printLogo() {
	_print(logo)
}

func _print(b grid) {
	fmt.Println()
	for _, r := range b {
		fmt.Printf("%s\n", strings.Join(r[:], " "))
	}
	fmt.Println()
}

func isKey(s string) bool {
	st := strings.TrimSpace(s)
	for _, v := range strings.Split("123456789", "") {
		if st == v {
			return true
		}
	}
	return false
}

func pos(key string) (int, int) {
	m := map[string]struct {
		row, col int
	}{
		"1": {0, 0},
		"2": {0, 1},
		"3": {0, 2},
		"4": {1, 0},
		"5": {1, 1},
		"6": {1, 2},
		"7": {2, 0},
		"8": {2, 1},
		"9": {2, 2},
	}
	pos := m[key]
	return pos.row, pos.col
}
