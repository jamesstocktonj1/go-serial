package main

import ( //"os"
	//"go-serial/parse"
	//"go-serial/serial"
	//"go-serial/settings"
	"go-serial/screen"
)

func main() {

	screen.InitScreen()
	defer screen.CloseScreen()

	screen.SetSize(45, 35)

	screen.CreatePorts([]string {"COM1", "COM2"}, true)
	screen.Update()

	screen.MainLoop()
}
