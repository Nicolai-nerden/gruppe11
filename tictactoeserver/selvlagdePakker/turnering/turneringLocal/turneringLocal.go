package turneringLocal

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
	"tx3server/selvlagdePakker/comm/localComm"
	"tx3server/selvlagdePakker/tictactoe/tictactoeLocal"
	"tx3server/selvlagdePakker/utf/supportedLangs"
	"tx3server/selvlagdePakker/utf/validInputs"
)

var playerList []tictactoe.Player
var proceed string
var simulation bool

// Turnering spiller turnering på en gitt connection.
func Turnering(c net.Conn) {
	chooseGameMode(c)

	if len(playerList) <= 2 { //Hvis det en eller to spillere, er dette finale.
		finale(c)
		goto end
	} else if len(playerList) == getNMax(len(playerList)) { // hvis antall spillere er en toerpotens (2, 4, 8, osv.) kan man gå rett til vanlig turnering.
		goto playOffs
	}
	sortByTime(len(playerList))  //Sorterer etter kortest tid.
	sortByScore(len(playerList)) //Sorterer etter flest poeng.

	drawQualifying(c, len(playerList)) //Trekker ut de som kvalifiserer seg basert på score og deretter tid brukt.
playOffs:
	playOffs(c, playerList) //Starter turnering med de kvalifiserte spillerne gjentar seg selv intill det er en vinner.
end:
}

// chooseGameMode velger modus, enten simulasjon eller PvP. Hvis simulasjon lager den random spillere og kvalikkresultater.
// Hvis PvP legger den til ønskede spillere og starter kvalikk.
func chooseGameMode(c net.Conn) {
	clientcommunication.ClientPrintln(c, "\nDette spillet har to moduser. Simulasjon og PvP.")
	clientcommunication.ClientPrintln(c, "1. PvP: Her spiller man turneringen. Fyll inn navn for hver deltaker og spill mot hverandre.")
	clientcommunication.ClientPrintln(c, "2. Simulasjon: Simulerer spillere, deres kvalifiseringsresultater og turnering.")
	clientcommunication.ClientPrintln(c, "   Du trenger kun å trykke enter-knappen underveis for å simulere hver match i simuleringsmodus.")

	clientcommunication.ClientPrint(c, "\nStøttede skriftspråk: ")
	for i := 0; i < len(supportedLangs.SupportedLangs); i++ {
		if i == len(supportedLangs.SupportedLangs)-1 {
			clientcommunication.ClientPrint(c, "og "+supportedLangs.SupportedLangs[i]) //linjeshift
			break
		}
		clientcommunication.ClientPrint(c, supportedLangs.SupportedLangs[i]+", ")
	}
chooseMode:
	clientcommunication.ClientPrintln(c, "\nHvilken modus ønsker du å bruke?")
	clientcommunication.ClientPrintln(c, "\n1. PvP")
	clientcommunication.ClientPrintln(c, "2. Simulasjon")
	input := clientcommunication.ClientRead(c)
	modus, valid := validInputs.CheckIfValid(input)

	if !valid {
		clientcommunication.ClientPrintln(c, "\nFant ikke input i støttede skriftsrpåk. Prøv igjen.")
		goto chooseMode
	}
	if modus == 1 {
		simulation = false
	addPlayers:
		addPlayers(c)
		if len(playerList) < 1 {
			clientcommunication.ClientPrintln(c, "\nDet må minst være en spiller. Prøv igjen.")
			goto addPlayers
			return
		} else if len(playerList) <= 2 || len(playerList) == getNMax(len(playerList)) { //Hvis det en eller to spillere, er dette finale. Eller hvis det ikke er behov for kvalikk.
			return
		}
		kvalikk(c, playerList, simulation)
	} else if modus == 2 {
		simulation = true
		var playerAmount int
	insertPlayers:
		clientcommunication.ClientPrintln(c, "\nSkriv inn antallet spillere du ønsker å simulere med vanlige tall.")
		clientcommunication.ClientPrintln(c, "(Denne inputen støtter kun de vanlige tall slik som \"20\" for at man skal kunne velge opp mot evig antall simulerte spillere")
		clientcommunication.ClientPrintln(c, "Dette er for at inputvalidatoren vår kun går opp til 9. Alle andre inputs i spillet kan skrives med skrifttegn fra de andre støttede skriftspråkene.)")

		playerAmount, _ = strconv.Atoi(clientcommunication.ClientRead(c))
		if playerAmount < 1 {
			clientcommunication.ClientPrintln(c, "Det må minst være en simulert spiller.")
			goto insertPlayers
		}
		simulateplayerList(playerAmount) //Henter simulerte navn.
		simulateGameInput(playerAmount)  //Henter inn simulerte tid- og poengresultater.
	} else {
		clientcommunication.ClientPrintln(c, "Ugyldig input. Prøv igjen.")
		goto chooseMode
	}

}

// addPlayers legger til spillere helt til man skriver "start", returnerer dermed tilbake til chooseGameMode
func addPlayers(c net.Conn) {
	var inputPlayer string
nySpiller:
	clientcommunication.ClientPrintln(c, "\nSkriv inn navn på deltager du vil legge til.")
	clientcommunication.ClientPrintln(c, "Skriv \"start\" for å starte turneringen.")
	inputPlayer = clientcommunication.ClientRead(c)

	if strings.TrimSpace(strings.ToLower(inputPlayer)) == "start" {
		return
	} else if len(inputPlayer) < 1 {
		clientcommunication.ClientPrintln(c, "\nDu kan ikke legge inn tomme navn. Prøv igjen.")
		goto nySpiller
	}

	playerList = append(playerList, tictactoe.Player{inputPlayer, 0, 0})
	goto nySpiller
}

func startQualifying([]tictactoe.Player) []tictactoe.Player {
	return playerList
}

func simulateplayerList(playerAmount int) {
	rand.Seed(time.Now().UnixNano()) //forandrer seedet etter hva tiden er. Slik at det gir ulike resultater.
	var adjektiv = []string{"Gnålete", "Håpløs", "Keeg", "Nerdete"}
	var substantiv = []string{"løve", "flodhest", "ape", "sjøku"}

	for i := 0; i < playerAmount; i++ {
		p := tictactoe.Player{adjektiv[rand.Intn(4)] + " " + substantiv[rand.Intn(4)], 0, 0}
		playerList = append(playerList, p)
	}
}

// simulateGameInput simulerer resultat av kvalifiseringsrunde for simulerte spillere.
func simulateGameInput(playerAmount int) {
	possibleResults := []int{0, 1, 3}

	for i := 0; i < playerAmount; i++ { //Simulerer poengsummer fra spill.

		var gameResult = possibleResults[rand.Intn(3)]
		var gameResult2 = possibleResults[rand.Intn(3)]

		playerList[i].Score = (gameResult + gameResult2)

		//Simulerer tid hver spiller har brukt på trekk denne runden.
		var gameTime1 = rand.Intn(70) + 30
		var gameTime2 = rand.Intn(70) + 30

		playerList[i].TimeUsed = (gameTime1 + gameTime2)
	}

}

// sortByTime sorterer spiller listen etter tid brukt i kvalikken.
func sortByTime(playerAmount int) {

	var sortedPlayerList []tictactoe.Player

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

	var sortedPlayerList []tictactoe.Player

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
func drawQualifying(c net.Conn, playerAmount int) {

	var advancingPlayers []tictactoe.Player
	amountQualPlayers := getNMax(playerAmount)

	for i := 0; i < amountQualPlayers; i++ { //Henter ut de øverste 2^nMax antall spillerne i poengtabellen.

		advancingPlayers = append(advancingPlayers, playerList[i])
	}

	clientcommunication.ClientPrintln(c, "")
	clientcommunication.ClientPrintln(c, "Disse Spillerne er videre: ")
	clientcommunication.ClientPrintln(c, "")

	for i := 0; i < len(advancingPlayers); i++ {
		out := fmt.Sprintf("%d. %s Poeng: %d Tid brukt: %d", i+1, advancingPlayers[i].Name, advancingPlayers[i].Score, advancingPlayers[i].TimeUsed)
		clientcommunication.ClientPrintln(c, out)
	}

	clientcommunication.ClientPrintln(c, "")

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
func playOffs(c net.Conn, remainingPlayers []tictactoe.Player) {

	var games []string
	var winners []tictactoe.Player

	for i := 0; i < len(remainingPlayers)/2; i++ { //Printer ut de ulike matchuppene.
		opponents := remainingPlayers[i].Name + " vs " + remainingPlayers[(len(remainingPlayers)-i)-1].Name
		games = append(games, opponents)
		clientcommunication.ClientPrintln(c, strconv.Itoa(i+1)+". match"+": "+opponents)
	}

	for i := 0; i < len(games); i++ { //Spiller ut turneringsrundens matcher

		clientcommunication.ClientPrintln(c, "\nNeste match:")
		clientcommunication.ClientPrintln(c, "\n-- "+games[i]+" --")
		clientcommunication.ClientPrintln(c, "\nTrykk enter for å starte matchen.")
		proceed = clientcommunication.ClientRead(c)

		//Starter match. Den beste mot den dårligste, nest beste mot nest dårligste osv.
		matchWinner := tictactoe.PlayGame(c, remainingPlayers[i], remainingPlayers[(len(remainingPlayers)-i)-1], simulation)
		clientcommunication.ClientPrintln(c, "\n##################################")
		clientcommunication.ClientPrintln(c, "Vinner av matchen: "+matchWinner.Name)
		clientcommunication.ClientPrintln(c, "##################################")

		//legger til vinnerne i en egen slice
		winners = append(winners, matchWinner)
	}

	playerList = winners //Oppdaterer listen med spillere som vant runden.

	newRound(c) //Starter en ny runde hvis det ikke er kåret en vinner.

}

// newRound starter ny turneringsrunde hvis ikke det er finale.
func newRound(c net.Conn) {

	if len(playerList) <= 2 { //Hvis det kun er to igjen, er dette finale.
		finale(c)
	} else {
		clientcommunication.ClientPrintln(c, "\nDet er "+strconv.Itoa(len(playerList)/2)+" turneringsrunder igjen.")
		clientcommunication.ClientPrintln(c, "\nTrykk enter for å starte neste turneringsrunde.")
		proceed = clientcommunication.ClientRead(c)
		playOffs(c, playerList)
	}
}

// kvalikk spiller kvalikk mellom spillerne lagt inn i pvp modus.
func kvalikk(c net.Conn, playerList []tictactoe.Player, simulation bool) []tictactoe.Player {

	var games [][]tictactoe.Player

	for i := 0; i < len(playerList); i++ { // Setter opp og printer ut de ulike matchuppene.
		if i < len(playerList)-1 {
			games = append(games, []tictactoe.Player{playerList[i], playerList[i+1]})
		} else {
			games = append(games, []tictactoe.Player{playerList[i], playerList[0]})
		}
	}

	games = gamesSorted(c, games)

	for i := 0; i < len(games); i++ { //Spiller ut turneringsrundens matcher
		clientcommunication.ClientPrintln(c, "\nNeste match:")
		clientcommunication.ClientPrintln(c, "\n-- "+games[i][0].Name+" vs "+games[i][1].Name+" --")
		clientcommunication.ClientPrintln(c, "\nTrykk enter for å starte matchen.")
		proceed = clientcommunication.ClientRead(c)

		//Starter matcher
		matchWinner := tictactoe.PlayGame(c, games[i][0], games[i][1], simulation)
		matchWinner.Score += 3
		clientcommunication.ClientPrintln(c, "\n##################################")
		clientcommunication.ClientPrintln(c, "Vinner av matchen: "+matchWinner.Name)
		clientcommunication.ClientPrintln(c, "##################################")

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
func gamesSorted(c net.Conn, games [][]tictactoe.Player) [][]tictactoe.Player {
	var gamesSorted [][]tictactoe.Player

	for i := 1; i <= 2; i++ {
		for g := 0; g < len(games); g++ {
			k := g + i
			if k%2 != 0 {
				gamesSorted = append(gamesSorted, games[g])
			}
		}
	}
	for i := 0; i < len(gamesSorted); i++ {
		clientcommunication.ClientPrintln(c, strconv.Itoa(i+1)+". match"+": "+gamesSorted[i][0].Name+" vs "+gamesSorted[i][1].Name)
	}
	return gamesSorted
}

// finale setter opp finale mellom de to finalistene.
func finale(c net.Conn) {
	opponents := playerList[0].Name + " vs " + playerList[len(playerList)-1].Name
	clientcommunication.ClientPrintln(c, "\nchochocho chachacha DET ER TID FOR FINALE.")
	clientcommunication.ClientPrintln(c, "Finalen spilles mellom:")
	clientcommunication.ClientPrintln(c, opponents)
	clientcommunication.ClientPrintln(c, "\nSmash enter for å starte FINALEN")
	proceed = clientcommunication.ClientRead(c)

	tournamentWinner := tictactoe.PlayGame(c, playerList[0], playerList[len(playerList)-1], simulation)

	clientcommunication.ClientPrintln(c, "\nVINNEREN AV TURNERINGEN ER:")
	clientcommunication.ClientPrintln(c, "\n##################################")
	clientcommunication.ClientPrintln(c, "---------- "+strings.ToUpper(tournamentWinner.Name)+" ----------")
	clientcommunication.ClientPrintln(c, "##################################")
	clientcommunication.ClientPrintln(c, "\n\n")
}
