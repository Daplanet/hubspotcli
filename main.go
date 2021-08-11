package main

import (
	"os"

	"github.com/denzuko/hubspotcli/cmd"
)
  //log "github.com/sirupsen/logrus"

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	} else {
     os.Exit(0)
	}
}
