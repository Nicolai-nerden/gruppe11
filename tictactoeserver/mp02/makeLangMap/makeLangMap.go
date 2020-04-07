package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	makeLangMap()
}

// print bytes printer ut en map med et språk sine tall 1 - 9 i bytes versjoner slik at de kan brukes i tictactoe
func makeLangMap() {
	var nameOfLang string
	tableOfChars := make(map[int][]byte)
	fmt.Println("Hva vil du mappet til skriftspråket du legge til skal hete? (Helst på engelsk)")
	fmt.Print("Name of language: ")
	fmt.Scanln(&nameOfLang)

	fmt.Println("\nBruk en oversetter på nettet og copy/paste inn tegnet/tegnene for dette språkets tall en etter en.")
	fmt.Println("Husk å ikke ha andre tegn enn de du skal konvertere.")
	for i := 1; i <= 9; i++ {
		var input string
		fmt.Print(strconv.Itoa(i) + ": ")
		fmt.Scan(&input)
		inputTrimmed := strings.Fields(input)
		tableOfChars[i] = []byte(inputTrimmed[0])
	}

	fmt.Println("\nLim inn dette i golangfilen:")
	fmt.Println()
	fmt.Print("var " + nameOfLang + " = map[int][]byte{")
	for i := 1; i <= 9; i++ {
		fmt.Print("\n 	" + strconv.Itoa(i) + ": {")
		for g := 0; g < len(tableOfChars[i]); g++ {
			fmt.Print(tableOfChars[i][g])
			if g+1 != len(tableOfChars[i]) {
				fmt.Print(", ")
			}
		}
		fmt.Print("},")
	}
	fmt.Println("\n}")

}
