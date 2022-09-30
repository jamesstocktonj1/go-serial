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

	screen.SetSize(screen.DEFAULT_WIDTH, screen.DEFUALT_HEIGHT)

	screen.CreatePorts([]string {"COM1", "COM2", "COM5"}, false)
	screen.SelectPort(0)
	screen.Update()

	screen.MainLoop()
}
