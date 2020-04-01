package main

import (
	"fmt"
	"gruppe11/mp03/overflow"
	"gruppe11/mp03/sum"
	"os"
	"reflect"
	"strconv"
	"unicode"
	//"github.com/pkg/profile"
)

var isDecimal bool
var firstInput bool

func main() { // Tar inn to variabler, sjekker om de innholder bokstaver og deretter konverterer til float64 eller en passende inttype.
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	arg1 := os.Args[1]
	arg2 := os.Args[2]
	isDecimal = false
	firstInput = true

	a := correctValues(arg1)
	b := correctValues(arg2)

	if isDecimal {
		aConverted := convertToFloat(a)
		bConverted := convertToFloat(b)
		sum := sum.SumFloat64(float64(aConverted), float64(bConverted))
		fmt.Println(sum)
		fmt.Print("Variable type: ")
		fmt.Println(reflect.TypeOf(sum).Kind())

	} else {

		aConverted := convertToInt(a)
		bConverted := convertToInt(b)

		adjustIntType(aConverted, bConverted)
	}
}

func adjustIntType(a int, b int) { //finner ut hvilken inttype som passer. Printer derreter ut resultatet og inntypen ved hjelp av "sum"-pakken.

	total := a + b

	if !overflow.Int8Overflow(total) {
		endTotal := sum.SumInt8(int8(a), int8(b))
		fmt.Println(endTotal)
		fmt.Print("Variable type: ")
		fmt.Println(reflect.TypeOf(endTotal).Kind())

	} else if !overflow.Int32Overflow(total) {
		endTotal := sum.SumInt32(int32(a), int32(b))
		fmt.Println(endTotal)
		fmt.Print("Variable type: ")
		fmt.Println(reflect.TypeOf(endTotal).Kind())

	} else if !overflow.Uint32Overflow(total) {

		endTotal := sum.SumUint32(uint32(a), uint32(b))
		fmt.Println(endTotal)
		fmt.Print("Variable type: ")
		fmt.Println(reflect.TypeOf(endTotal).Kind())

	} else if !overflow.Int64Overflow(total) {

		endTotal := sum.SumInt64(int64(a), int64(b))
		fmt.Println(endTotal)
		fmt.Print("Variable type: ")
		fmt.Println(reflect.TypeOf(endTotal).Kind())
	} else {
		fmt.Println("There's a glitch in the matrix")
	}

}

func convertToFloat(x string) float64 { //converterer
	y, _ := strconv.ParseFloat(x, 64)
	return y
}

func correctValues(input string) string { //teste resultatet helt til det ikke er en bokstav og returnerer resultatet i form av en String.
	var resultat = input
	if !erGyldig(resultat) {
		if firstInput {
			fmt.Println("Denne første inputen er ikke gyldig. \n Prøv igjen")
		} else {
			fmt.Println("Denne andre inputen er ikke gyldig. \n Prøv igjen")
		}
		fmt.Scanln(&resultat)
		correctValues(resultat)
	}
	firstInput = false
	return resultat
}

func convertToInt(conv string) int { //konverterer string til int og returnerer dette.
	var x, _ = strconv.Atoi(conv)
	return x
}

func erGyldig(s string) bool { //Tester om inputen er et riktig inntastet siffer/desimal. Returnerer true eller false.
	for _, r := range s {
		if r == 46 { // Hvis den finner et punktum betyr dette at det er et desimaltall.
			isDecimal = true //Setter til at vi skal regne med desimaler.
		} else if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}
