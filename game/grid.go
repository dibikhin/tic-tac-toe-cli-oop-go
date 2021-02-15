package game

import (
	"fmt"
	"strings"
)

// Grid

type grid [3][3]string

type cell struct {
	row, col int
}

func (b *grid) setCell(c cell, player string) *grid {
	b[c.row][c.col] = player
	return b
}

func toCell(key string) cell {
	m := map[string]cell{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
	return m[key] // TODO: detect and propagate errors?
}

func (b grid) isFilled(c cell) bool {
	v := b[c.row][c.col]
	return v != "_"
}

func (b grid) hasEmpty() bool {
	for _, row := range b {
		for _, v := range row {
			if v == "_" {
				return true
			}
		}
	}
	return false
}

func (b grid) isWinner(s string) bool {
	// Something better needed, too naive

	// Horizontal
	x0 := b[0][0] == s && b[0][1] == s && b[0][2] == s
	x1 := b[1][0] == s && b[1][1] == s && b[1][2] == s
	x2 := b[2][0] == s && b[2][1] == s && b[2][2] == s

	// Vertical
	x3 := b[0][0] == s && b[1][0] == s && b[2][0] == s
	x4 := b[0][1] == s && b[1][1] == s && b[2][1] == s
	x5 := b[0][2] == s && b[1][2] == s && b[2][2] == s

	// Diagonal
	x6 := b[0][0] == s && b[1][1] == s && b[2][2] == s
	x7 := b[0][2] == s && b[1][1] == s && b[2][0] == s

	return x0 || x1 || x2 || x3 || x4 || x5 || x6 || x7
}

// IO

func (b grid) print() {
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")

	_print(b)
}

func _print(b grid) {
	fmt.Println()
	for _, row := range b {
		fmt.Printf("%s\n", strings.Join(row[:], " "))
	}
	fmt.Println()
}
