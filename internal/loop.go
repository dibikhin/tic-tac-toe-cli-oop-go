// Package internal implements 3x3 Tic-tac-toe for 2 friends (cannot play with computer yet).
// Players choose their marks, put them, the game checks winner or draw.
package internal

import (
	"errors"
	"fmt"
)

// Constants, Private
var (
	// errCouldNotStart arises when `Loop()` is run without running `Setup()` first.
	errCouldNotStart = errors.New("game: couldn't start the game loop, set up the game first")
)

// Public

// Game Loop()

type again = bool

// Loop prompts players to take turns.
// The `board` is returned for tests only.
func Loop(gam *game) (*game, again, error) {
	if gam == nil || !gam.isReady() {
		return _deadGame, false, errCouldNotStart
	}
	more := gam.turn(gam.player1)
	if !more {
		return gam, false, nil
	}
	more = gam.turn(gam.player2)
	return gam, more, nil
}

// Private

func (g *game) turn(playr player) bool {
	c := g.inputLoop(playr)
	g.board.setCell(c, playr.mark)
	g.printBoard()

	if g.board.isWinner(playr.mark) {
		fmt.Printf("%v won!\n", playr)
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
