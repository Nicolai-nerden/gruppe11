package main

import (
	"fmt"
	"strconv"
)

var lastAnswer int
var tur int = 1
var runde int = 0
var game string
var gameOver bool
var takenSpots = []int{}

func main() {

	fmt.Println("Slik ser brettet ut: \n -------------  \n | 1 | 2 | 3 | \n -------------  \n | 4 | 5 | 6 | \n -------------  \n | 7 | 8 | 9 | \n -------------  \n Skriv inn tallet som representerer posisjonen du vil sette inn din første brikke")
	playGame()
}

func playGame() {

	fmt.Scanln(&lastAnswer)
	valueCheck()
	printGame()
	checkWinOrDraw()

}
func valueCheck() {
	checkValidity()
	for i := 0; i < len(takenSpots); i++ {
		if lastAnswer == takenSpots[i] {
			fmt.Println("Denne plassen er allerede tatt. Prøv igjen.")
			fmt.Scanln(&lastAnswer)
		}
	}
	takenSpots = append(takenSpots, lastAnswer)
}

func checkValidity() {
	if lastAnswer > 0 && lastAnswer < 10 {
	} else {
		fmt.Println("Du må velge et tall mellom 1 og 9 da det kun er 9 ruter på brettet. Prøv igjen")
		fmt.Scanln(&lastAnswer)
		checkValidity()
	}
}

func printGame() {

	game = ""
	var printSignal bool

	for refNum := 1; refNum <= 9; refNum++ {
		for i := 0; i < len(takenSpots); i++ {
			if refNum == takenSpots[i] {

				printSignal = true
				tur = i

				break
			} else {

				printSignal = false
			}
		}

		if printSignal {

			if tur%2 == 0 {
				game += "X"

			} else {
				game += "O"

			}
		} else {
			game += strconv.Itoa(refNum)
		}
	}

	fmt.Println("Slik ser brettet ut: \n -------------  \n |", game[0:1], "|", game[1:2], "|", game[2:3], "| \n -------------  \n |", game[3:4], "|", game[4:5], "|", game[5:6], "| \n -------------  \n |", game[6:7], "|", game[7:8], "|", game[8:9], "| \n -------------  ")
}

func checkWinOrDraw() {

	if game[0:1] == "X" && game[1:2] == "X" && game[2:3] == "X" || game[3:4] == "X" && game[4:5] == "X" && game[5:6] == "X" || game[6:7] == "X" && game[7:8] == "X" && game[8:9] == "X" || game[0:1] == "X" && game[3:4] == "X" && game[6:7] == "X" || game[1:2] == "X" && game[4:5] == "X" && game[7:8] == "X" || game[2:3] == "X" && game[5:6] == "X" && game[8:9] == "X" || game[0:1] == "X" && game[4:5] == "X" && game[8:9] == "X" || game[2:3] == "X" && game[4:5] == "X" && game[6:7] == "X" {

		fmt.Println("Spiller 1 vant spillet.")
		gameOver = true

	} else if game[0:1] == "O" && game[1:2] == "O" && game[2:3] == "O" || game[3:4] == "O" && game[4:5] == "O" && game[5:6] == "O" || game[6:7] == "O" && game[7:8] == "O" && game[8:9] == "O" || game[0:1] == "O" && game[3:4] == "O" && game[6:7] == "O" || game[1:2] == "O" && game[4:5] == "O" && game[7:8] == "O" || game[2:3] == "O" && game[5:6] == "O" && game[8:9] == "O" || game[0:1] == "O" && game[4:5] == "O" && game[8:9] == "O" || game[2:3] == "O" && game[4:5] == "O" && game[6:7] == "O" {

		fmt.Println("Spiller 2 vant spillet.")
		gameOver = true

	} else {

		runde++

		if runde == 9 {

			fmt.Println("Spillet er over.")
		} else if gameOver == true {

			fmt.Println("Spillet er over.")
		} else {
			playGame()
		}
	}

}
