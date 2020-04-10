package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/Nicolai-nerden/gruppe11/mp01/clientcommunication"
)

var agentList []net.Conn

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
	clientcommunication.ClientPrint(c, "Skriv inn navnet ditt: ")
	agentName := clientcommunication.ClientRead(c)
	agentNum := len(agentList)
	agentList = append(agentList, c)

	fmt.Printf("Tjener %s\n", c.RemoteAddr().String())
	for {

		msg := clientcommunication.ClientRead(c)

		if msg == "STOP" {

			break
		}

		printAll(c, agentName+": "+msg)
	}

	c.Close()
	agentList = append(agentList[:agentNum], agentList[agentNum+1:]...)
	fmt.Println(c.RemoteAddr().String(), "closed")
}

func printAll(c net.Conn, msg string) {
	for i := 0; i < len(agentList); i++ {
		if agentList[i] != c {
			clientcommunication.ClientPrintln(agentList[i], msg)
		}
	}
}
