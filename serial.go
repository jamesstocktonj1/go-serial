package main

import (
	"fmt"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
	"os"
	"flag"
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

		portUsb := ""
		if port.IsUSB {
			portUsb = "USB"
		}
		fmt.Printf("  %s\t\t%s\n", port.Name, portUsb)
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


func OpenSerialPort(flagSet *flag.FlagSet) serial.Port {

	if *listPorts {
		FormatListSerialPort()
		return nil
	}
	
	if *defaultPort {
		portsList := ListSerialPorts()

		if len(portsList) == 0 {
			fmt.Println("Error: No serial devices connected")
			return nil
		} else {
			*port = portsList[0]
		}

	} else {

		if len(flagSet.Args()) == 0 {
			fmt.Println("Error: No serial device specified")
			return nil
		}
		
		*port = flagSet.Args()[0]

		if !ParseComPort() {
			fmt.Print("Error: Invalid serial device, device should take form ")

			if os.Geteuid() == -1 {
				fmt.Println("COMx")
			} else {
				fmt.Println("/dev/tty")
			}
			return nil
		}
	}

	var localSetting = serial.Mode{*baud, *dataBits, serial.NoParity, serial.OneStopBit}
	localPort, err := serial.Open(*port, &localSetting)

	if err != nil {
		fmt.Println("Error: Failed to open serial port")
		return nil
	}

	return localPort
}