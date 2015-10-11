package commandline

import "time"

// CommandLine type represents options read from command line arguments.
type CommandLine struct {
	Timeout time.Duration
	Host    string
	Port    uint64
}
