package tictactoe

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//Player er et struct som lagrer verdiene til spillere i tictactoe og turnering.
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

func simulateGame(p1 Player, p2 Player) Player { //Simulerer et spill, 50/50 hvem som vinner
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

func printBoard() { // Printer ut brettet i terminalen.

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

func placeMove(p1 Player, p2 Player) (Player, Player) {
	start := time.Now() // Tar tidspunktet siden fra når trekket starter.
	var move = moveAndValidate()
	timeUsed := time.Since(start).Milliseconds() // Regner ut tiden som har gått siden trekket startet.

	i := 0
	for board[i] != move {
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

func moveAndValidate() string {
	// Scanner inn input, sjekker om første tegnet er et tall mellom 1-9 i bytes. Starter på nytt hvis ikke.
	// Hvis det er det returnerer den dette tegnet tilbake som en string.
Start:
	var move string
	validInputs := [][]byte{{49}, {50}, {51}, {52}, {53}, {54}, {55}, {56}, {57}} //gyldige svar i bytes
	i := 0

	fmt.Scanln(&move)
	if move == "" { //hvis tom input
		fmt.Println("Tomt trekk. Prøv igjen.")
		goto Start
	}
	for (i < len(validInputs)) && ([]byte(move)[0] != validInputs[i][0]) { //Sjekker om inputen er et registrert gyldig trekk.
		i++
	}

	if i >= len(validInputs) { //Hvis i er større enn limiten til loopen betyr dette at den ikke fant en samsvarende byte.
		fmt.Println("Ugyldig trekk. Prøv igjen.")
		goto Start
	}

	byteToInt, _ := strconv.Atoi(string(validInputs[i]))

	if board[byteToInt] == "X" || board[byteToInt] == "O" {
		fmt.Println("Dette trekket er allerede tatt. Prøv igjen.")
		goto Start
	}

	return string(validInputs[i])
}

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
