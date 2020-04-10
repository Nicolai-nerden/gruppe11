package turneringMultiplayer

import (
	"fmt"
	"math"
	"tx3server/selvlagdePakker/tictactoe/tictactoeMultiplayer"
	"tx3server/selvlagdePakker/comm/multiplayerComm"
	"net"
	"strconv"
	"strings"
	"github.com/teamwork/reload"
)

var playerList []tictactoeMultiplayer.Player
var proceed string
var simulation bool

// Turnering spiller turnering på en gitt connection.
func Turnering(agents []net.Conn, names []string) {

	addPlayers(agents, names)

	if len(playerList) == 2 { //Hvis det en eller to spillere, er dette finale.
		finale(agents, playerList[0], playerList[1])
		goto end
	} else if len(playerList) == getNMax(len(playerList)) { // hvis antall spillere er en toerpotens (2, 4, 8, osv.) kan man gå rett til vanlig turnering.
		goto playOffs
	}
	playerList = kvalikk(agents, playerList)
	sortByTime(len(playerList))  //Sorterer etter kortest tid.
	sortByScore(len(playerList)) //Sorterer etter flest poeng.

	drawQualifying(agents, len(playerList)) //Trekker ut de som kvalifiserer seg basert på score og deretter tid brukt.
playOffs:
	playOffs(agents, playerList) //Starter turnering med de kvalifiserte spillerne gjentar seg selv intill det er en vinner.
end:
}

func addPlayers(agents []net.Conn, names []string) {
	for i := 0; i < len(agents); i++ {
		playerList = append(playerList, tictactoeMultiplayer.Player{agents[i], names[i], 0, 0})
	}
}

// sortByTime sorterer spiller listen etter tid brukt i kvalikken.
func sortByTime(playerAmount int) {

	var sortedPlayerList []tictactoeMultiplayer.Player

	for g := 0; g < playerAmount; g++ {

		shortestTimeUsed := playerList[0].TimeUsed
		shortestIndex := 0

		for i := 0; i < len(playerList); i++ { //henter ut den laveste verdien av timeUsed og den tilhørende indexen.
			if playerList[i].TimeUsed <= shortestTimeUsed {
				shortestTimeUsed = playerList[i].TimeUsed
				shortestIndex = i
			}
		}

		sortedPlayerList = append(sortedPlayerList, playerList[shortestIndex])
		playerList = append(playerList[:shortestIndex], playerList[shortestIndex+1:]...)

	}

	playerList = sortedPlayerList //oppdaterer tabellen, versjonen sortert etter tid.
}

// sortByScore sorterer spillerlisten etter poengsummen fra kvalikken
func sortByScore(playerAmount int) {

	var sortedPlayerList []tictactoeMultiplayer.Player

	for g := 0; g < playerAmount; g++ {

		highestScore := playerList[0].Score
		highestIndex := 0

		for i := 0; i < len(playerList); i++ { //henter ut den høyeste verdien av score og den tilhørende indexen.
			if playerList[i].Score > highestScore {
				highestScore = playerList[i].Score
				highestIndex = i
			}
		}

		sortedPlayerList = append(sortedPlayerList, playerList[highestIndex])
		playerList = append(playerList[:highestIndex], playerList[highestIndex+1:]...)

	}

	playerList = sortedPlayerList //oppdaterer tabellen, til versjonen sortert etter tid og score.
}

// drawQualifying velger ut de øverste 2^n av spillerlisten som er sortert etter poeng sum>tid brukt. Altså teller poeng mer enn tid brukt.
func drawQualifying(agents []net.Conn, playerAmount int) {

	var advancingPlayers []tictactoeMultiplayer.Player
	amountQualPlayers := getNMax(playerAmount)

	for i := 0; i < amountQualPlayers; i++ { //Henter ut de øverste 2^nMax antall spillerne i poengtabellen.

		advancingPlayers = append(advancingPlayers, playerList[i])
	}

	if len(advancingPlayers) == 2 { // hvis finale
		finale(agents, advancingPlayers[0], advancingPlayers[1])
		reload.Exec()
		return 
	}

	multiplayerComm.PrintAll(agents, "")
	multiplayerComm.PrintAll(agents, "Disse Spillerne er videre: ")
	multiplayerComm.PrintAll(agents, "")

	for i := 0; i < len(advancingPlayers); i++ {
		out := fmt.Sprintf("%d. %s Poeng: %d Tid brukt: %d", i+1, advancingPlayers[i].Name, advancingPlayers[i].Score, advancingPlayers[i].TimeUsed)
		multiplayerComm.PrintAll(agents, out)
	}

	multiplayerComm.PrintAll(agents, "")

	playerList = advancingPlayers

}

//sjekker høyeste 2^n som er innenfor antall spillere
func getNMax(playerAmount int) int {

	nMax := 0
	for int(math.Exp2(float64(nMax))) <= playerAmount {
		nMax++
	}

	if int(math.Exp2(float64(nMax))) != playerAmount { //Hvis nMax ikke er lik antall spillere betyr det at nMax er en for stor.
		nMax--
	}

	return int(math.Exp2(float64(nMax)))
}

// playOffs starter turneringen og setter opp matcher og deretter starter nye turneringsrunder helt til det er finale.
func playOffs(agents []net.Conn, remainingPlayers []tictactoeMultiplayer.Player) {

	var games []string
	var winners []tictactoeMultiplayer.Player

	for i := 0; i < len(remainingPlayers)/2; i++ { //Printer ut de ulike matchuppene.
		opponents := remainingPlayers[i].Name + " vs " + remainingPlayers[(len(remainingPlayers)-i)-1].Name
		games = append(games, opponents)
		multiplayerComm.PrintAll(agents, strconv.Itoa(i+1)+". match"+": "+opponents)
	}

	for i := 0; i < len(games); i++ { //Spiller ut turneringsrundens matcher

		multiplayerComm.PrintAll(agents, "\nNeste match:")
		multiplayerComm.PrintAll(agents, "\n-- "+games[i]+" --")
		multiplayerComm.PrintAll(agents, "\n"+remainingPlayers[i].Name+", trykk enter for å starte matchen.")
		proceed = multiplayerComm.ClientRead(remainingPlayers[i].Conn)

		//Starter match. Den beste mot den dårligste, nest beste mot nest dårligste osv.
		matchWinner := tictactoeMultiplayer.PlayGame(agents, remainingPlayers[i], remainingPlayers[(len(remainingPlayers)-i)-1])
		multiplayerComm.PrintAll(agents, "\n##################################")
		multiplayerComm.PrintAll(agents, "Vinner av matchen: "+matchWinner.Name)
		multiplayerComm.PrintAll(agents, "##################################")

		//legger til vinnerne i en egen slice
		winners = append(winners, matchWinner)
	}

	playerList = winners //Oppdaterer listen med spillere som vant runden.

	newRound(agents) //Starter en ny runde hvis det ikke er kåret en vinner.

}

// newRound starter ny turneringsrunde hvis ikke det er finale.
func newRound(agents []net.Conn) {

	if len(playerList) == 2 { //Hvis det kun er to igjen, er dette finale.
		finale(agents, playerList[0], playerList[1])
	} else {
		multiplayerComm.PrintAll(agents, "\nDet er "+strconv.Itoa(len(playerList)/2)+" turneringsrunder igjen.")
		multiplayerComm.PrintAll(agents, "\n"+playerList[0].Name+"Trykk enter for å starte neste turneringsrunde.")
		proceed = multiplayerComm.ClientRead(playerList[0].Conn)
		playOffs(agents, playerList)
	}
}

// kvalikk spiller kvalikk mellom spillerne lagt inn i pvp modus.
func kvalikk(agents []net.Conn, playerList []tictactoeMultiplayer.Player) []tictactoeMultiplayer.Player {

	var games [][]tictactoeMultiplayer.Player

	for i := 0; i < len(playerList); i++ { // Setter opp og printer ut de ulike matchuppene.
		if i < len(playerList)-1 {
			games = append(games, []tictactoeMultiplayer.Player{playerList[i], playerList[i+1]})
		} else {
			games = append(games, []tictactoeMultiplayer.Player{playerList[i], playerList[0]})
		}
	}

	games = gamesSorted(agents, games)

	for i := 0; i < len(games); i++ { //Spiller ut turneringsrundens matcher
		multiplayerComm.PrintAll(agents, "\nNeste match:")
		multiplayerComm.PrintAll(agents, "\n-- "+games[i][0].Name+" vs "+games[i][1].Name+" --")
		multiplayerComm.PrintAll(agents, "\n"+ games[i][0].Name+", trykk enter for å starte matchen.")
		proceed = multiplayerComm.ClientRead(games[i][0].Conn)

		//Starter matcher
		matchWinner := tictactoeMultiplayer.PlayGame(agents, games[i][0], games[i][1])
		matchWinner.Score += 3
		multiplayerComm.PrintAll(agents, "\n##################################")
		multiplayerComm.PrintAll(agents, "Vinner av matchen: "+matchWinner.Name)
		multiplayerComm.PrintAll(agents, "##################################")

		//legger til poeng og tid brukt til vinneren sin struct.
		g := 0
		for matchWinner.Name != playerList[g].Name {
			g++
		}
		playerList[g].Score = matchWinner.Score
		playerList[g].TimeUsed = matchWinner.TimeUsed
	}

	return playerList
}

// gamesSorted fordeler matchene slik at alle spillere spiller kvalifiseringsmatchene med et jevnt mellomrom
// i stedetfor at hver spiller spiller to matcher rett etter hverandre.
func gamesSorted(agents []net.Conn, games [][]tictactoeMultiplayer.Player) [][]tictactoeMultiplayer.Player {
	var gamesSorted [][]tictactoeMultiplayer.Player

	for i := 1; i <= 2; i++ {
		for g := 0; g < len(games); g++ {
			k := g + i
			if k%2 != 0 {
				gamesSorted = append(gamesSorted, games[g])
			}
		}
	}
	for i := 0; i < len(gamesSorted); i++ {
		multiplayerComm.PrintAll(agents, strconv.Itoa(i+1)+". match"+": "+gamesSorted[i][0].Name+" vs "+gamesSorted[i][1].Name)
	}
	return gamesSorted
}

// finale setter opp finale mellom de to finalistene.
func finale(agents []net.Conn, p1 tictactoeMultiplayer.Player, p2 tictactoeMultiplayer.Player) {
	opponents := p1.Name + " vs " + p2.Name
	multiplayerComm.PrintAll(agents, "\nchochocho chachacha DET ER TID FOR FINALE.")
	multiplayerComm.PrintAll(agents, "Finalen spilles mellom:")
	multiplayerComm.PrintAll(agents, opponents)
	multiplayerComm.PrintAll(agents, "\n"+p1.Name+", mash enter for å starte FINALEN")
	proceed = multiplayerComm.ClientRead(p1.Conn)

	tournamentWinner := tictactoeMultiplayer.PlayGame(agents, p1, p2)

	multiplayerComm.PrintAll(agents, "\nVINNEREN AV TURNERINGEN ER:")
	multiplayerComm.PrintAll(agents, "\n##################################")
	multiplayerComm.PrintAll(agents, "---------- "+strings.ToUpper(tournamentWinner.Name)+" ----------")
	multiplayerComm.PrintAll(agents, "##################################")
	multiplayerComm.PrintAll(agents, "\n\n")
}
