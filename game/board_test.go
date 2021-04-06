package game

import (
	"reflect"
	"testing"
)

func Test_board_isFilled(t *testing.T) {
	tests := []struct {
		name string
		b    board
		args cell
		want bool
	}{
		{"filled", board{
			{"X", "_", "X"},
			{"O", "X", "O"},
			{"X", "_", "O"},
		}, cell{0, 0}, true},
		{"empty", board{
			{"X", "_", "X"},
			{"O", "_", "O"},
			{"X", "_", "O"},
		}, cell{0, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isFilled(tt.args); got != tt.want {
				t.Errorf("board.isFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_hasEmpty(t *testing.T) {
	tests := []struct {
		name string
		b    board
		want bool
	}{
		{"has empty", board{
			{"X", "_", "X"},
			{"O", "_", "O"},
			{"X", "_", "O"},
		}, true},
		{"all filled", board{
			{"X", "O", "X"},
			{"O", "X", "O"},
			{"O", "X", "O"},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.hasEmpty(); got != tt.want {
				t.Errorf("board.hasEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_isWinner(t *testing.T) {
	tests := []struct {
		name string
		b    board
		arg  string
		want bool
	}{
		{"first row, X", board{
			{"X", "X", "X"},
			{"O", "_", "_"},
			{"O", "_", "_"},
		}, "X", true},
		{"last col, O", board{
			{"X", "X", "O"},
			{"_", "_", "O"},
			{"_", "_", "O"},
		}, "O", true},
		{"left diagonal, O", board{
			{"X", "X", "O"},
			{"_", "O", "_"},
			{"O", "_", "_"},
		}, "O", true},
		{"draw", board{
			{"X", "O", "O"},
			{"O", "X", "X"},
			{"O", "X", "O"},
		}, "O", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isWinner(tt.arg); got != tt.want {
				t.Errorf("board.isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_setCell(t *testing.T) {
	type args struct {
		c      cell
		player string
	}
	tests := []struct {
		name string
		b    *board
		args args
		want *board
	}{
		{
			"1,1",
			&board{
				{"_", "_", "_"},
				{"_", "_", "_"},
				{"_", "_", "_"},
			},
			args{key("5").toCell(), "X"},
			&board{
				{"_", "_", "_"},
				{"_", "X", "_"},
				{"_", "_", "_"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.setCell(tt.args.c, tt.args.player)
			if !reflect.DeepEqual(tt.b, tt.want) {
				t.Errorf("board.setCell() = %v, want %v", tt.b, tt.want)
			}
		})
	}
}
