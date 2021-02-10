// Tic Tac Toe inspired by 'A Tour of Go'
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		fmt.Print("Player 1, your turn: ")

		scanner.Scan()
		p1turn := scanner.Text()

		if p1turn != "" {
			num, err := strconv.Atoi(p1turn)
			if err != nil {
				panic(err)
			}
			row, col := pos(num)
			board[row][col] = "X"
			print(board)
		}
		fmt.Print("Player 2, your turn: ")

		scanner.Scan()
		p2turn := scanner.Text()

		if p2turn != "" {
			num, err := strconv.Atoi(p2turn)
			if err != nil {
				panic(err)
			}
			row, col := pos(num)
			board[row][col] = "O"
			print(board)
		}
	}
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

func pos(num int) (row, col int) {
	m := map[int]struct {
		row, col int
	}{
		1: {0, 0},
		2: {0, 1},
		3: {0, 2},
		4: {1, 0},
		5: {1, 1},
		6: {1, 2},
		7: {2, 0},
		8: {2, 1},
		9: {2, 2},
	}
	pos := m[num]
	return pos.row, pos.col
}
