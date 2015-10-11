package client

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const defaultBufferSize = 4096

// TelnetClient represents a TCP client which is responsible for writing input data and printing response.
type TelnetClient struct {
	destination     *net.TCPAddr
	responseTimeout time.Duration
}

// NewTelnetClient method creates new instance of TCP client.
func NewTelnetClient(options Options) *TelnetClient {
	tcpAddr := createTCPAddr(options)
	resolved := resolveTCPAddr(tcpAddr)

	return &TelnetClient{
		destination:     resolved,
		responseTimeout: options.Timeout(),
	}
}

func createTCPAddr(options Options) string {
	var buffer bytes.Buffer
	buffer.WriteString(options.Host())
	buffer.WriteByte(':')
	buffer.WriteString(fmt.Sprintf("%d", options.Port()))
	return buffer.String()
}

func resolveTCPAddr(addr string) *net.TCPAddr {
	resolved, error := net.ResolveTCPAddr("tcp", addr)
	if nil != error {
		log.Fatalf("Error occured while resolving TCP address \"%v\": %v", addr, error)
	}

	return resolved
}

// ProcessData method processes data: reads from input and writes to output.
func (t *TelnetClient) ProcessData(inputData io.Reader) {
	connection, error := net.DialTCP("tcp", nil, t.destination)
	if nil != error {
		log.Fatalf("Error occured while connecting to address \"%v\": %v", t.destination.String(), error)
	}

	defer connection.Close()
	t.writeInputData(connection, inputData)
	t.printServerResponse(connection)
}

func (t *TelnetClient) writeInputData(connection *net.TCPConn, dataReader io.Reader) {
	reader := bufio.NewReader(dataReader)
	b, error := reader.ReadByte()

	for nil == error {
		connection.Write([]byte{b})
		b, error = reader.ReadByte()
	}

	t.assertEOF(error)
}

func (t *TelnetClient) printServerResponse(connection *net.TCPConn) {
	var someRead bool
	receivedChannel := make(chan []byte, 1)

	go t.readData(connection, receivedChannel)

	for {
		select {
		case data := <-receivedChannel:
			someRead = true
			fmt.Printf("%v", string(data))
		case <-time.After(t.responseTimeout):
			if !someRead {
				log.Println("Nothing read. Timeout?")
			}
			return
		}
	}
}

func (t *TelnetClient) readData(connection *net.TCPConn, received chan<- []byte) {
	buffer := make([]byte, defaultBufferSize)
	var error error
	var n int

	for nil == error {
		n, error = connection.Read(buffer)
		received <- buffer[:n]
	}

	t.assertEOF(error)
}

func (t *TelnetClient) assertEOF(error error) {
	if "EOF" != error.Error() {
		log.Fatalf("Error occured while writing to TCP socket: %v", error)
	}
}
