package main

import (
	"flag"
	"os"
	"strings"
)

func HandleCommandParsing() *flag.FlagSet {

	flagSet := flag.NewFlagSet("Go Serial", flag.ContinueOnError)

	port = new(string)
	*port = "COMX"
	baud = flagSet.Int("baud", 9600, "select Baud Rate")

	dataBits = flagSet.Int("data", 8, "select number of data bits")
	stopBits = flagSet.Int("stop", 1, "select number of stop bits")
	parity = flagSet.String("parity", "None", "select parity type [None, Odd Even, Mark, Space]")

	verbose = flagSet.Bool("v", false, "verbose output e.g. connect/disconnect")
	listPorts = flagSet.Bool("l", false, "list available ports")
	defaultPort = flagSet.Bool("d", false, "select Default Serial Port")

	return flagSet
}

func ParseComPort() bool {

	// linux
	portFormat := "/dev/tty"

	// windows
	if os.Geteuid() == -1 {
		portFormat = "COM"
	}

	return strings.HasPrefix(*port, portFormat)
}

func ParseParity() bool {

	var parityValues = [5]string{"None", "Odd", "Even", "Mark", "Space"}

	for _, p := range parityValues {

		if *parity == p {
			return true
		}
	}

	return false
}
