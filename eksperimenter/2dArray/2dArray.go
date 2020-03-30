package main

import "fmt"

func main() {
	winningCombos := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 4, 7}, {2, 5, 8}, {3, 6, 9}, {1, 5, 9}, {3, 5, 7}}

	fmt.Println(len(winningCombos))
	winningCombos = append(winningCombos, []int{100, 101, 102})

	fmt.Println(winningCombos[2][2])
}
