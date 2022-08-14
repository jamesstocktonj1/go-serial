package main

import (
	"fmt"
	//"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func ListSerialPorts() []string {

	ports, err := enumerator.GetDetailedPortsList()

	if err != nil {
		fmt.Println("Error Getting Serial Ports")
	}

	var portList []string
	for _, port := range ports {
		portList = append(portList, port.Name)
	}

	return portList
}

func FormatListSerialPort() {
	
	ports, err := enumerator.GetDetailedPortsList()

	if err != nil {
		fmt.Println("Error Getting Serial Ports")
	}

	fmt.Println("Ports:")
	for _, port := range ports {
		fmt.Printf("  %s\n", port.Name)
	}
}

func IsSerialPort(port string, ports []string) bool {

	for _, p := range ports {

		if p == port {
			return true
		}
	}

	return false
}


func OpenSerialPort(flagSet *flag.FlagSet) {

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
}