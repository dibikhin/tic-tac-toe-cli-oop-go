package game

import "strconv"

// Cell

type cell struct {
	row, col int
}

// Key

type key string

func (k key) isKey() bool {
	n, err := strconv.Atoi(string(k))
	if err != nil {
		return false
	}
	return n >= 1 && n <= 9
}

func (k key) toCell() cell {
	coords := map[key]cell{
		"1": {0, 0}, "2": {0, 1}, "3": {0, 2},
		"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
		"7": {2, 0}, "8": {2, 1}, "9": {2, 2},
	}
	return coords[k] // TODO: detect and propagate errors?
}
