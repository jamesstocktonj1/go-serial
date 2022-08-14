package main

import (
	"go.bug.st/serial"
)

var (
	port        *string
	baud        *int
	dataBits    *int
	stopBits    *int
	parity      *string
	verbose     *bool
	listPorts   *bool
	defaultPort *bool
	serialPort  serial.Port
)
