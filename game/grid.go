package game

import (
	// "fmt"
	"strings"
)

// Grid

type grid [3][3]mark

type cell struct {
	row, col int
}

func (b grid) String() string {
	var rows []mark
	for _, row := range b {
		s := strings.Join(row[:], " ")
		rows = append(rows, s)
	}
	return strings.Join(rows, "\n")
}

func (b *grid) setCell(c cell, m mark) {
	b[c.row][c.col] = m
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

func (b grid) isWinner(m mark) bool {
	// Something better needed, too naive

	// Horizontal
	h0 := b[0][0] == m && b[0][1] == m && b[0][2] == m
	h1 := b[1][0] == m && b[1][1] == m && b[1][2] == m
	h2 := b[2][0] == m && b[2][1] == m && b[2][2] == m

	// Vertical
	v0 := b[0][0] == m && b[1][0] == m && b[2][0] == m
	v1 := b[0][1] == m && b[1][1] == m && b[2][1] == m
	v2 := b[0][2] == m && b[1][2] == m && b[2][2] == m

	// Diagonal
	d0 := b[0][0] == m && b[1][1] == m && b[2][2] == m
	d1 := b[0][2] == m && b[1][1] == m && b[2][0] == m

	return h0 || h1 || h2 || v0 || v1 || v2 || d0 || d1
}

// IO

func (b grid) print() {
	println("print")
	// fmt.Println()
	// fmt.Println()
	// fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	// fmt.Println()

	// fmt.Println(b)
	// fmt.Println()
}
