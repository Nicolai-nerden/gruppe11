package main

import (
	"armyboyz/mp02/supportedLangs"
	"bytes"
	"fmt"
	"strings"
)

func main() {
Start:
	var move string
	fmt.Println("Skriv inn trekket du ønsker å gjøre.")
	fmt.Scanln(&move)

	if len(move) == 0 { //hvis tom input
		fmt.Println("Tomt trekk. Prøv igjen.")
		goto Start
	}

	input := []byte(strings.Fields(move)[0]) //trimmer inputen og gjør den om til bytes
	//numberOfLangs := len(supportedLangs.ValidInputs)
	validNumbers := 9
	matchedLang := 0 // må starte på -1 for at første språket skal være index av 0
	matchedKey := 0

	for matchedKey <= validNumbers && !bytes.Equal(input, supportedLangs.ValidInputs[matchedLang][matchedKey]) { //for hver key (tall) i språket
		matchedKey++
		fmt.Println(input)
		fmt.Println("Keybytes:", supportedLangs.ValidInputs[matchedLang][matchedKey])
	}

	fmt.Println(supportedLangs.ValidInputs[matchedLang][matchedKey])
}
