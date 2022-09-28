package settings

import (
	"go.bug.st/serial"
	"sync"
)

const Version = "0.1.0"

var (
	Port        	*string
	Baud        	*int
	DataBits    	*int
	StopBits    	*int
	Parity      	*string
	Verbose     	*bool
	ListPorts   	*bool
	DefaultPort 	*bool
	PrintVersion	*bool
	SerialPort  	serial.Port
	SerialSync  	sync.WaitGroup
)
