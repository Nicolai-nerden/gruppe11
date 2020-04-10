package multiplayerComm

import (
	"bufio"
	"net"

	"github.com/teamwork/reload"
)

// ClientRead leser input fra klienten
func ClientRead(c net.Conn) string {
	input, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
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

// PrintAll Printer til alle klientene koblet til.
func PrintAll(agentList []net.Conn, msg string) {
	for i := 0; i < len(agentList); i++ {
		ClientPrintln(agentList[i], msg)
	}
}
