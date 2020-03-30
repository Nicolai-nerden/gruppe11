package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	bytesToInt()

}

func printBytes() {
	for i := 1; i <= 9; i++ {
		fmt.Println([]byte(strconv.Itoa(i)))
	}
}

func printString() {

	fmt.Println([]byte("第一"))
}

func kindOf() {
	var move string
	validInputs := [][]byte{{49}, {50}, {51}, {52}, {53}, {54}, {55}, {56}, {57}}
	i := 0
	fmt.Scanln(&move)

	a := []byte(move)
	b := validInputs[i]
	c := reflect.TypeOf(a).Kind()
	d := reflect.TypeOf(b).Kind()

	fmt.Println(c)
	fmt.Println(d)
}

func moveAndValidate() {
	var move string
	validInputs := [][]byte{{49}, {50}, {51}, {52}, {53}, {54}, {55}, {56}, {57}}
	i := 0

	fmt.Println(len(validInputs))
	fmt.Scanln(&move)

	for (i < len(validInputs)) && ([]byte(move)[0] != validInputs[i][0]) {
		i++
	}

	fmt.Println(string(validInputs[i]))

}

func bytesToInt() {
	validInputs := [][]byte{{49}, {50}, {51}, {52}, {53}, {54}, {55}, {56}, {57}}

	a, _ := strconv.Atoi(string(validInputs[0]))
	b, _ := strconv.Atoi(string(validInputs[1]))

	fmt.Println(a + b)

}
