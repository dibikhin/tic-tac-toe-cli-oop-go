package game

import (
	"fmt"
	"strings"
)

// Cell

type cell struct {
	row, col int
}

func toCell(key string) cell {
	m := map[string]cell{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
	return m[key] // TODO: detect and propagate errors?
}

// Board

const _blank = "_"

type mark = string // to avoid conversions

type board [3][3]mark

func (b board) String() string {
	var rows []mark
	for _, row := range b {
		s := strings.Join(row[:], " ")
		rows = append(rows, s)
	}
	return strings.Join(rows, "\n")
}

func (b *board) setCell(c cell, m mark) {
	b[c.row][c.col] = m
}

func (b board) isFilled(c cell) bool {
	v := b[c.row][c.col]
	return v != _blank
}

func (b board) hasEmpty() bool {
	for _, row := range b {
		for _, v := range row {
			if v == _blank {
				return true
			}
		}
	}
	return false
}

func (b board) isWinner(m mark) bool {
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

func (b board) print() {
	var _ fmt.Stringer = board{}

	fmt.Println()
	fmt.Println()
	fmt.Println("Press 1 to 9 to mark an empty cell (5 is center), then press ENTER. Board:")
	fmt.Println()

	fmt.Println(b)
	fmt.Println()
}
