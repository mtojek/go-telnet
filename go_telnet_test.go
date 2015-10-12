package main

import (
	"testing"

	"github.com/mtojek/go-telnet/commandline"
	"github.com/stretchr/testify/assert"
)

func TestCreateTelnetClient(t *testing.T) {
	// given
	commandline.SetCommandLineArgs("localhost", "80")
	sut := newGoTelnet()

	// when
	result := sut.createTelnetClient()

	// then
	assert := assert.New(t)
	assert.NotNil(result)
}
