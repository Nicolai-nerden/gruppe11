package tictactoeMultiplayer

import (
	"tx3server/selvlagdePakker/comm/multiplayerComm"
	"tx3server/selvlagdePakker/utf/validInputs"
	"net"
	"strconv"

	"time"
)

// Player er et struct som lagrer verdiene til spillere i tictactoe og turnering.
type Player struct {
	Conn     net.Conn
	Name     string
	Score    int
	TimeUsed int
}

var board map[int]string
var runde int
var winner Player
var agentList []net.Conn

// PlayGame starter Tic Tac Toe. Tar inn to spillere, Returnerer vinneren.
func PlayGame(agents []net.Conn, p1 Player, p2 Player) Player {
	agentList = agents
	runde = 1
	board = map[int]string{ // Lagrer spillets trekk. Tomme felt forblir tall som representerer posisjonen dens på brettet.
		1: "1", 2: "2", 3: "3",
		4: "4", 5: "5", 6: "6",
		7: "7", 8: "8", 9: "9"}

	multiplayerComm.PrintAll(agents, printBoard())
	multiplayerComm.PrintAll(agents, "\n"+p1.Name+" Starter.")
	newRoundOrGameOver(agents, p1, p2)
	return winner
}

// printBoard printer ut hvordan brettet ser ut med eventuelle trekk i terminalen.
func printBoard() string {

	line := " ------------- "
	wall := " | "
	shiftLine := 0

	boardArch := "\n" + line + "\n"

	for g := 0; g < 3; g++ { // Printer ut brettet, en linje om gangen.
		for i := 1; i <= 3; i++ {
			boardArch += wall + board[i+shiftLine]
		}
		boardArch += wall
		shiftLine += 3
		boardArch += "\n" + line + "\n"
	}

	return boardArch

}

// place Move plasserer trekk på brettet.
func placeMove(p1 Player, p2 Player) (Player, Player) {

	var p net.Conn

	if runde%2 == 0 {
		p = p2.Conn
	} else {
		p = p1.Conn
	}

	start := time.Now() // Tar tidspunktet siden fra når trekket starter.
	move := checkIfTaken(p)
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
	multiplayerComm.ClientPrintln(c, "\nSkriv inn tallet som representerer ruten du vil sette brikken din på.")
	move = multiplayerComm.ClientRead(agentList, c)

	if len(move) == 0 { //hvis tom input
		multiplayerComm.ClientPrintln(c, "Tomt trekk. Prøv igjen.")
		goto Again
	}
	parsedMove, valid := validInputs.CheckIfValid(move)

	if !valid {
		multiplayerComm.ClientPrintln(c, "Fant ikke input i støttede skriftsrpåk. Prøv igjen.")
		goto Again
	}

	if board[parsedMove] == "X" || board[parsedMove] == "O" {
		multiplayerComm.ClientPrintln(c, "Dette trekket er allerede tatt. Prøv igjen.")
		goto Again
	}
	return parsedMove
}

// newRoundOrGameOver sjekker om det er en vinner, uavgjort eller om den skal starte en ny runde.
func newRoundOrGameOver(agents []net.Conn, p1 Player, p2 Player) {
newRound:
	winningCombos := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 4, 7}, {2, 5, 8}, {3, 6, 9}, {1, 5, 9}, {3, 5, 7}}
	p1, p2 = placeMove(p1, p2) // plasserer brikken i på brettet og oppdaterer hver spiller har brukt.
	multiplayerComm.PrintAll(agents, printBoard())

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
		multiplayerComm.PrintAll(agents, "\nBrettet er fullt. Vinneren blir dermed avgjort på tid.")
		if p1.TimeUsed <= p2.TimeUsed {
			winner = p1
		} else {
			winner = p2
		}

		multiplayerComm.PrintAll(agents, "\n "+p1.Name+" brukte "+strconv.Itoa(p1.TimeUsed)+" millisekunder.")
		multiplayerComm.PrintAll(agents, " "+p2.Name+" brukte "+strconv.Itoa(p2.TimeUsed)+" millisekunder.")

		return
	}

	goto newRound
}
