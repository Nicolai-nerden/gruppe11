package main

import (
	"armyboyz/tictactoe"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
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
	playerList = startQualifying(playerList)
	sortByTime(len(playerList))  //Sorterer etter kortest tid.
	sortByScore(len(playerList)) //Sorterer etter flest poeng.

	drawQualifying(len(playerList)) //Trekker ut de som kvalifiserer seg basert på score og deretter tid brukt.

	startTournament(playerList) //Starter turnering med de kvalifiserte spillerne gjentar seg selv intill det er en vinner.
end:
}

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

func addPlayers() { // Legger til spillere helt til man skriver "start", returnerer dermed tilbake til main()
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

func simulateGameInput(playerAmount int) { // Simulerer resultat av kvalifiseringsrunde
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

func startTournament(remainingPlayers []tictactoe.Player) {

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

func newRound() {

	if len(playerList) <= 2 { //Hvis det kun er to igjen, er dette finale.
		finale()
	} else {
		fmt.Println("\nDet er ", (len(playerList) / 2), "turneringsrunder igjen.")
		fmt.Println("\nTrykk enter for å starte neste turneringsrunde.")
		fmt.Scanln(&proceed)
		startTournament(playerList)
	}
}

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
