package main

import (
	"fmt"
	"os"
	//"flag"
)

func main() {

	flagSet := HandleCommandParsing()
	flagSet.Parse(os.Args[1:])

	OpenSerialPort(flagSet)

	fmt.Printf("Serial Port: %s\n", *port)
}
