// Package Game implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet)
// Players choose their mark, put them, the game checks the winner or draw
package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Game types

type reader func() string

// Player

type player struct {
	mark mark
	num  int
}

func (p player) String() string {
	return fmt.Sprintf(`Player %v ("%v")`, p.num, p.mark)
}

func arrange(m mark) (player, player) {
	if strings.ToLower(m) == "x" {
		return player{"X", 1}, player{"O", 2}
	} else {
		return player{"O", 1}, player{"X", 2}
	}
}

// Private package state.
// It's here to simplify dependency injection.

var (
	_ready = false

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

	_ready = true
}

// Loop prompts players to take turns.
// The `read` param is a strategy to prevent mocking
// The `board` is returned for tests only
func Loop(read reader) (board, bool) {
	if !_ready {
		return _board, false
	}
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

	// TODO have to check and propagate _scanner.Err() ?
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
		{_blank, _blank, _blank},
		{_blank, _blank, _blank},
		{_blank, _blank, _blank},
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

func turn(them player, read reader) bool {
	prompt(them)

	cell := inputLoop(read, them)
	_board.setCell(cell, them.mark)
	_board.print()

	if _board.isWinner(them.mark) {
		fmt.Printf("%v won!\n", them)
		return false
	}
	if !_board.hasEmpty() {
		fmt.Println("Draw!")
		return false
	}
	return true
}

func inputLoop(read reader, pers player) cell {
	var cel cell
	for {
		turn := read()
		if !isKey(turn) {
			_board.print()
			prompt(pers)

			continue
		}
		cel = toCell(turn)
		if _board.isFilled(cel) {
			_board.print()
			prompt(pers)

			continue
		}
		break
	}
	return cel
}

// IO

func printLogo() {
	fmt.Println()
	fmt.Println(_logo)
	fmt.Println()
}

func prompt(s fmt.Stringer) {
	fmt.Printf("%v, your turn: ", s)
}
