package game

import "fmt"

type player struct {
	mark mark
	num  int
}

func (p player) String() string {
	return fmt.Sprintf(`Player %v ("%v")`, p.num, p.mark)
}

// IO

// Implicit check for `fmt.Stringer` impl
func prompt(s fmt.Stringer) {
	fmt.Printf("%v, your turn: ", s)
}

func (g game) printBoard() {
	// WARN: possible nil
	var _ fmt.Stringer = g.board // Explicit check for fmt.Stringer

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(g.board)
	fmt.Println()
}
