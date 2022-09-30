package screen

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

const (
	MAX_WIDTH      = 50
	MAX_HEIGHT     = 1000
	INPUT_HEIGHT   = 3
	LOGGING_HEIGHT = 4
)

var (
	screen_width  = 0
	screen_height = 0
	port_count    = 0
	screen_log    = false
	logger	LoggingParagraph
	outputs	[]OutputParagraph
	input 	InputParagraph
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

	logger = CreateLogging()
	for i, p := range ports {
		outputs = append(outputs, CreateOutput(p, i))
	}
	input = CreateInput()
}

func Update() {
	if screen_log {
		logger.Render()
	}

	for _, p := range outputs {
		p.Render()
	}
	input.Render()
}

func MainLoop() {
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			log.Fatal()
		case "<Resize>":
			Update()
		default:
			input.UpdateInput(e)
			Update()
		}

		if input.Send {
			outputs[0].AddText(input.ReadBuffer())
		}
	}
}

func CloseScreen() {
	ui.Close()
}
