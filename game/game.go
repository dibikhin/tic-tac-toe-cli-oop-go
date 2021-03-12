// Package Game implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet)
// Players choose their mark, put them, the game checks the winner or draw
package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mark = string // to avoid conversions

type player struct {
	mark mark
	num  int
}

func (p player) String() string {
	return fmt.Sprintf(`Player %v ("%v")`, p.num, p.mark)
}

type reader func() string

// Package state.
// It's here to simplify dependency injection.

var (
	_logo  board
	_board board

	_player1 player
	_player2 player

	_scanner *bufio.Scanner
)

// Public

// Setup initialized the game and helps players to choose mark.
// The `read` param is a strategy to prevent mocking
func Setup(read reader) {
	_init()
	printLogo()
	setupGame(read)
}

// Loop prompts players to take turns.
// The `read` param is a strategy to prevent mocking
// The `board` is returned for testing
func Loop(read reader) (board, bool) {
	ok := turn(_player1, read)
	if !ok {
		return _board, false
	}
	ok = turn(_player2, read)
	return _board, ok
}

// IO

// Read gets players's input and returns it as a text.
// It's a default impl of the `reader` Strategy. It's used for testing to prevent mocking.
func Read() string {
	_scanner.Scan()
	return strings.TrimSpace(_scanner.Text())
}

// Private

// Setup

func _init() {
	_logo = board{
		{"X", " ", "X"},
		{"O", "X", "O"},
		{"X", " ", "O"},
	}
	_board = board{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	_scanner = bufio.NewScanner(os.Stdin)
}

func setupGame(read reader) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	mark1 := read()
	_player1, _player2 = arrange(mark1)

	fmt.Println()
	fmt.Println(_player1)
	fmt.Println(_player2)

	_board.print()
}

// Game

func turn(p player, read reader) bool {
	prompt(p)

	cell := inputLoop(read, p)
	_board.setCell(cell, p.mark)
	_board.print()

	if _board.isWinner(p.mark) {
		fmt.Printf("%v won!\n", p)
		return false
	}
	if !_board.hasEmpty() {
		fmt.Println("Draw!")
		return false
	}
	return true
}

func inputLoop(read reader, p player) cell {
	var c cell
	for {
		turn := read()
		if !isKey(turn) {
			_board.print()
			prompt(p)

			continue
		}
		c = toCell(turn)
		if _board.isFilled(c) {
			_board.print()
			prompt(p)

			continue
		}
		break
	}
	return c
}

// IO

func printLogo() {
	fmt.Println()
	fmt.Println(_logo)
	fmt.Println()
}

func prompt(p player) {
	var _ fmt.Stringer = player{}
	fmt.Printf("%v, your turn: ", p)
}

// Other

func arrange(s string) (player, player) {
	if strings.ToLower(s) == "x" {
		return player{"X", 1}, player{"O", 2}
	} else {
		return player{"O", 1}, player{"X", 2}
	}
}

func isKey(s string) bool {
	k, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return k >= 1 && k <= 9
}
