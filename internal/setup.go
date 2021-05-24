package internal

import (
	"bufio"
	"os"
	"strings"
)

// Setup initializes the game and helps players to choose marks.
// The param is a strategy for user input to be stubbed.
// One can pass nothing, the default reader is used in the case.
func Setup(rs ...reader) error {
	var err error
	_game, err = makeGame(DefaultReader, rs...)
	if err != nil {
		return err
	}
	printLogo(_game.logo)

	_game.setPlayers(_game.chooseMarks())
	_game.print()

	return nil
}

// DefaultReader gets player's input and returns it as a text.
// It's exposed as a default impl of the `reader` Strategy.
func DefaultReader() string {
	// NOTE: it's easier to create it in place on demand vs. to store
	// and to initialize it somewhere. The `NewScanner` is very cheap inside actually
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	return strings.TrimSpace(s.Text())

	// TODO: have to check and propagate _scanner.Err() ?
}

// Private

// Factory
func makeGame(def reader, rs ...reader) (*game, error) {
	gam := newGame()
	gam.setReader(def)

	if len(rs) > 0 {
		fst := rs[0]
		if fst == nil {
			return nil, errNilReader
		}
		gam.setReader(fst)
	}
	return gam, nil
}
