package main

import (
	"os"

	"github.com/mtojek/go-telnet/client"
	"github.com/mtojek/go-telnet/commandline"
)

type goTelnet struct{}

func newGoTelnet() *goTelnet {
	return new(goTelnet)
}

func (g *goTelnet) run() {
	commandLine := commandline.Read()
	telnetClient := client.NewTelnetClient(commandLine)
	telnetClient.ProcessData(os.Stdin)
}
