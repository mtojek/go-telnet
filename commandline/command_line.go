package commandline

import "time"

// CommandLine type represents options read from command line arguments.
type CommandLine struct {
	host    string
	port    uint64
	timeout time.Duration
}

// Host method returns a given host.
func (c *CommandLine) Host() string {
	return c.host
}

// Port method returns a given port.
func (c *CommandLine) Port() uint64 {
	return c.port
}

// Timeout method returns a given server response timeout after EOF of input file.
func (c *CommandLine) Timeout() time.Duration {
	return c.timeout
}
