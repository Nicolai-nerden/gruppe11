package tictactoe

import (
	"fmt"
	"gruppe11/mp02/validInputs"
	"math/rand"
	"strconv"
	"time"
)

// Player er et struct som lagrer verdiene til spillere i tictactoe og turnering.
type Player struct {
	Name     string
	Score    int
	TimeUsed int
}

var board map[int]string
var runde int
var winner Player

// PlayGame starter Tic Tac Toe. Tar inn to spillere, Returnerer vinneren.
func PlayGame(p1 Player, p2 Player, simulation bool) Player {

	if simulation { //hvis turneringen er i simuleringsmodus
		winner = simulateGame(p1, p2)
		return winner
	}

	runde = 1
	board = map[int]string{ // Lagrer spillets trekk. Tomme felt forblir tall som representerer posisjonen dens på brettet.
		1: "1", 2: "2", 3: "3",
		4: "4", 5: "5", 6: "6",
		7: "7", 8: "8", 9: "9"}

	printBoard()
	fmt.Println("\n" + p1.Name + " Starter.")
	newRoundOrGameOver(p1, p2)
	return winner
}

// simulateGame simulerer et spill, 50/50 hvem som vinner
func simulateGame(p1 Player, p2 Player) Player {
	rand.Seed(time.Now().UnixNano()) //forandrer seedet etter hva tiden er.
	var rng int = rand.Intn(2)
	var winner Player

	if rng == 0 {
		winner = p1
	} else {
		winner = p2
	}
	return winner
}

// printBoard printer ut hvordan brettet ser ut med eventuelle trekk i terminalen.
func printBoard() {

	line := " ------------- "
	wall := " | "
	shiftLine := 0

	fmt.Println("\n" + line)

	for g := 0; g < 3; g++ { // Printer ut brettet, en linje om gangen.
		for i := 1; i <= 3; i++ {
			fmt.Print(wall, board[i+shiftLine])
		}
		fmt.Println(wall)
		shiftLine += 3
		fmt.Println(line)
	}

}

// place Move plasserer trekk på brettet.
func placeMove(p1 Player, p2 Player) (Player, Player) {
	start := time.Now() // Tar tidspunktet siden fra når trekket starter.
	move := checkIfTaken()
	timeUsed := time.Since(start).Milliseconds() // Regner ut tiden som har gått siden trekket startet.

	i := 0
	for board[i] != strconv.Itoa(move) {
		i++
	}
	if runde%2 == 0 {
		board[i] = "O"
		p2.TimeUsed += int(timeUsed)
	} else {
		board[i] = "X"
		p1.TimeUsed += int(timeUsed)
	}

	return p1, p2
}

func checkIfTaken() int {
Again:
	move := validInputs.CheckIfValid()
	if board[move] == "X" || board[move] == "O" {
		fmt.Println("Dette trekket er allerede tatt. Prøv igjen.")
		goto Again
	}
	return move
}

// newRoundOrGameOver sjekker om det er en vinner, uavgjort eller om den skal starte en ny runde.
func newRoundOrGameOver(p1 Player, p2 Player) {
newRound:
	winningCombos := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 4, 7}, {2, 5, 8}, {3, 6, 9}, {1, 5, 9}, {3, 5, 7}}
	fmt.Println("\nSkriv inn tallet som representerer posisjonen du vil sette brikken din på")
	p1, p2 = placeMove(p1, p2) // plasserer brikken i på brettet og oppdaterer hver spiller har brukt.
	printBoard()

	for i := 0; i < len(winningCombos); i++ { //Sjekker om det er noen vinner kombinasjoner på brettet.
		if board[winningCombos[i][0]] == board[winningCombos[i][1]] && board[winningCombos[i][1]] == board[winningCombos[i][2]] {
			if board[winningCombos[i][0]] == "X" {
				winner = p1
				return
			}
			winner = p2
			return

		}
	}

	runde++
	if runde > 9 { //Sjekker om brettet er fylt opp.
		fmt.Println("\nBrettet er fullt. Vinneren blir dermed avgjort på tid.")
		if p1.TimeUsed <= p2.TimeUsed {
			winner = p1
		} else {
			winner = p2
		}

		fmt.Println("\n "+p1.Name+" brukte", p1.TimeUsed, "millisekunder.")
		fmt.Println(" "+p2.Name+" brukte", p2.TimeUsed, "millisekunder.")

		return
	}

	goto newRound
}
