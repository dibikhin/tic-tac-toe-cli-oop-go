package internal

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
			{"X", __, "X"},
			{"O", "X", "O"},
			{"X", __, "O"},
		}, cell{0, 0}, true},
		{"empty", board{
			{"X", __, "X"},
			{"O", __, "O"},
			{"X", __, "O"},
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
			{"X", __, "X"},
			{"O", __, "O"},
			{"X", __, "O"},
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
			{"O", __, __},
			{"O", __, __},
		}, "X", true},
		{"last col, O", board{
			{"X", "X", "O"},
			{__, __, "O"},
			{__, __, "O"},
		}, "O", true},
		{"left diagonal, O", board{
			{"X", "X", "O"},
			{__, "O", __},
			{"O", __, __},
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
				{__, __, __},
				{__, __, __},
				{__, __, __},
			},
			args{key("5").toCell(), "X"},
			&board{
				{__, __, __},
				{__, "X", __},
				{__, __, __},
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
