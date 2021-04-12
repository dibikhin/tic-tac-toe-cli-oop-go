# Tic-tac-toe

Classic 3x3 Tic-tac-toe for 2 friends to play locally in terminal.

[![GoReportCard example](https://goreportcard.com/badge/github.com/dibikhin/tic-tac-toe-go)](https://goreportcard.com/report/github.com/dibikhin/tic-tac-toe-go) [![Maintainability](https://api.codeclimate.com/v1/badges/229dc45729c3983e99a9/maintainability)](https://codeclimate.com/github/dibikhin/tic-tac-toe-go/maintainability) [![example branch parameter](https://github.com/dibikhin/tic-tac-toe-go/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/dibikhin/tic-tac-toe-go/actions/workflows/go.yml)

# Overview

Get someone and play locally in your terminal using keyboard only. Cannot play w/ computer yet so can play with yourself at worst :)

## Getting Started

### Prerequisites
- Install [Go](https://golang.org/doc/install) (tested on go1.15.7 linux/amd64)

### Installing
Not needed, runs as is, just clone:
```
$ cd my_projects
$ git clone https://github.com/dibikhin/tic-tac-toe-go.git
```

## Running the tests
```
$ cd tic-tac-toe-go
$ cd game
$ go test -v
...
>PASS
>ok      tictactoe/game  0.010s
```

## Running
```
$ cd tic-tac-toe-go
$ clear && go run main.go
> Hey! This is 3x3 Tic-tac-toe for 2 friends :)
>
> X   X
> O X O
> X   O
>
> Press 'x' or 'o' to choose mark for Player 1:
```

NOTE: Hit `ctrl+c` to exit.

## Internals
- The 3x3 size is hardcoded
- The UI is CLI
- No timeouts for turns
- Dirty input tolerant

### Technical
- The game is one active app (no client/server)
- A basic DI is inside
  - basic IoC in `main.go` — the `Read()` strategy
  - no mocks
  - inner DI in `loop.go`
- Well-tested
- ~90% code coverage
- IO extracted but not isolated

## Project Structure
- `/game` — The game package
- `|-board.go`
- `|-game.go`
- `|-key.go`
- `|-loop.go`
- `|-player.go`
- `main.go` — Entry point

## Authors
- [Roman Dibikhin](https://github.com/dibikhin)

## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments
Thanks to:
- [A Tour of Go](https://tour.golang.org/welcome/1) — For the idea
- [Tic-tac-toe](https://en.wikipedia.org/wiki/Tic-tac-toe) — A lot of insights about the game
