// Package game implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet)
// Players choose their marks, put them, the game checks winner or draw
package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ErrCouldNotStart arises when `Loop()` is run without running `Setup()` first
var ErrCouldNotStart = errors.New("couldn't start, set up the game first")

// ErrNilReader arises when `Setup()` is run with nil reader
var ErrNilReader = errors.New("reader is nil, pass the default reader at least")

type (
	reader func() string

	game struct {
		isReady bool

		logo board

		board   board
		player1 player
		player2 player

		scanner *bufio.Scanner
		reader  reader
	}
)

func newGame(read reader) *game {
	return &game{
		logo: board{
			{"X", " ", "X"},
			{"O", "X", "O"},
			{"X", " ", "O"}},
		board: board{
			{_blank, _blank, _blank}, {_blank, _blank, _blank}, {_blank, _blank, _blank}},
		scanner: bufio.NewScanner(os.Stdin),
		reader:  read,
	}
}

func (g *game) setPlayers(p1, p2 player) {
	g.player1 = p1
	g.player2 = p2
	g.isReady = true
}

// Private package state.
// It's here to simplify dependency injection.
// There was no need to expose the private state as context.
var _game *game

// Public

// Setup initializes the game and helps players to choose marks.
// The `read` param is a strategy for user input to prevent mocking
func Setup(read reader) error {
	if read == nil {
		return ErrNilReader
	}
	_game = newGame(read)
	printLogo(_game.logo)

	_game.setPlayers(chooseMarks(read))
	printGame(*_game)

	return nil
}

// Loop prompts players to take turns.
// The `board` is returned for tests only
func Loop() (board, bool, error) {
	if !_game.isReady {
		return _game.board, false, ErrCouldNotStart
	}
	more := turn(_game.player1, _game.reader, &_game.board)
	if !more {
		return _game.board, false, nil
	}
	more = turn(_game.player2, _game.reader, &_game.board)
	return _game.board, more, nil
}

// Read gets players's input and returns it as a text.
// It's a default impl of the `reader` Strategy. It's exposed to be used
// for testing to prevent mocking.
// (IO)
func Read() string {
	_game.scanner.Scan()
	return strings.TrimSpace(_game.scanner.Text())

	// TODO: have to check and propagate _scanner.Err() ?
}

// Private

// Setup() IO

func printLogo(s fmt.Stringer) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()

	fmt.Println("(Use `ctrl+c` to exit)")
	fmt.Println()
}

func chooseMarks(read reader) (player, player) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")
	mark1 := read()
	p1, p2 := arrange(mark1)
	return p1, p2
}

func printGame(g game) {
	fmt.Println()

	fmt.Println(g.player1)
	fmt.Println(g.player2)

	g.board.print()
}

// Other

func arrange(m mark) (player, player) {
	if strings.ToLower(m) == "x" {
		return player{"X", 1}, player{"O", 2}
	}
	return player{"O", 1}, player{"X", 2}
}

// Game Loop()

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
		turn := key(read())
		if !turn.isKey() {
			board.print()
			prompt(pers)

			continue
		}
		cel = turn.toCell()
		if board.isFilled(cel) {
			board.print()
			prompt(pers)

			continue
		}
		break
	}
	return cel
}
