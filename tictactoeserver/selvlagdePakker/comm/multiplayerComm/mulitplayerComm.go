package multiplayerComm

import (
	"bufio"
	"net"

	"github.com/teamwork/reload"
)

// ClientRead leser input fra klienten
func ClientRead(agentList []net.Conn, c net.Conn) string {
	input, err := bufio.NewReader(c).ReadString('\n')
	if input == "^C\n" || err != nil {
		PrintAll(agentList, "Noen koblet fra. Restarter server")
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
