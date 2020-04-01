package validInputs

import (
	"bytes"
	"fmt"
	"gruppe11/mp02/supportedLangs"
	"strings"
)

// CheckIfValid sjekker om tegnet skrevet inn et gyldig trekk, på en av de støttede skriftspråkene.
func CheckIfValid() int {
Start:
	var move string
	fmt.Scanln(&move)

	if len(move) == 0 { //hvis tom input
		fmt.Println("Tomt trekk. Prøv igjen.")
		goto Start
	}

	input := []byte(strings.Fields(move)[0]) //trimmer inputen og gjør den om til bytes
	numberOfLangs := len(supportedLangs.ValidInputs)
	validNumbers := 9
	matchedLang := 0
	matchedKey := 0

	for matchedLang < numberOfLangs && !bytes.Equal(input, supportedLangs.ValidInputs[matchedLang][matchedKey]) { // for hvert språk
		matchedKey = 1                                                                                               // resetter key ved nytt språk
		for matchedKey <= validNumbers && !bytes.Equal(input, supportedLangs.ValidInputs[matchedLang][matchedKey]) { //for hver key (tall) i språket
			matchedKey++
		}
		if bytes.Equal(input, supportedLangs.ValidInputs[matchedLang][matchedKey]) {
			break
		}
		matchedLang++
	}

	if matchedLang >= numberOfLangs {
		fmt.Println("Ugyldig input. Prøv igjen")
		goto Start
	}

	return matchedKey
}
