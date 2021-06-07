package internal

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Emulating importing the package itself (self-importing is prohibited as "import cycle not allowed in test")
type _Board = Board

var (
	_Setup = Setup
	_Loop  = Loop
)

// A blackbox test. It uses public (exported) members of the package only.
// It's here to simplify coverage computation.
func TestLoop(t *testing.T) {
	// NOTE: intentionally kept dirty to lower maintenance

	// WARN: editing this can hang up this test
	c := -2
	// -2 is ignored;
	// -1 is for testing wrong input;
	// 0 is for choosing player; 1..7 are for players turns
	reader := func() string {
		c++
		x0 := strconv.Itoa(c)
		fmt.Println(x0)
		return x0
	}
	tests := []struct {
		name  string
		board _Board
		more  bool
	}{
		{"O: 1, X: 2",
			_Board{
				{"O", "X", "_"},
				{"_", "_", "_"},
				{"_", "_", "_"},
			},
			true},
		{"O: 3, X: 4",
			_Board{
				{"O", "X", "O"},
				{"X", "_", "_"},
				{"_", "_", "_"},
			},
			true},
		{"O: 5, X: 6",
			_Board{
				{"O", "X", "O"},
				{"X", "O", "X"},
				{"_", "_", "_"},
			},
			true},
		{"O: 7",
			_Board{
				{"O", "X", "O"},
				{"X", "O", "X"},
				{"O", "_", "_"},
			},
			false},
	}

	err := _Setup(reader) // NOTE: setting up is mandatory
	if err != nil {
		t.Errorf("Error = %v, want nil", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, _ := _Loop()
			// assert.Equal is for verbose output
			if !assert.Equal(t, tt.board, got) {
				t.Errorf("Loop() got = %v, want %v", got, tt.board)
			}
			if got1 != tt.more {
				t.Errorf("Loop() got1 = %v, want %v", got1, tt.more)
			}
		})
	}
}