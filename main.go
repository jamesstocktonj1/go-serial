package main

import (
	"fmt"
	"os"
	//"flag"
)

func main() {

	flagSet := HandleCommandParsing()
	flagSet.Parse(os.Args[1:])

	ser := OpenSerialPort(flagSet)

	if ser == nil {
		os.Exit(1)
	} else {
		serialPort = ser
	}

	serialSync.Add(1)
	go SerialInputHandler(serialPort)

	serialSync.Wait()

	//serialPort.Close()

	fmt.Printf("Serial Port: %s\n", *port)
}
