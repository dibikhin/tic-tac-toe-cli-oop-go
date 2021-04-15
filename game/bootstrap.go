package game

// PlayGame starts the game by setting it up and running the game loop.
// It's a default bootstrapper.
// Other public functions are exposed for testing purposes.
func PlayGame() error {
	err := Setup(Read)
	if err != nil {
		return err
	}
	_, more, err := Loop()
	if err != nil {
		return err
	}
	for more {
		_, more, _ = Loop()
	}
	return nil
}
