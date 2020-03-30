package main

import (
	"fmt"
	"time"
)

func main() {

	var move string

	start := time.Now()

	fmt.Scanln(&move)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(elapsed.Milliseconds())      //konverterer til ms
	fmt.Println(int(elapsed.Milliseconds())) //konverterer fra int64 til int
}
