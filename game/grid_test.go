package game

import (
	"reflect"
	"testing"
)

func Test_grid_isFilled(t *testing.T) {
	tests := []struct {
		name string
		b    grid
		args cell
		want bool
	}{
		{"filled", grid{
			{"X", "_", "X"},
			{"O", "X", "O"},
			{"X", "_", "O"},
		}, cell{0, 0}, true},
		{"empty", grid{
			{"X", "_", "X"},
			{"O", "_", "O"},
			{"X", "_", "O"},
		}, cell{0, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isFilled(tt.args); got != tt.want {
				t.Errorf("grid.isFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_hasEmpty(t *testing.T) {
	tests := []struct {
		name string
		b    grid
		want bool
	}{
		{"has empty", grid{
			{"X", "_", "X"},
			{"O", "_", "O"},
			{"X", "_", "O"},
		}, true},
		{"all filled", grid{
			{"X", "O", "X"},
			{"O", "X", "O"},
			{"O", "X", "O"},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.hasEmpty(); got != tt.want {
				t.Errorf("grid.hasEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_isWinner(t *testing.T) {
	tests := []struct {
		name string
		b    grid
		arg  string
		want bool
	}{
		{"first row, X", grid{
			{"X", "X", "X"},
			{"O", "_", "_"},
			{"O", "_", "_"},
		}, "X", true},
		{"last col, O", grid{
			{"X", "X", "O"},
			{"_", "_", "O"},
			{"_", "_", "O"},
		}, "O", true},
		{"left diagonal, O", grid{
			{"X", "X", "O"},
			{"_", "O", "_"},
			{"O", "_", "_"},
		}, "O", true},
		{"draw", grid{
			{"X", "O", "O"},
			{"O", "X", "X"},
			{"O", "X", "O"},
		}, "O", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isWinner(tt.arg); got != tt.want {
				t.Errorf("grid.isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_setCell(t *testing.T) {
	type args struct {
		c      cell
		player string
	}
	tests := []struct {
		name string
		b    *grid
		args args
		want *grid
	}{
		{
			"1,1",
			&grid{
				{"_", "_", "_"},
				{"_", "_", "_"},
				{"_", "_", "_"},
			},
			args{toCell("5"), "X"},
			&grid{
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
				t.Errorf("grid.setCell() = %v, want %v", tt.b, tt.want)
			}
		})
	}
}
