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

	if len(ports) == 0 {
		fmt.Println("No Serial Devices")
	}

	var portList []string;

	for _, port := range ports {
		fmt.Printf("%v\n", port.Name)
		portList = append(portList, port.Name)
	}

	return portList
}


func IsSerialPort(port string, ports []string) bool {

	for _, p := range ports {
		
		if p == port {
			return true
		}
	}

	return false
}