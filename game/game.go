package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type game struct {
	isReady bool

	logo board

	board   board
	player1 player
	player2 player

	scanner *bufio.Scanner
	reader  reader
}

// Pure
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

// Pure
func arrange(m mark) (player, player) {
	if strings.ToLower(m) == "x" {
		return player{"X", 1}, player{"O", 2}
	}
	return player{"O", 1}, player{"X", 2}
}
