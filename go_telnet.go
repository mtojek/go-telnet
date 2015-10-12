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
	telnetClient := g.createTelnetClient()
	telnetClient.ProcessData(os.Stdin)
}

func (g *goTelnet) createTelnetClient() *client.TelnetClient {
	commandLine := commandline.Read()
	telnetClient := client.NewTelnetClient(commandLine)
	return telnetClient
}
