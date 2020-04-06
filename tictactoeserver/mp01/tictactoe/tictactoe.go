package tictactoe

import (
	"tx3server/mp01/clientcommunication"
	"tx3server/mp02/validInputs"
	"math/rand"
	"net"
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
func PlayGame(c net.Conn, p1 Player, p2 Player, simulation bool) Player {

	if simulation { //hvis turneringen er i simuleringsmodus
		winner = simulateGame(p1, p2)
		return winner
	}

	runde = 1
	board = map[int]string{ // Lagrer spillets trekk. Tomme felt forblir tall som representerer posisjonen dens på brettet.
		1: "1", 2: "2", 3: "3",
		4: "4", 5: "5", 6: "6",
		7: "7", 8: "8", 9: "9"}

	printBoard(c)
	clientcommunication.ClientPrintln(c, "\n"+p1.Name+" Starter.")
	newRoundOrGameOver(c, p1, p2)
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
func printBoard(c net.Conn) {

	line := " ------------- "
	wall := " | "
	shiftLine := 0

	clientcommunication.ClientPrintln(c, "\n"+line)

	for g := 0; g < 3; g++ { // Printer ut brettet, en linje om gangen.
		for i := 1; i <= 3; i++ {
			clientcommunication.ClientPrint(c, wall+board[i+shiftLine])
		}
		clientcommunication.ClientPrintln(c, wall)
		shiftLine += 3
		clientcommunication.ClientPrintln(c, line)
	}

}

// place Move plasserer trekk på brettet.
func placeMove(c net.Conn, p1 Player, p2 Player) (Player, Player) {
	start := time.Now() // Tar tidspunktet siden fra når trekket starter.
	move := checkIfTaken(c)
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

func checkIfTaken(c net.Conn) int {
	move := ""
Again:
	clientcommunication.ClientPrintln(c, "\nSkriv inn tallet som representerer ruten du vil sette brikken din på.")
	move = clientcommunication.ClientRead(c)

	if len(move) == 0 { //hvis tom input
		clientcommunication.ClientPrintln(c, "Tomt trekk. Prøv igjen.")
		goto Again
	}
	parsedMove, valid := validInputs.CheckIfValid(move)

	if !valid {
		clientcommunication.ClientPrintln(c, "Fant ikke input i støttede skriftsrpåk. Prøv igjen.")
		goto Again
	}

	if board[parsedMove] == "X" || board[parsedMove] == "O" {
		clientcommunication.ClientPrintln(c, "Dette trekket er allerede tatt. Prøv igjen.")
		goto Again
	}
	return parsedMove
}

// newRoundOrGameOver sjekker om det er en vinner, uavgjort eller om den skal starte en ny runde.
func newRoundOrGameOver(c net.Conn, p1 Player, p2 Player) {
newRound:
	winningCombos := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 4, 7}, {2, 5, 8}, {3, 6, 9}, {1, 5, 9}, {3, 5, 7}}
	p1, p2 = placeMove(c, p1, p2) // plasserer brikken i på brettet og oppdaterer hver spiller har brukt.
	printBoard(c)

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
		clientcommunication.ClientPrintln(c, "\nBrettet er fullt. Vinneren blir dermed avgjort på tid.")
		if p1.TimeUsed <= p2.TimeUsed {
			winner = p1
		} else {
			winner = p2
		}

		clientcommunication.ClientPrintln(c, "\n "+p1.Name+" brukte "+strconv.Itoa(p1.TimeUsed)+" millisekunder.")
		clientcommunication.ClientPrintln(c, " "+p2.Name+" brukte "+strconv.Itoa(p2.TimeUsed)+" millisekunder.")

		return
	}

	goto newRound
}
