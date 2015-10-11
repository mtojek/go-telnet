package commandline

import "gopkg.in/alecthomas/kingpin.v2"

// Read method returns valid options read from command line args.
func Read() *CommandLine {
	host := kingpin.Arg("host", "Target host").Required().String()
	port := kingpin.Arg("port", "Target port").Required().Uint64()
	timeout := kingpin.Flag("timeout", "Byte receiving timeout after the input EOF occurs").Short('t').Default("1s").Duration()

	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("1.0").Author("Marcin Tojek")
	kingpin.CommandLine.Name = "go-telnet"
	kingpin.CommandLine.Help = "Read bytes from stdin and pass them to the remote host."

	kingpin.Parse()

	return &CommandLine{
		host:    *host,
		port:    *port,
		timeout: *timeout,
	}
}
