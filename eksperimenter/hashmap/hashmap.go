package main

import "fmt"

func main() {
	board := map[int]string{
		1: "en",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
	}

	standard := map[int][]byte{
		1: {49},
		2: {50},
		3: {51},
		4: {52},
		5: {53},
		6: {54},
		7: {55},
		8: {56},
		9: {57},
	}

	i := 1
	k := 2

	fmt.Println(board[i+k])
}
