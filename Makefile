.PHONY: all test build start run clean

ttt = tictactoe.bin

all: clean test build

test:
	@echo "\nTesting..."
	go test -v ./...

build:
	@echo "\nBuilding..."
	@go version
	go build -o ${ttt}

start: ${ttt}
	@echo "\nStarting..."
	./${ttt}

run:
	@echo "\nRunning..."
	clear && go run main.go

clean:
	@echo "\nCleaning up..."
	rm -f ${ttt}