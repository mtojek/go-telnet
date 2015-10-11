package main

import (
	"fmt"

	"github.com/mtojek/go-telnet/commandline"
)

type goTelnet struct{}

func newGoTelnet() *goTelnet {
	return new(goTelnet)
}

func (g *goTelnet) run() {
	commandLine := commandline.Read()
	fmt.Println(commandLine)
}
