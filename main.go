package main

import (
	"go-serial/screen"
	"go-serial/serial"
)

var (
	SerialList []serial.SerialPort
)

func SerialLoop(s serial.SerialPort, p screen.OutputParagraph) {

	for {
		buf := s.Read()
		
		p.SetAvailable(s.Online)
		p.AddText(buf)
	}
}

func ScreenLoop() {

}

func main() {
	serialList := []string {"COM8"}

	screen.InitScreen()
	defer screen.CloseScreen()

	screen.SetSize(screen.DEFAULT_WIDTH, screen.DEFUALT_HEIGHT)

	screen.CreatePorts(serialList, false)
	SerialList = append(SerialList, serial.CreateSerial(serialList[0]))

	screen.SelectPort(0)
	screen.Update()

	go SerialLoop(SerialList[0], screen.Outputs[0])

	screen.MainLoop()
}
