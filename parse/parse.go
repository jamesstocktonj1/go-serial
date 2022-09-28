package parse

import (
	"flag"
	"os"
	"strings"
	"go-serial/settings"
)

func HandleCommandParsing() *flag.FlagSet {

	flagSet := flag.NewFlagSet("Go Serial", flag.ContinueOnError)

	settings.Port = new(string)
	*settings.Port = "COMX"
	settings.Baud = flagSet.Int("baud", 9600, "select Baud Rate")

	settings.DataBits = flagSet.Int("data", 8, "select number of data bits")
	settings.StopBits = flagSet.Int("stop", 1, "select number of stop bits")
	settings.Parity = flagSet.String("parity", "None", "select parity type [None, Odd Even, Mark, Space]")

	settings.Verbose = flagSet.Bool("v", false, "verbose output e.g. connect/disconnect")
	settings.ListPorts = flagSet.Bool("l", false, "list available ports")
	settings.DefaultPort = flagSet.Bool("d", false, "select Default Serial Port")
	settings.PrintVersion = flagSet.Bool("version", false, "print software version")

	return flagSet
}

func ParseComPort() bool {

	// linux
	portFormat := "/dev/tty"

	// windows
	if os.Geteuid() == -1 {
		portFormat = "COM"
	}

	return strings.HasPrefix(*settings.Port, portFormat)
}

func ParseParity() bool {

	var parityValues = [5]string{"None", "Odd", "Even", "Mark", "Space"}

	for _, p := range parityValues {

		if *settings.Parity == p {
			return true
		}
	}

	return false
}
