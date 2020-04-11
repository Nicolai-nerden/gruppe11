package main

import (
	"fmt"
	"github.com/teamwork/reload"
	"tx3server/selvlagdePakker/comm/multiplayerComm"
	"tx3server/selvlagdePakker/turnering/turneringMultiplayer"
	"tx3server/selvlagdePakker/turnering/turneringLocal"
	"net"
	"strings"
	"strconv"
)

var agentList []net.Conn
var agentNames []string
var started bool = false

// setter opp serveren og godkjenner forespørsler fra nye klienter. Og sender dem til handle connection
func main() {
	fmt.Println("Serveren kjører.")

	l, err := net.Listen("tcp4", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

// handleConnections Bestemmer hvordan hver klient behandles.
func handleConnection(c net.Conn) {
	if chooseMode(c) { // Spør om du vil spille lokalt
		turneringLocal.Turnering(c)
		c.Close()
		return
	} else if started {
		multiplayerComm.ClientPrintln(c, "\nEt spill er allerede i gang. Be verten starte på nytt eller prøv igjen når de er ferdige.")
		c.Close()
		return
	}
name: 
	multiplayerComm.ClientPrint(c, "\nSkriv inn navnet ditt: ")
	name := multiplayerComm.ClientRead(agentList, c)
	if len(name) == 0 {
		multiplayerComm.ClientPrint(c, "Du kan ikke ha tomt navn. Prøv igjen.\n")
		goto name
	}
	agentNames = append(agentNames, name)
	multiplayerComm.PrintAll(agentList, name + " ble med.")
	agentList = append(agentList, c)
	multiplayerComm.PrintAll(agentList, "Antall spillere med: "+ strconv.Itoa(len(agentList)))
	fmt.Printf("Tjener %s\n", c.RemoteAddr().String())
	if len(agentList) == 1 {
		start()
	} else {
		multiplayerComm.ClientPrintln(c, "Venter på at "+agentNames[0]+" skal starte spill...")
	}
}

// hosten får mulighet til å starte spillet. Tillater ikke spill med færre enn 2
func start(){
	for {
	//start:
		multiplayerComm.ClientPrintln(agentList[0], "\nDu er spillets vert. Skriv \"start\" Når spillet skal settes igang.")
		startSig := multiplayerComm.ClientRead(agentList, agentList[0])

		if len(agentList) == 1 {
			multiplayerComm.ClientPrintln(agentList[0], "\nDet må minst være 2 spillere for å spille multiplayer. \nVent på at en til har koblet seg til.")
		} else if strings.Fields(startSig)[0] == "start" {
			started = true
			turneringMultiplayer.Turnering(agentList, agentNames)
			end()
			reload.Exec() //restarter serveren, for å kunne starte ny turnering.
		}

	}
}
// sletter alle forbindelser
func end(){
	for i:=0; i < len(agentList); i++ {
		agentList[i].Close()
	}
}

func chooseMode(c net.Conn) bool {
	multiplayerComm.ClientPrintln(c, "\nSkriv \"lokal\" for å spille lokalt på maskinen din.\n TRYKK ENTER FOR MULTIPLAYER.")
	answer := multiplayerComm.ClientRead(agentList, c)

	if len(answer) == 0 {
		return false
	} else if answer == "lokal" {
		return true
	}
	return false
	
}
