package main

import (
	"fmt"
	"os"
	//"flag"
)

func main() {

	flagSet := HandleCommandParsing()
	flagSet.Parse(os.Args[1:])

	if *listPorts {
		FormatListSerialPort()
		os.Exit(3)
	}
	
	if *defaultPort {
		portsList := ListSerialPorts()

		if len(portsList) == 0 {
			fmt.Println("Error: No serial devices connected")
			os.Exit(1)
		} else {
			*port = portsList[0]
		}

	} else {

		if len(flagSet.Args()) == 0 {
			fmt.Println("Error: No serial device specified")
			os.Exit(1)
		}
		
		*port = flagSet.Args()[0]

		if !ParseComPort() {
			fmt.Print("Error: Invalid serial device, device should take form ")

			if os.Geteuid() == -1 {
				fmt.Println("COMx")
			} else {
				fmt.Println("/dev/tty")
			}
			os.Exit(1)
		}
	}

	fmt.Printf("Serial Port: %s\n", *port)
}
