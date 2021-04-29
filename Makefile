.PHONY: all test lint run build start clean

ttt = tictactoe.bin

all: clean test build

test:
	@echo "\nTesting..."
	go test -v -coverpkg=./game ./game

lint:
	golangci-lint run -v

run:
	@echo "\nRunning..."
	clear && go run main.go

build:
	@echo "\nBuilding..."
	@go version
	go build -o ${ttt}

start: ${ttt}
	@echo "\nStarting..."
	./${ttt}

clean:
	@echo "\nCleaning up..."
	rm -f ${ttt}