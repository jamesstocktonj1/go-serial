package screen

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

const (
	MAX_WIDTH      = 50
	MAX_HEIGHT     = 1000
	DEFAULT_WIDTH  = 100
	DEFUALT_HEIGHT = 35
	INPUT_HEIGHT   = 3
	LOGGING_HEIGHT = 4
)

var (
	screen_width  = 0
	screen_height = 0
	port_count    = 0
	port_select	  = 0
	screen_log    = false
	Logger	LoggingParagraph
	Outputs	[]OutputParagraph
	Input 	InputParagraph
)

func InitScreen() {
	err := ui.Init()
	if err != nil {
		log.Println(err)
	}

	screen_width, screen_height = ui.TerminalDimensions()
}

func SetSize(width int, height int) {
	screen_width = width
	screen_height = height
}

func CreatePorts(ports []string, verbose bool) {
	port_count = len(ports)
	screen_log = verbose

	Logger = CreateLogging()
	for i, p := range ports {
		Outputs = append(Outputs, CreateOutput(p, i))
	}
	Input = CreateInput()
}

func SelectPort(index int) {
	port_select = index % port_count

	for i, p := range Outputs {
		p.Select(i == port_select)
	}
}

func Update() {
	if screen_log {
		Logger.Render()
	}

	for _, p := range Outputs {
		p.Render()
	}
	Input.Render()
}

func MainLoop() {
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			log.Fatal()
		case "<Right>":
			SelectPort(port_select + 1)
		case "<Left>":
			SelectPort(port_select - 1)
		case "<Resize>":
			Update()
		default:
			Input.UpdateInput(e)
			Update()
		}

		if Input.Send {
			Outputs[port_select].AddText(Input.ReadBuffer())
		}
	}
}

func CloseScreen() {
	ui.Close()
}
