package main

import (
	"fmt"
	"tx3server/mp01/tictactoeturnering"
	"math/rand"
	"net"
	"time"
)

func main() {
	fmt.Println("Serveren kjører.")

	l, err := net.Listen("tcp4", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	defer c.Close()
	tictactoeturnering.Turnering(c)
	fmt.Println("closed")
}
