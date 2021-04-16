// Package game implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet).
// Players choose their marks, put them, the game checks winner or draw.
package game

import (
	"errors"
	"fmt"
	"strings"
)

type reader func() string

// Constants
var (
	// ErrCouldNotStart arises when `Loop()` is run without running `Setup()` first.
	ErrCouldNotStart = errors.New("couldn't start the game loop, set up the game first")

	// ErrNilReader arises when `Setup()` is run with nil reader.
	ErrNilReader = errors.New("the reader is nil, pass the default reader at least")
)

// Private package state.
// It's here to simplify dependency injection.
// There was no need to expose the private state as context
var _game *game

// Public

// Setup initializes the game and helps players to choose marks.
// The `read` param is a strategy for user input to prevent mocking.
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

// Read gets players's input and returns it as a text.
// It's a default impl of the `reader` Strategy. It's exposed to be used
// for testing to prevent mocking.
func Read() string {
	_game.scanner.Scan()
	return strings.TrimSpace(_game.scanner.Text())

	// TODO: have to check and propagate _scanner.Err() ?
}

// Game Loop()

// Loop prompts players to take turns.
// The `board` is returned for tests only.
func Loop() (board, bool, error) {
	if _game == nil || !_game.isReady {
		return board{}, false, ErrCouldNotStart
	}
	more := turn(_game.player1, _game.reader, &_game.board)
	if !more {
		return _game.board, false, nil
	}
	more = turn(_game.player2, _game.reader, &_game.board)
	return _game.board, more, nil
}

// Private

func turn(them player, read reader, board *board) bool {
	c := inputLoop(read, them, *board)
	board.setCell(c, them.mark)
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

func inputLoop(read reader, pers player, board board) cell {
	prompt(pers)
	for {
		turn := key(read())
		if !turn.isKey() {
			board.print()
			prompt(pers)

			continue
		}
		cel := turn.toCell()
		if board.isFilled(cel) {
			board.print()
			prompt(pers)

			continue
		}
		return cel
	}
}
