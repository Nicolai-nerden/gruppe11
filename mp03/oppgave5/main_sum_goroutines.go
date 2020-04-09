package main

import (
	"fmt"
	"modul3/oppgave3/sum"
	"modul3/oppgave5/overflowchannels"
	"os"
	"reflect"
	"strconv"
	"unicode"

	"github.com/pkg/profile"
)

var isDecimal bool
var firstInput bool

func main() { // Tar inn to variabler, sjekker om de innholder bokstaver og deretter konverterer til float64 eller en passende inttype.
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

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

	chInt8, chInt32, chUint32, chInt64 := make(chan bool), make(chan bool), make(chan bool), make(chan bool)

	go overflowchannels.Int8Overflow(total, chInt8)
	go overflowchannels.Int32Overflow(total, chInt32)
	go overflowchannels.Uint32Overflow(total, chUint32)
	go overflowchannels.Int64Overflow(total, chInt64)

	if !<-chInt8 {
		endTotal := sum.SumInt8(int8(a), int8(b))
		fmt.Println(endTotal)
		fmt.Print("Variable type: ")
		fmt.Println(reflect.TypeOf(endTotal).Kind())

	} else if !<-chInt32 {
		endTotal := sum.SumInt32(int32(a), int32(b))
		fmt.Println(endTotal)
		fmt.Print("Variable type: ")
		fmt.Println(reflect.TypeOf(endTotal).Kind())

	} else if !<-chUint32 {

		endTotal := sum.SumUint32(uint32(a), uint32(b))
		fmt.Println(endTotal)
		fmt.Print("Variable type: ")
		fmt.Println(reflect.TypeOf(endTotal).Kind())

	} else if !<-chInt64 {
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
