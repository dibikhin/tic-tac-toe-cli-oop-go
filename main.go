// Tic Tac Toe inspired by 'A Tour of Go'
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grid [][]string

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

func main() {
	fmt.Println("Hey! This is PvP Tic Tac Toe :)")
	print(logo)

	// Setting up
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	player1 := scanner.Text()

	var player2 string
	if player1 == "x" {
		player2 = "o"
	} else {
		player2 = "x"
	}
	fmt.Println()
	fmt.Println("Player 1 is:", strings.ToUpper(player1))
	fmt.Println("Player 2 is:", strings.ToUpper(player2))

	print(board)

	// Game loop
	for {
		if !hasEmpty(board) {
			break
		}
		fmt.Print("Player 1, your turn: ")
		p1turn := getInput(scanner)
		for !isTurn(p1turn) {
			print(board)
			fmt.Print("Player 1, your turn: ")
			p1turn = getInput(scanner)
		}
		row, col := pos(p1turn)
		board[row][col] = "X"
		print(board)

		if !hasEmpty(board) {
			break
		}
		fmt.Print("Player 2, your turn: ")
		p2turn := getInput(scanner)
		for !isTurn(p2turn) {
			print(board)
			fmt.Print("Player 2, your turn: ")
			p2turn = getInput(scanner)
		}
		row, col = pos(p2turn)
		board[row][col] = "O"
		print(board)
	}
}

func getInput(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

func print(b grid) {
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark, then press ENTER (e.g. 5 is center). Board:")

	fmt.Println()
	for _, r := range b {
		fmt.Printf("%s\n", strings.Join(r, " "))
	}
	fmt.Println()
}

func hasEmpty(b grid) bool {
	for _, r := range b {
		for _, v := range r {
			if v == "_" {
				return true
			}
		}
	}
	return false
}

func isTurn(s string) bool {
	st := strings.TrimSpace(s)
	for _, v := range strings.Split("123456789", "") {
		if st == v {
			return true
		}
	}
	return false
}

func pos(key string) (row, col int) {
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
