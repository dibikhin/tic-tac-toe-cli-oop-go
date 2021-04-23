package game

import (
	"fmt"
	"strings"
)

// User input strategy for stubbing in tests.
//
// NOTE: An interface is more idiomatic in this case. BUT it's overkill to define
// a type with constructor, an interface and its fake implementation in tests vs. this
// func, its impl and its fake impl in tests.
type reader func() string

// Game

type game struct {
	logo board

	board   board
	player1 player
	player2 player

	read reader
}

// Pure
func newGame() *game {
	return &game{
		logo: board{
			{"X", " ", "X"},
			{"O", "X", "O"},
			{"X", " ", "O"}},
		board: board{
			{_blank, _blank, _blank},
			{_blank, _blank, _blank},
			{_blank, _blank, _blank}},
	}
}

// Pure
func (g game) isReady() bool {
	return g.read != nil &&
		g.player1 != player{} &&
		g.player2 != player{} &&
		g.board != board{}
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

	// WARN: possible nil
	mark1 := g.read()
	p1, p2 := arrange(mark1)
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
func arrange(m mark) (player, player) {
	if strings.ToLower(m) == "x" {
		return player{"X", 1}, player{"O", 2}
	}
	return player{"O", 1}, player{"X", 2}
}
