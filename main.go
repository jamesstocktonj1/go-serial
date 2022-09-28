package main

import (
	"fmt"
	"os"
	"go-serial/parse"
	"go-serial/serial"
	"go-serial/settings"
)

func main() {

	flagSet := parse.HandleCommandParsing()
	flagSet.Parse(os.Args[1:])

	ser := serial.OpenSerialPort(flagSet)
	if ser == nil {
		os.Exit(1)
	} else {
		settings.SerialPort = ser
	}

	settings.SerialSync.Add(2)
	go serial.SerialInputHandler(settings.SerialPort)
	go serial.SerialOutputHandler(settings.SerialPort)

	settings.SerialSync.Wait()

	settings.SerialPort.Close()

	fmt.Printf("Serial Port: %s\n", *settings.Port)
}
