package serial

import (
	"fmt"
	"go-serial/screen"
	"go.bug.st/serial"
)

const (
	//DEFAULT_SETTINGS = serial.Mode{BaudRate: 9600}
	BUFFER_SIZE = 100
)

type SerialPort struct {
	Port         serial.Port
	PortName     string
	PortSettings serial.Mode
	Online 		 bool
}

func CreateSerial(portName string) SerialPort {
	p := SerialPort {
		PortName: portName,
		Online: true,
	}

	var err error
	p.Port, err = serial.Open(p.PortName, &serial.Mode{BaudRate: 9600})
	if err != nil {
		errorMsg := fmt.Sprintf("Could not open port %s", p.PortName)
		screen.Logger.Error(errorMsg)
		p.Online = false
	}

	if p.Online {
		p.Port.ResetInputBuffer()
		p.Port.ResetOutputBuffer()
	}

	return p
}

/*
func ListSerialPorts() []string {

	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		screen.Logger.Error("Error getting port list")
	}

	var portList []string
	for _, port := range ports {
		portList = append(portList, port.Name)
	}

	return portList
}
*/

func (p *SerialPort) Read() string {
	if !p.Online {
		return ""
	}

	buff := make([]byte, BUFFER_SIZE)

	n, err := p.Port.Read(buff)
	if err != nil || n == 0 {
		errorMsg := fmt.Sprintf("Could not read port %s", p.PortName)
		screen.Logger.Error(errorMsg)
		p.Online = false
		return ""
	} else {
		p.Online = true
	}

	return string(buff[:n])
}

func (p *SerialPort) Write(data string) {
	if !p.Online {
		return
	}

	_, err := p.Port.Write([]byte(data))
	if err != nil {
		errorMsg := fmt.Sprintf("Could not write port %s", p.PortName)
		screen.Logger.Error(errorMsg)
		p.Online = false
	} else {
		p.Online = true
	}
}

func (p *SerialPort) Close() {
	if p.Online {
		p.Port.Close()
		p.Online = false
	}
}