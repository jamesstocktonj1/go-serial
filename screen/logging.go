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

	p.SetRect(0, 0, screen_width, LOGGING_HEIGHT)

	return LoggingParagraph{
		Paragraph: p,
	}
}

func (p *LoggingParagraph) SetText(data string) {
	p.Paragraph.Text = data
	p.RenderLogging()
}

func (p *LoggingParagraph) RenderLogging() {
	ui.Render(p.Paragraph)
}
