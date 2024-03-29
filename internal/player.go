package internal

import "fmt"

type player struct {
	mark mark
	num  int
}

func (p player) String() string {
	return fmt.Sprintf(`Player %v ("%v")`, p.num, p.mark)
}

func (p player) isEmpty() bool {
	return p == player{}
}

// IO

// Implicit check for `fmt.Stringer` impl
func prompt(s fmt.Stringer) {
	fmt.Printf("%v, your turn: ", s)
}

func (g game) printBoard() {
	// Explicit check for the interface
	var _ fmt.Stringer = g.board

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(g.board)
	fmt.Println()
}
