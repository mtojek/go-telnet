package client

import (
	"testing"
	"time"

	"net"

	"bytes"
	"net/http"

	"fmt"

	"bufio"
	"os"

	"github.com/mtojek/localserver"
	"github.com/stretchr/testify/assert"
)

func TestNewTelnetClient(t *testing.T) {
	// given
	ip := net.IPv4(1, 2, 3, 4)
	port := 5678
	timeout := 5 * time.Second
	options := newMockedTelnetClientOptions(ip.String(), uint64(port), timeout)

	expectedAddr := &net.TCPAddr{
		IP:   ip,
		Port: port,
	}

	// when
	sut := NewTelnetClient(options)

	// then
	assert := assert.New(t)
	assert.Equal(expectedAddr, sut.destination, "TCP destination should be defined.")
	assert.Equal(timeout, sut.responseTimeout, "Response timeout is different than defined.")
}

func TestProcessDataFromBufferedString(t *testing.T) {
	// given
	ip := net.IPv4(127, 0, 0, 1)
	port := 45678
	timeout := 2 * time.Second
	options := newMockedTelnetClientOptions(ip.String(), uint64(port), timeout)
	sut := NewTelnetClient(options)

	request := "GET /first HTTP/1.1\nHost: localhost\n\n"
	buffer := bytes.NewBuffer([]byte(request))
	localServer := localserver.NewLocalServer(fmt.Sprintf("%v:%d", ip.String(), port), "http")
	http.HandleFunc("/first", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("a_response\n"))
	})
	localServer.StartHTTP()

	// when
	sut.ProcessData(buffer, os.Stdout)

	// then
	localServer.Stop()
	http.DefaultServeMux = http.NewServeMux()
}

func TestProcessDataFromFile(t *testing.T) {
	// given
	ip := net.IPv4(127, 0, 0, 1)
	port := 45679
	timeout := 2 * time.Second
	options := newMockedTelnetClientOptions(ip.String(), uint64(port), timeout)
	sut := NewTelnetClient(options)

	fi, err := os.Open("../resources/input-data/localhost_1.bin")
	if err != nil {
		panic(err)
	}

	defer fi.Close()

	buffer := bufio.NewReader(fi)
	localServer := localserver.NewLocalServer(fmt.Sprintf("%v:%d", ip.String(), port), "http")
	http.HandleFunc("/second", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("b_response\n"))
	})
	localServer.StartHTTP()

	// when
	sut.ProcessData(buffer, os.Stdout)

	// then
	localServer.Stop()
	http.DefaultServeMux = http.NewServeMux()
}
