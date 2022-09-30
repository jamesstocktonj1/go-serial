package screen

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type OutputParagraph struct {
	Paragraph *widgets.Paragraph
	Width     int
	Height    int
	Open 	  bool
}

func CreateOutput(portName string, index int) OutputParagraph {
	p := widgets.NewParagraph()
	p.Title = portName

	pStart := 0
	if screen_log {
		pStart = LOGGING_HEIGHT
	}

	p.SetRect((screen_width/port_count) * index, pStart, (screen_width / port_count) * (index+1), screen_height-INPUT_HEIGHT)

	return OutputParagraph{
		Paragraph: p,
		Open: true,
	}
}

func (p *OutputParagraph) Select(s bool) {
	if s {
		p.Paragraph.TitleStyle.Fg = ui.ColorWhite
	} else {
		p.SetAvailable(p.Open)
	}
	p.Render()
}

func (p *OutputParagraph) SetAvailable(s bool) {
	p.Open = s
	if s {
		p.Paragraph.TitleStyle.Fg = ui.ColorGreen
	} else {
		p.Paragraph.TitleStyle.Fg = ui.ColorRed
	}
}

func (p *OutputParagraph) SetText(data string) {
	p.Paragraph.Text = data
	p.Render()
}

func (p *OutputParagraph) AddText(data string) {
	p.Paragraph.Text += data
	p.Render()
}

func (p *OutputParagraph) Render() {
	ui.Render(p.Paragraph)
}
