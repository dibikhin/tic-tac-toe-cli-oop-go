// Package game implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet).
// Players choose their marks, put them, the game checks winner or draw.
package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Constants, Private
var (
	// errCouldNotStart arises when `Loop()` is run without running `Setup()` first.
	errCouldNotStart = errors.New("couldn't start the game loop, set up the game first")

	// errNilReader arises when `Setup()` is run with nil reader.
	errNilReader = errors.New("the reader is nil, pass the default reader at least")
)

// Private package state.
// It's here to simplify dependency injection.
// There was no need to expose the private state as context
var _game *game

// Public

// Setup initializes the game and helps players to choose marks.
// The `read` param is a strategy for user input to be stubbed.
func Setup(read reader) error {
	if read == nil {
		return errNilReader
	}
	_game = newGame()
	printLogo(_game.logo)

	_game.setReader(read)
	_game.setPlayers(_game.chooseMarks())
	_game.print()

	return nil
}

// Read gets players's input and returns it as a text.
// It's exposed as a default impl of the `reader` Strategy.
func Read() string {
	// NOTE: it's easier to create it in place on demand vs. to store
	// and to initialize it somewhere. The `NewScanner` is very cheap inside actually
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	return strings.TrimSpace(s.Text())

	// TODO: have to check and propagate _scanner.Err() ?
}

// Game Loop()

// Loop prompts players to take turns.
// The `board` is returned for tests only.
func Loop() (board, bool, error) {
	if _game == nil || !_game.isReady() {
		return board{}, false, errCouldNotStart
	}
	more := _game.turn(_game.player1)
	if !more {
		return _game.board, false, nil
	}
	more = _game.turn(_game.player2)
	return _game.board, more, nil
}

// Private

func (g *game) turn(plr player) bool {
	c := g.inputLoop(plr)
	g.board.setCell(c, plr.mark)
	g.printBoard()

	if g.board.isWinner(plr.mark) {
		fmt.Printf("%v won!\n", plr)
		return false
	}
	if !g.board.hasEmpty() {
		fmt.Println("Draw!")
		return false
	}
	return true
}

func (g *game) inputLoop(them player) cell {
	prompt(them)
	for {
		turn := key(g.read())
		if !turn.isKey() {
			g.printBoard()
			prompt(them)

			continue
		}
		cel := turn.toCell()
		if g.board.isFilled(cel) {
			g.printBoard()
			prompt(them)

			continue
		}
		return cel
	}
}
