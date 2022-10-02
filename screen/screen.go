package screen

import (
	"log"
	"time"

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
	ScreenWidth  = 0
	ScreenHeight = 0
	PortCount    = 0
	PortSelect	  = 0
	ScreenLog    = false
	Logger	LoggingParagraph
	Outputs	[]OutputParagraph
	Input 	InputParagraph
)

func InitScreen() {
	err := ui.Init()
	if err != nil {
		log.Println(err)
	}

	ScreenWidth, ScreenHeight = ui.TerminalDimensions()
}

func SetSize(width int, height int) {
	ScreenWidth = width
	ScreenHeight = height
}

func CreatePorts(ports []string, verbose bool) {
	PortCount = len(ports)
	ScreenLog = verbose

	Logger = CreateLogging()
	for i, p := range ports {
		Outputs = append(Outputs, CreateOutput(p, i))
	}
	Input = CreateInput()
}

func SelectPort(index int) {
	PortSelect = index % PortCount

	for i, p := range Outputs {
		p.Select(i == PortSelect)
	}
}

func Update() {
	if ScreenLog {
		Logger.Render()
	}

	for _, p := range Outputs {
		p.Render()
	}
	Input.Render()
}

func MainLoop() {
	t := time.Tick(time.Millisecond)
	uiEvents := ui.PollEvents()
	for range t{
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			log.Fatal()
		case "<C-a>":
			ScreenLog = !ScreenLog
			Update()
		case "<Right>":
			SelectPort(PortSelect + 1)
		case "<Left>":
			SelectPort(PortSelect - 1)
		case "<Enter>":
			Outputs[PortSelect].AddText(Input.ReadBuffer())
		case "<Resize>":
			Update()
		default:
			Input.UpdateInput(e)
			Update()
		}
	}
}

func CloseScreen() {
	ui.Close()
}
