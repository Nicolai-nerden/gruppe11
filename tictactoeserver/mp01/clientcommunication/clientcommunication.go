package clientcommunication

import (
	"bufio"
	"fmt"
	"net"

	"github.com/teamwork/reload"
)

// ClientRead leser input fra klienten
func ClientRead(c net.Conn) string {
	input, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println("En bruker terminerte økten sin. Må gjenoppstarte server...")
		reload.Exec()
	}

	return input[:len(input)-1]
}

// ClientPrintln skriver på klienten med linjeskift.
func ClientPrintln(c net.Conn, output string) {

	c.Write([]byte(output + "\n"))
}

// ClientPrint skriver på klient uten linjeskift.
func ClientPrint(c net.Conn, output string) {

	c.Write([]byte(output))
}
