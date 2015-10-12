package client

import "time"

type mockedTelnetClientOptions struct {
	host    string
	port    uint64
	timeout time.Duration
}

func newMockedTelnetClientOptions(host string, port uint64, timeout time.Duration) *mockedTelnetClientOptions {
	return &mockedTelnetClientOptions{host, port, timeout}
}

func (m *mockedTelnetClientOptions) Host() string {
	return m.host
}

func (m *mockedTelnetClientOptions) Port() uint64 {
	return m.port
}

func (m *mockedTelnetClientOptions) Timeout() time.Duration {
	return m.timeout
}
