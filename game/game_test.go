package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_arrange(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		want1 player
		want2 player
	}{
		{"empty", "", player{"O", 1}, player{"X", 2}},
		{"space", "    ", player{"O", 1}, player{"X", 2}},
		{"junk", "^G$FDTeg39dslf*^58)#", player{"O", 1}, player{"X", 2}},
		{"small x", "x", player{"X", 1}, player{"O", 2}},
		{"big X", "X", player{"X", 1}, player{"O", 2}},
		{"small o", "o", player{"O", 1}, player{"X", 2}},
		{"big O", "O", player{"O", 1}, player{"X", 2}},
		{"zero", "0", player{"O", 1}, player{"X", 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := arrange(tt.arg)
			if got1 != tt.want1 {
				t.Errorf("arrange() got = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("arrange() got1 = %v, want %v", got2, tt.want2)
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

func Test_toCell(t *testing.T) {
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
			c := toCell(tt.arg)
			if c.row != tt.want {
				t.Errorf("pos() got = %v, want %v", c.row, tt.want)
			}
			if c.col != tt.want1 {
				t.Errorf("pos() got1 = %v, want %v", c.col, tt.want1)
			}
		})
	}
}

func TestLoop(t *testing.T) {
	// NOTE: intentionally kept dirty to lower maintenance

	f := true
	tests := []struct {
		name  string
		read  reader
		want  grid
		want1 bool
	}{
		{"O: 1, X: 2",
			func() string {
				if f {
					f = !f
					return "1"
				}
				f = !f
				return "2"
			},
			grid{
				{"O", "X", "_"},
				{"_", "_", "_"},
				{"_", "_", "_"},
			},
			true},

		{"O: 3, X: 4",
			func() string {
				if f {
					f = !f
					return "3"
				}
				f = !f
				return "4"
			},
			grid{
				{"O", "X", "O"},
				{"X", "_", "_"},
				{"_", "_", "_"},
			},
			true},
		{"O: 5, X: 6",
			func() string {
				if f {
					f = !f
					return "5"
				}
				f = !f
				return "6"
			},
			grid{
				{"O", "X", "O"},
				{"X", "O", "X"},
				{"_", "_", "_"},
			},
			true},
		{"O: 7",
			func() string { return "7" },
			grid{
				{"O", "X", "O"},
				{"X", "O", "X"},
				{"O", "_", "_"},
			},
			false},
	}

	Setup(func() string { return "o" }) // NOTE

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Loop(tt.read)
			// assert.Equal is for verbose output
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Loop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Loop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
