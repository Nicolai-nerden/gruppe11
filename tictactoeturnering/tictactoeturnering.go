package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"tictactoe"
	"time"
)

var playerList []tictactoe.Player
var proceed string
var simulation bool

func main() {

	chooseGameMode()

	if len(playerList) <= 2 { //Hvis det en eller to spillere, er dette finale.
		finale()
		goto end
	}
	sortByTime(len(playerList))  //Sorterer etter kortest tid.
	sortByScore(len(playerList)) //Sorterer etter flest poeng.

	drawQualifying(len(playerList)) //Trekker ut de som kvalifiserer seg basert på score og deretter tid brukt.

	playOffs(playerList) //Starter turnering med de kvalifiserte spillerne gjentar seg selv intill det er en vinner.
end:
}

// chooseGameMode velger modus, enten simulasjon eller PvP. Hvis simulasjon lager den random spillere og kvalikkresultater.
// Hvis PvP legger den til ønskede spillere og starter kvalikk.
func chooseGameMode() {
	var modus string
	fmt.Println("\nDette spillet har to moduser. Simulasjon og PvP.")
	fmt.Println("Simulasjon: Simulerer spillere, deres kvalifiserings resultater og turnering. Du trenger kun å trykke enter-knappen underveis for å simulere hver match")
	fmt.Println("PvP: Her spiller man turneringen. Fyll inn navn for hver deltaker og spill mot hverandre.")
chooseMode:
	fmt.Println("\nHvilken modus ønsker du å bruke?")
	fmt.Println("\n1. PvP")
	fmt.Println("2. Simulasjon")
	fmt.Scanln(&modus)
	modusConv, _ := strconv.Atoi(modus)
	if modusConv == 1 {
		simulation = false
		addPlayers()
		kvalikk(playerList, simulation)
	} else if modusConv == 2 {
		simulation = true
		var playerAmount int
	insertPlayers:
		fmt.Println("Skriv inn antallet spillere du ønsker å simulere.")
		fmt.Scanln(&playerAmount)
		if playerAmount < 1 {
			fmt.Println("Det må minst være en simulert spiller.")
			goto insertPlayers
		}
		simulateplayerList(playerAmount) //Henter simulerte navn.
		simulateGameInput(playerAmount)  //Henter inn simulerte tid- og poengresultater.
	} else {
		fmt.Println("Ugyldig input. Prøv igjen.")
		goto chooseMode
	}
}

// addPlayers legger til spillere helt til man skriver "start", returnerer dermed tilbake til chooseGameMode
func addPlayers() {
	var inputPlayer string
nySpiller:
	fmt.Println("\nSkriv inn navn på deltager du vil legge til.")
	fmt.Println("Skriv \"start\" for å starte turneringen.")
	fmt.Scanln(&inputPlayer)

	if strings.TrimSpace(strings.ToLower(inputPlayer)) == "start" {
		return
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
func drawQualifying(playerAmount int) {

	var advancingPlayers []tictactoe.Player
	nMax := 0

	//sjekker høyeste 2^n som er innenfor antall spillere
	for int(math.Exp2(float64(nMax))) <= playerAmount {
		nMax++
	}

	if int(math.Exp2(float64(nMax))) != playerAmount { //Hvis nMax ikke er lik antall spillere betyr det at nMax er en for stor.
		nMax--
	}

	amountQualPlayers := int(math.Exp2(float64(nMax)))

	for i := 0; i < amountQualPlayers; i++ { //Henter ut de øverste 2^nMax antall spillerne i poengtabellen.

		advancingPlayers = append(advancingPlayers, playerList[i])
	}

	fmt.Println()
	fmt.Println("Disse Spillerne er videre: ")
	fmt.Println()

	for i := 0; i < len(advancingPlayers); i++ {
		fmt.Println(strconv.Itoa(i+1)+".", advancingPlayers[i].Name, "Poeng:", advancingPlayers[i].Score, "Tid brukt:", advancingPlayers[i].TimeUsed)
	}

	fmt.Println()

	playerList = advancingPlayers

}

// playOffs starter turneringen og setter opp matcher og deretter starter nye turneringsrunder helt til det er finale.
func playOffs(remainingPlayers []tictactoe.Player) {

	var games []string
	var winners []tictactoe.Player

	for i := 0; i < len(remainingPlayers)/2; i++ { //Printer ut de ulike matchuppene.
		opponents := remainingPlayers[i].Name + " vs " + remainingPlayers[(len(remainingPlayers)-i)-1].Name
		games = append(games, opponents)
		fmt.Println(strconv.Itoa(i+1) + ". match" + ": " + opponents)
	}

	for i := 0; i < len(games); i++ { //Spiller ut turneringsrundens matcher

		fmt.Println("\nNeste match:")
		fmt.Println("\n-- " + games[i] + " --")
		fmt.Println("\nTrykk enter for å starte matchen.")
		fmt.Scanln(&proceed)

		//Starter match. Den beste mot den dårligste, nest beste mot nest dårligste osv.
		matchWinner := tictactoe.PlayGame(remainingPlayers[i], remainingPlayers[(len(remainingPlayers)-i)-1], simulation)
		fmt.Println("\n##################################")
		fmt.Println("Vinner av matchen:", matchWinner.Name)
		fmt.Println("##################################")

		//legger til vinnerne i en egen slice
		winners = append(winners, matchWinner)
	}

	playerList = winners //Oppdaterer listen med spillere som vant runden.

	newRound() //Starter en ny runde hvis det ikke er kåret en vinner.

}

// newRound starter ny turneringsrunde hvis ikke det er finale.
func newRound() {

	if len(playerList) <= 2 { //Hvis det kun er to igjen, er dette finale.
		finale()
	} else {
		fmt.Println("\nDet er ", (len(playerList) / 2), "turneringsrunder igjen.")
		fmt.Println("\nTrykk enter for å starte neste turneringsrunde.")
		fmt.Scanln(&proceed)
		playOffs(playerList)
	}
}

// kvalikk spiller kvalikk mellom spillerne lagt inn i pvp modus.
func kvalikk(playerList []tictactoe.Player, simulation bool) []tictactoe.Player {

	var games [][]tictactoe.Player

	for i := 0; i < len(playerList); i++ { // Setter opp og printer ut de ulike matchuppene.
		if i < len(playerList)-1 {
			games = append(games, []tictactoe.Player{playerList[i], playerList[i+1]})
		} else {
			games = append(games, []tictactoe.Player{playerList[i], playerList[0]})
		}
	}

	games = gamesSorted(games)

	for i := 0; i < len(games); i++ { //Spiller ut turneringsrundens matcher
		fmt.Println("\nNeste match:")
		fmt.Println("\n-- " + games[i][0].Name + " vs " + games[i][1].Name + " --")
		fmt.Println("\nTrykk enter for å starte matchen.")
		fmt.Scanln(&proceed)

		//Starter matcher
		matchWinner := tictactoe.PlayGame(games[i][0], games[i][1], simulation)
		matchWinner.Score += 3
		fmt.Println("\n##################################")
		fmt.Println("Vinner av matchen:", matchWinner.Name)
		fmt.Println("##################################")

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
func gamesSorted(games [][]tictactoe.Player) [][]tictactoe.Player {
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
		fmt.Println(strconv.Itoa(i+1) + ". match" + ": " + gamesSorted[i][0].Name + " vs " + gamesSorted[i][1].Name)
	}
	return gamesSorted
}

// finale setter opp finale mellom de to finalistene.
func finale() {
	opponents := playerList[0].Name + " vs " + playerList[len(playerList)-1].Name
	fmt.Println("\nchochocho chachacha DET ER TID FOR FINALE.")
	fmt.Println("Finalen spilles mellom:")
	fmt.Println(opponents)
	fmt.Println("\nSmash enter for å starte FINALEN")
	fmt.Scanln(&proceed)

	tournamentWinner := tictactoe.PlayGame(playerList[0], playerList[len(playerList)-1], simulation)

	fmt.Println("\nVINNEREN AV TURNERINGEN ER:")
	fmt.Println("\n##################################")
	fmt.Println("----------", strings.ToUpper(tournamentWinner.Name), "----------")
	fmt.Println("##################################")
	fmt.Println("\n\n")
}
