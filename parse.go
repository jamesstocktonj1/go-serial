package main

import (
	"fmt"
	"flag"
)


func HandleCommandParsing() *flag.FlagSet {

	flagSet := flag.NewFlagSet("Example FlagSet", flag.ContinueOnError)

	port = flagSet.String("com", "COMX", "select Serial Port")
	baud = flagSet.Int("baud", 9600, "select Baud Rate")

	dataBits = flagSet.Int("data", 8, "select number of data bits")
	stopBits = flagSet.Int("stop", 1, "select number of stop bits")
	parity = flagSet.String("parity", "None", "select parity type [None, Odd Even, Mark, Space]")

	verbose = flagSet.Bool("v", false, "verbose output e.g. connect/disconnect")
	listPorts = flagSet.Bool("ls", false, "list available ports")
	defaultPort = flagSet.Bool("d", false, "select Default Serial Port")

	return flagSet
}


func ParseComPort() {

}

func ParseParity() {

}