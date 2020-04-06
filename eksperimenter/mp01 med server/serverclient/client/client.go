package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Nummer som skal oversettes: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)
	}
}
