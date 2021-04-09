package game

import (
	"strings"
)

type (
	mark  = string // to avoid conversions
	board [3][3]mark
)

const _blank = "_"

func (b board) String() string {
	var dump []string
	for _, row := range b {
		s := strings.Join(row[:], " ")
		dump = append(dump, s)
	}
	return strings.Join(dump, "\n")
}

// Private

func (b *board) setCell(c cell, m mark) {
	b[c.row][c.col] = m
}

// Pure
func (b board) isFilled(c cell) bool {
	return b[c.row][c.col] != _blank
}

// Pure
func (b board) hasEmpty() bool {
	for _, row := range b {
		for _, m := range row {
			if m == _blank {
				return true
			}
		}
	}
	return false
}

// Pure
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

// No IO allowed in this file for SRP
