package main

import (
	"fmt"
	"os"
	//"flag"
)

const version = "0.1.0"

func main() {

	flagSet := HandleCommandParsing()
	flagSet.Parse(os.Args[1:])

	ser := OpenSerialPort(flagSet)
	if ser == nil {
		os.Exit(1)
	} else {
		serialPort = ser
	}

	serialSync.Add(2)
	go SerialInputHandler(serialPort)
	go SerialOutputHandler(serialPort)

	serialSync.Wait()

	//serialPort.Close()

	fmt.Printf("Serial Port: %s\n", *port)
}
