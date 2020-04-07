package validInputs

import (
	"bytes"
	"tx3server/mp02/supportedLangs"
	"strings"
)

// CheckIfValid sjekker om tegnet skrevet inn et gyldig trekk, på en av de støttede skriftspråkene.
func CheckIfValid(move string) (int, bool) {

	validity := true
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
		validity = false
	}

	return matchedKey, validity
}
