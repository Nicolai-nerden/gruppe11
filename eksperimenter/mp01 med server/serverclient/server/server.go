package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

var i = 1
var sum int

func main() {

	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8081")

	conn, _ := ln.Accept()

	for {
		input, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(input))

		if i == 1 {
			a, _ := strconv.Atoi(string(input[0]))
			sum += a
			conn.Write([]byte("Første nummer mottatt. Nummer: " + input))
			i++
		} else if i == 2 {
			b, _ := strconv.Atoi(string(input[0]))
			sum += b
			// send new string back to client
			conn.Write([]byte("Summen av tallene: " + strconv.Itoa(sum) + "\n"))
			i = 1
			sum = 0
		}
	}
}
