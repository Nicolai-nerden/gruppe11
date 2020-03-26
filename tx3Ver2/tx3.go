package main

import (
	"fmt"
)

type player struct {
	name     string
	score    int
	timeUsed int
}

var board []int
var runde int

func main() {

	p1 := player{"Nicolai", 6, 140}
	p2 := player{"Martin", 5, 300}

	playGame(p1, p2)
}

func playGame(p1 player, p2 player) {

	board = []string{1, 2, 3, 4, 6, 7, 8, 9} //lagrer spillets trekk. Tomme felt forblir tall som representerer posisjonen dens på brettet.
	runde = 1

	printBoard()
	fmt.Println("\nSkriv inn tallet som representerer posisjonen du vil sette inn din første brikke")

	placeMove()
	printBoard()

}

func printBoard() {

	line := " ------------- "
	wall := " | "
	nextLine := 0

	fmt.Println("\n" + line)

	for g := 0; g < 3; g++ { //printer ut brettet, en linje om gangen.
		for i := 0; i < 3; i++ {
			fmt.Print(wall, board[i]+nextLine)
		}
		fmt.Println(wall)
		nextLine += 3
		fmt.Println(line)
	}

}

func placeMove() {
	var move int
	fmt.Scanln(&move)

	i := 0
	for board[i] != move {
		i++
	}

	if runde%2 == 0 {
		board[i] == "O"
	}
}
