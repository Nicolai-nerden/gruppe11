package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[1])

	fmt.Println(a + b)

}
