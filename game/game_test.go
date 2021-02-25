package game

import (
	// "reflect"
	// "strconv"
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

// func TestLoop(t *testing.T) {
// 	type args struct {
// 		read reader
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  grid
// 		want1 bool
// 	}{
// 		{"qqq", args{rdr},
// 			grid{
// 				{"_", "_", "_"},
// 				{"_", "X", "_"},
// 				{"_", "_", "_"},
// 			},
// 			true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := Loop(tt.args.read)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Loop() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("Loop() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

// var c = 0

// func rdr() string {
// 	if c > 9 {
// 		return ""
// 	}
// 	c++
// 	return strconv.Itoa(c)
// }
