package game

import (
	. "tictactoe/internal"
)

// Public

// Play starts the game by setting it up and running the game loop.
// It's a default bootstrapper.
// Other public functions are exposed for testing purposes.
func Play() error {
	err := Setup()
	if err != nil {
		return err
	}
	more := true
	for more {
		_, more, err = Loop()
		if err != nil {
			return err
		}
	}
	return nil
}
