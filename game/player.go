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

func prompt(s fmt.Stringer) {
	fmt.Printf("%v, your turn: ", s)
}

func (b board) print() {
	var _ fmt.Stringer = board{}

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(b)
	fmt.Println()
}
