package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// Liten netcatklient for serveren vår
func main() {
  conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "178.128.250.190", 8080))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go io.Copy(os.Stdout, conn)
  _, err = io.Copy(conn, os.Stdin)
}