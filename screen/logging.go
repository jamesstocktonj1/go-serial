package screen

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type LoggingParagraph struct {
	Paragraph *widgets.Paragraph
	Width     int
	Send      bool
}

func CreateLogging() LoggingParagraph {
	p := widgets.NewParagraph()
	p.Title = "Logging"

	p.SetRect(0, 0, ScreenWidth, LOGGING_HEIGHT)

	return LoggingParagraph{
		Paragraph: p,
	}
}

func (p *LoggingParagraph) Write(data string) {
	p.Paragraph.Text = data
	p.Paragraph.TextStyle.Fg = ui.ColorWhite
	p.Render()
}

func (p *LoggingParagraph) Error(data string) {
	p.Paragraph.Text = data
	p.Paragraph.TextStyle.Fg = ui.ColorYellow
}

func (p *LoggingParagraph) Render() {
	ui.Render(p.Paragraph)
}
