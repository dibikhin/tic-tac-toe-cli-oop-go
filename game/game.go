// Package Game implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet)
// Players choose their mark, put them, the game checks the winner or draw
package game

import (
	"bufio"
	"errors"
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
// There was no need to expose the private state as context.

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
	_logo, _board, _scanner = _init()
	printLogo(_logo)
	_player1, _player2 = setupGame(read, _board)

	_ready = true
}

// Loop prompts players to take turns.
// The `read` param is a strategy to prevent mocking
// The `board` is returned for tests only
func Loop(read reader) (board, bool, error) {
	if !_ready {
		return _board, false, errors.New("setup failed")
	}
	more := turn(_player1, read, &_board)
	if !more {
		return _board, false, nil
	}
	more = turn(_player2, read, &_board)
	return _board, more, nil
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

func _init() (board, board, *bufio.Scanner) {
	logo := board{
		{"X", " ", "X"},
		{"O", "X", "O"},
		{"X", " ", "O"},
	}
	board := board{
		{_blank, _blank, _blank},
		{_blank, _blank, _blank},
		{_blank, _blank, _blank},
	}
	scanner := bufio.NewScanner(os.Stdin)
	return logo, board, scanner
}

func setupGame(read reader, b board) (player, player) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")
	mark1 := read()
	p1, p2 := arrange(mark1)

	fmt.Println()
	fmt.Println(p1)
	fmt.Println(p2)

	b.print()

	return p1, p2
}

// Game

func turn(them player, read reader, board *board) bool {
	prompt(them)

	cell := inputLoop(read, them, board)
	board.setCell(cell, them.mark)
	board.print()

	if board.isWinner(them.mark) {
		fmt.Printf("%v won!\n", them)
		return false
	}
	if !board.hasEmpty() {
		fmt.Println("Draw!")
		return false
	}
	return true
}

func inputLoop(read reader, pers player, board *board) cell {
	var cel cell
	for {
		turn := read()
		if !isKey(turn) {
			board.print()
			prompt(pers)

			continue
		}
		cel = toCell(turn)
		if board.isFilled(cel) {
			board.print()
			prompt(pers)

			continue
		}
		break
	}
	return cel
}

// IO

func printLogo(s fmt.Stringer) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()
}

func prompt(s fmt.Stringer) {
	fmt.Printf("%v, your turn: ", s)
}
