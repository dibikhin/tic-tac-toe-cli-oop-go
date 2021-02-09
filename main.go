package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hey! This is Tic Tac Toe :)")
	fmt.Println()

	logo := [][]string{
		{"X", " ", "X"},
		{"O", "X", "O"},
		{"X", " ", "O"},
	}
	for _, r := range logo {
		fmt.Printf("%s\n", strings.Join(r, " "))
	}
	fmt.Println()

	fmt.Print("Press 'x' or 'o' to choose mark for Player 1: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	player1 := scanner.Text()

	fmt.Println("Player 1 is:", strings.ToUpper(player1))

	var player2 string
	if player1 == "x" {
		player2 = "o"
	} else {
		player2 = "x"
	}
	fmt.Println("Player 2 is:", strings.ToUpper(player2))

	fmt.Println("Press from 1 to 9 to put your mark (e.g. 5 is center).")

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	for _, r := range board {
		fmt.Printf("%s\n", strings.Join(r, " "))
	}
	fmt.Println()
	fmt.Print("Player 1, your turn: ")

	scanner.Scan()
	p1turn := scanner.Text()
	if p1turn != "" {
		board[1][1] = "X"

		fmt.Println()
		for _, r := range board {
			fmt.Printf("%s\n", strings.Join(r, " "))
		}
	}
	fmt.Println()
	fmt.Print("Player 2, your turn: ")

	scanner.Scan()
	p2turn := scanner.Text()
	if p2turn != "" {
		board[2][2] = "O"

		fmt.Println()
		for _, r := range board {
			fmt.Printf("%s\n", strings.Join(r, " "))
		}
	}
	fmt.Println()

	// for {
	// 	fmt.Print("Enter text: ")
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	text := scanner.Text()
	// 	fmt.Println(text)
	// }
}
