package game_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	game "tictactoe/game"
)

type Board = game.Board

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
		board Board
		more  bool
	}{
		{"O: 1, X: 2",
			Board{
				{"O", "X", "_"},
				{"_", "_", "_"},
				{"_", "_", "_"},
			},
			true},
		{"O: 3, X: 4",
			Board{
				{"O", "X", "O"},
				{"X", "_", "_"},
				{"_", "_", "_"},
			},
			true},
		{"O: 5, X: 6",
			Board{
				{"O", "X", "O"},
				{"X", "O", "X"},
				{"_", "_", "_"},
			},
			true},
		{"O: 7",
			Board{
				{"O", "X", "O"},
				{"X", "O", "X"},
				{"O", "_", "_"},
			},
			false},
	}

	err := game.Setup(reader) // NOTE: setting up is mandatory
	if err != nil {
		t.Errorf("Error = %v, want nil", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, _ := game.Loop()
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
