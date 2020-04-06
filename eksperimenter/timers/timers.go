package main

import (
	"time"

	"github.com/teamwork/reload"
)

func main() {

	timer1 := time.NewTimer(1 * time.Hour)

	<-timer1.C
	reload.Exec()
}
