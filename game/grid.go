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

func (b grid) String() string {
	var rows []string
	for _, row := range b {
		s := strings.Join(row[:], " ")
		rows = append(rows, s)
	}
	return strings.Join(rows, "\n")
}

func (b *grid) setCell(c cell, player string) {
	b[c.row][c.col] = player
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
	h0 := b[0][0] == s && b[0][1] == s && b[0][2] == s
	h1 := b[1][0] == s && b[1][1] == s && b[1][2] == s
	h2 := b[2][0] == s && b[2][1] == s && b[2][2] == s

	// Vertical
	v0 := b[0][0] == s && b[1][0] == s && b[2][0] == s
	v1 := b[0][1] == s && b[1][1] == s && b[2][1] == s
	v2 := b[0][2] == s && b[1][2] == s && b[2][2] == s

	// Diagonal
	d0 := b[0][0] == s && b[1][1] == s && b[2][2] == s
	d1 := b[0][2] == s && b[1][1] == s && b[2][0] == s

	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
}

// IO

func (b grid) print() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(b)
	fmt.Println()
}
