package game

import (
	"testing"
)

func Test_arrange(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		want  string
		want1 string
	}{
		{"empty", "", "O", "X"},
		{"space", "    ", "O", "X"},
		{"junk", "^G$FDTeg39dslf*^58)#", "O", "X"},
		{"small x", "x", "X", "O"},
		{"big X", "X", "X", "O"},
		{"small o", "o", "O", "X"},
		{"big O", "O", "O", "X"},
		{"zero", "0", "O", "X"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := arrange(tt.arg)
			if got != tt.want {
				t.Errorf("arrange() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("arrange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isKey(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"empty", "", false},
		{"space", "  ", false},
		{"too big", "42", false},
		{"too small", "0", false},
		{"letter", "a", false},
		{"junk", " bsd &% #&#", false},
		{"number 1 to 9 incl.", "5", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isKey(tt.arg); got != tt.want {
				t.Errorf("isKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pos(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		want  int
		want1 int
	}{
		{"out of range", "-1", 0, 0},
		{"1", "1", 0, 0},
		{"5", "5", 1, 1},
		{"8", "8", 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := pos(tt.arg)
			if got != tt.want {
				t.Errorf("pos() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("pos() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_grid_isFilled(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		b    grid
		args args
		want bool
	}{
		{"filled", grid{
			{"X", "_", "X"},
			{"O", "X", "O"},
			{"X", "_", "O"},
		}, args{0, 0}, true},
		{"empty", grid{
			{"X", "_", "X"},
			{"O", "_", "O"},
			{"X", "_", "O"},
		}, args{0, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isFilled(tt.args.row, tt.args.col); got != tt.want {
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
		{"nobody", grid{
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
