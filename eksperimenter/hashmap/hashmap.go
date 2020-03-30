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
	i := 1
	k := 2

	fmt.Println(board[i+k])
}
