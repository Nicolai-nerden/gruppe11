package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type player struct {
	name     string
	score    int
	timeUsed int
}

var playerList []player
var proceed string

func main() {

	playerAmount, _ := strconv.Atoi(os.Args[1]) //Skriv et tall på slutten av kommando linjen i terminalen for å sette hvor mange spillere som skal simuleres.

	simulateplayerList(playerAmount) //Henter simulerte navn.
	simulateGameInput(playerAmount)  //Henter inn simulerte tid- og poengresultater.

	sortByTime(len(playerList))  //Sorterer etter kortest tid.
	sortByScore(len(playerList)) //Sorterer etter flest poeng.

	drawQualifying(len(playerList)) //Trekker ut de som kvalifiserer seg basert på score og deretter tid brukt.

	startTournament(playerList) //Starter turnering med de kvalifiserte spillerne gjentar seg selv intill det er en vinner.
}

func simulateplayerList(playerAmount int) {

	//rand.Seed(time.Now().UnixNano()) //forandrer seedet etter hva tiden er. Slik at det gir ulike resultater.

	var adjektiv = []string{"Gnålete", "Håpløs", "Keeg", "Nerdete"}
	var substantiv = []string{"løve", "flodhest", "ape", "sjøku"}

	for i := 0; i < playerAmount; i++ {

		p := player{adjektiv[rand.Intn(4)] + " " + substantiv[rand.Intn(4)], 0, 0}

		playerList = append(playerList, p)
	}

}

func simulateGameInput(playerAmount int) { // Simulerer resultat av kvalifiseringsrunde

	possibleResults := []int{0, 1, 3}

	for i := 0; i < playerAmount; i++ { //Simulerer poengsummer fra spill.

		var gameResult = possibleResults[rand.Intn(3)]
		var gameResult2 = possibleResults[rand.Intn(3)]

		playerList[i].score = (gameResult + gameResult2)

		//Simulerer tid hver spiller har brukt på trekk denne runden.
		var gameTime1 = rand.Intn(70) + 30
		var gameTime2 = rand.Intn(70) + 30

		playerList[i].timeUsed = (gameTime1 + gameTime2)
	}

}

func sortByTime(playerAmount int) {

	var sortedPlayerList []player

	for g := 0; g < playerAmount; g++ {

		shortestTimeUsed := playerList[0].timeUsed
		shortestIndex := 0

		for i := 0; i < len(playerList); i++ { //henter ut den laveste verdien av timeUsed og den tilhørende indexen.
			if playerList[i].timeUsed <= shortestTimeUsed {
				shortestTimeUsed = playerList[i].timeUsed
				shortestIndex = i
			}
		}

		sortedPlayerList = append(sortedPlayerList, playerList[shortestIndex])
		playerList = append(playerList[:shortestIndex], playerList[shortestIndex+1:]...)

	}

	playerList = sortedPlayerList //oppdaterer tabellen, versjonen sortert etter tid.
}

func sortByScore(playerAmount int) {

	var sortedPlayerList []player

	for g := 0; g < playerAmount; g++ {

		highestScore := playerList[0].score
		highestIndex := 0

		for i := 0; i < len(playerList); i++ { //henter ut den høyeste verdien av score og den tilhørende indexen.
			if playerList[i].score > highestScore {
				highestScore = playerList[i].score
				highestIndex = i
			}
		}

		sortedPlayerList = append(sortedPlayerList, playerList[highestIndex])
		playerList = append(playerList[:highestIndex], playerList[highestIndex+1:]...)

	}

	playerList = sortedPlayerList //oppdaterer tabellen, til versjonen sortert etter tid og score.
}

func drawQualifying(playerAmount int) {

	var advancingPlayers []player
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
		fmt.Println(strconv.Itoa(i+1)+".", advancingPlayers[i].name, "Poeng:", advancingPlayers[i].score, "Tid brukt:", advancingPlayers[i].timeUsed)
	}

	fmt.Println()

	playerList = advancingPlayers

}

func startTournament(remainingPlayers []player) {

	var games []string
	var winners []player

	for i := 0; i < len(remainingPlayers)/2; i++ { //Printer ut de ulike matchuppene.
		opponents := remainingPlayers[i].name + " vs " + remainingPlayers[(len(remainingPlayers)-i)-1].name
		games = append(games, opponents)
		fmt.Println(strconv.Itoa(i+1) + ". match" + ": " + opponents)
	}

	for i := 0; i < len(games); i++ { //Spiller ut turneringsrundens matcher

		fmt.Println("\nNeste match:")
		fmt.Println("\n-- " + games[i] + " --")
		fmt.Println("\nTrykk enter for å starte matchen.")
		fmt.Scanln(&proceed)

		//Starter match. Den beste mot den dårligste, nest beste mot nest dårligste osv.
		matchWinner := playGame(remainingPlayers[i], remainingPlayers[(len(remainingPlayers)-i)-1])
		fmt.Println("\n##################################")
		fmt.Println("Vinner av matchen:", matchWinner.name)
		fmt.Println("##################################")

		//legger til vinnerne i en egen slice
		winners = append(winners, matchWinner)
	}

	playerList = winners //Oppdaterer listen med spillere som vant runden.

	newRound() //Starter en ny runde hvis det ikke er kåret en vinner.

}

func newRound() {

	if len(playerList) == 2 { //Hvis det kun er to igjen, er dette finale.
		fmt.Println("\nchochocho chachacha DET ER TID FOR FINALE.")
		fmt.Println("\nSmash enter for å starte FINALEN")
		fmt.Scanln(&proceed)

		tournamentWinner := playGame(playerList[0], playerList[1])

		fmt.Println("\nVINNEREN AV TURNERINGEN ER:")
		fmt.Println("\n##################################")
		fmt.Println("----------", strings.ToUpper(tournamentWinner.name), "----------")
		fmt.Println("##################################")
		fmt.Println("\n\n")

	} else {
		fmt.Println("\nDet er ", (len(playerList) / 2), "turneringsrunder igjen.")
		fmt.Println("\nTrykk enter for å starte neste turneringsrunde.")
		fmt.Scanln(&proceed)
		startTournament(playerList)
	}
}

func playGame(p1 player, p2 player) player { //Simulerer et spill, 50/50 hvem som vinner
	var rng int = rand.Intn(2)
	var winner player

	if rng == 0 {
		winner = p1
	} else {
		winner = p2
	}
	return winner
}
