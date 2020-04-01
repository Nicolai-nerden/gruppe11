package main

import "fmt"

func main() {
	start := 0
	goal := 20

	for goal != start {
		start++
		fmt.Println("Loop", start)
	}

}
