package client

import "time"

// Options interface defines which settings should be provided to the client.
type Options interface {
	Host() string
	Port() uint64
	Timeout() time.Duration
}
