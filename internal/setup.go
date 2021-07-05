package internal

// Setup initializes the game and helps players to choose marks.
// The param is a strategy for user input to be stubbed.
// One can pass nothing, the default reader is used in the case.
// Example:
// ctx, err := Setup()
// OR
// ctx, err := Setup(DefaultReader)
// OR
// ctx, err := Setup(yourReaderFunc)
func Setup(rs ...reader) (*game, error) {
	alt, err := extractReader(rs)
	if err != nil {
		return nil, err
	}
	gam := makeGame(DefaultReader, alt)

	printLogo(gam.logo)

	gam.setPlayers(gam.chooseMarks())
	gam.print()

	return gam, nil
}

// Factory, Pure
func makeGame(def, alt reader) *game {
	gam := newGame()
	if alt != nil {
		gam.setReader(alt)
		return gam
	}
	gam.setReader(def)
	return gam
}
