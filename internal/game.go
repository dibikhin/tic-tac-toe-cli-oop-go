package internal

import (
	"fmt"
	"strings"
)

// Game

type game struct {
	logo board

	board   board
	player1 player
	player2 player

	read reader
}

// Constants

var _deadGame = &game{board: _deadBoard}

// Public

// Pure
func (g game) Board() board {
	return g.board
}

// Private

// Pure
func newGame() *game {
	return &game{
		logo:  _logo,
		board: _blankBoard,

		// others are omitted for flexibility
	}
}

// Pure
func (g game) isReady() bool {
	return g.read != nil &&
		!g.player1.isEmpty() &&
		!g.player2.isEmpty() &&
		!g.board.isEmpty()
}

func (g *game) setReader(read reader) {
	// WARN: possible nils
	g.read = read
}

func (g *game) setPlayers(p1, p2 player) {
	// WARN: possible nil
	g.player1 = p1
	g.player2 = p2
}

// Setup() IO

func printLogo(s fmt.Stringer) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()

	fmt.Println("(Use `ctrl+c` to exit)")
	fmt.Println()
}

func (g game) chooseMarks() (player, player) {
	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	// WARN: possible nil, ignored
	mark1 := g.read()
	p1, p2 := arrangePlayers(mark1)
	return p1, p2
}

func (g game) print() {
	fmt.Println()

	fmt.Println(g.player1)
	fmt.Println(g.player2)

	g.printBoard()
}

// Other

// Pure
func arrangePlayers(m mark) (player, player) {
	if strings.ToLower(m) == "x" {
		return player{"X", 1}, player{"O", 2}
	}
	return player{"O", 1}, player{"X", 2}
}
