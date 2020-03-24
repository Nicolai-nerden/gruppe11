package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

var playerNames []string
var playerScores []int
var playerTimeUsed []int
var playerAmount = 10
var amountQualPlayers int
var advancingPlayers []string
var winner string

func main() {

	createNames(playerAmount) //Henter simulerte navn.
	gameInput(playerAmount)   //Henter inn simulerte tid- og poengresultater.

	stockByTime(playerAmount)  //Sorterer etter kortest tid.
	stockByScore(playerAmount) //Sorterer etter flest poeng.

	drawQualifying(playerAmount) //Trekker ut de som kvalifiserer seg basert på score og deretter tid brukt.
	startTournamentRound(advancingPlayers) //Starter turnering med de kvalifiserte spillerne gjentar seg selv intill det er en vinner.


}

func checkWinner(){
	if len(advancingPlayers) == 1 {
		fmt.Println("The winner of the tournament is:", advancingPlayers[0])
	} else {
		startTournamentRound(advancingPlayers)
	}
}

func startTournamentRound(players []string) {

	var games []string
	var winners []string
	var proceed string
	//fmt.Println(advancingPlayers)

	for i := 0; i < len(players)/2; i++ { //lager en slice med de ulike matchuppene.
		var opponents string = "Game " + strconv.Itoa(i+1) + ": " + players[i] + " vs " + players[(len(players)-i)-1]
		games = append(games, opponents)
		fmt.Println(opponents)
	}

	for i := 0; i < len(games); i++ {
		fmt.Println("\n" + games[i], "\nPress enter to continue.")
		fmt.Scanln(&proceed)
		playGame(players[i], players[(len(players)-i)-1])
		winners = append(winners, winner)
	}

	advancingPlayers = winners

	checkWinner()
}

func playGame(p1 string, p2 string){ //Simulerer et spill, 50/50 hvem som vinner
	var rng int = rand.Intn(2)

	if rng == 0 {
		winner = p1
	}else {
		winner = p2
	}

	fmt.Println("Winner:", winner, "\n ")
}

func drawQualifying(players int) {

	nMax := 0

	//sjekker høyeste 2^n som er innenfor antall spillere
	for int(math.Exp2(float64(nMax))) <= players {
		nMax++
	}

	if int(math.Exp2(float64(nMax))) != players {
		nMax--
	}
	//fmt.Println(nMax)
	//fmt.Println(int(math.Exp2(float64(nMax))))

	amountQualPlayers = int(math.Exp2(float64(nMax)))

	//Henter ut de øverste 2^nMax antall spillerne i poengtabellen.

	for i := 0; i < amountQualPlayers; i++ {

		advancingPlayers = append(advancingPlayers, playerNames[i])
	}

	/*fmt.Println("Disse Spillerne er videre: ")

	for i := 0; i < len(advancingPlayers); i++ {
		fmt.Println(playerNames[i], "Poeng:", playerScores[i], "Tid brukt:", playerTimeUsed[i])
	}*/

}

func stockByTime(players int) {
	var stockedTimePoints = []int{}
	var stockedNames = []string{}
	var stockedScores = []int{}

	for g := 0; g < players; g++ {

		smallestValue := playerTimeUsed[0]
		smallestIndex := 0

		for i := 0; i < len(playerTimeUsed); i++ { //Finner minste verdi i slicen

			if playerTimeUsed[i] <= smallestValue {
				smallestValue = playerTimeUsed[i]
				smallestIndex = i
			}
		}

		//Legger til verdien vi fant ut var størst i den nye slicen og rokkerer de navn og poengTabellene derretter.
		stockedTimePoints = append(stockedTimePoints, playerTimeUsed[smallestIndex])
		stockedNames = append(stockedNames, playerNames[smallestIndex])
		stockedScores = append(stockedScores, playerScores[smallestIndex])

		//Fjerner elementet som vi la til i den rangerte slicen fra den vi hentet fra.
		playerTimeUsed = append(playerTimeUsed[:smallestIndex], playerTimeUsed[smallestIndex+1:]...)
		playerNames = append(playerNames[:smallestIndex], playerNames[smallestIndex+1:]...)
		playerScores = append(playerScores[:smallestIndex], playerScores[smallestIndex+1:]...)
	}

	//fmt.Println(stockedNames,": ", stockedTimePoints)

	playerTimeUsed = stockedTimePoints
	playerNames = stockedNames
	playerScores = stockedScores
}

func stockByScore(players int) {

	var stockedTimePoints = []int{}
	var stockedNames = []string{}
	var stockedScores = []int{}

	for g := 0; g < players; g++ {

		biggestValue := playerScores[0]
		biggestIndex := 0

		for i := 0; i < len(playerScores); i++ { //Finner største verdi i slicen

			if playerScores[i] > biggestValue {
				biggestValue = playerScores[i]
				biggestIndex = i
			}
		}

		//Legger til verdien vi fant ut var størst i den nye slicen og rokkerer de navn og poengTabellene derretter.
		stockedTimePoints = append(stockedTimePoints, playerTimeUsed[biggestIndex])
		stockedNames = append(stockedNames, playerNames[biggestIndex])
		stockedScores = append(stockedScores, playerScores[biggestIndex])

		//Fjerner elementet som vi la til i den rangerte slicen fra den vi hentet fra.
		playerTimeUsed = append(playerTimeUsed[:biggestIndex], playerTimeUsed[biggestIndex+1:]...)
		playerNames = append(playerNames[:biggestIndex], playerNames[biggestIndex+1:]...)
		playerScores = append(playerScores[:biggestIndex], playerScores[biggestIndex+1:]...)

	}

	//fmt.Println(stockedScores, ":", stockedTimePoints)

	playerTimeUsed = stockedTimePoints
	playerNames = stockedNames
	playerScores = stockedScores

}

func gameInput(players int) { // Simulerer resultat av kvalifiseringsrunde

	possibleResults := []int{0, 1, 3}

	for i := 0; i < players; i++ {
		//Simulerer poengsummer fra spill.
		var gameResult = possibleResults[rand.Intn(3)]
		var gameResult2 = possibleResults[rand.Intn(3)]
		playerScores = append(playerScores, (gameResult + gameResult2))
		// fmt.Println(playerScores[i])

		//Simulerer tid hver spiller har brukt på trekk denne runden.
		var gameTime1 = rand.Intn(70) + 30
		var gameTime2 = rand.Intn(70) + 30
		playerTimeUsed = append(playerTimeUsed, gameTime1+gameTime2)
		// fmt.Println(playerTimeUsed[i])

	}
}

func createNames(players int) {
	var adjektiv = []string{"Gnålete", "Håpløs", "Keeg", "Nerdete"}
	var substantiv = []string{"løve", "flodhest", "ape", "sjøku"}

	for i := 0; i < players; i++ {

		playerNames = append(playerNames, adjektiv[rand.Intn(4)]+" "+substantiv[rand.Intn(4)])
		//fmt.Println(playerNames[i])
	}

}

