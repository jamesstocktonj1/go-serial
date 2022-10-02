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
	Selected  bool
}

func CreateOutput(portName string, index int) OutputParagraph {
	p := widgets.NewParagraph()
	p.Title = portName

	pStart := 0
	if ScreenLog {
		pStart = LOGGING_HEIGHT
	}

	p.SetRect((ScreenWidth/PortCount) * index, pStart, (ScreenWidth / PortCount) * (index+1), ScreenHeight-INPUT_HEIGHT)

	pWidth := ScreenWidth / PortCount
	pHeight := ScreenHeight - (INPUT_HEIGHT+pStart)

	return OutputParagraph{
		Paragraph: p,
		Width: pWidth,
		Height: pHeight,
		Open: false,
		Selected: false,
	}
}

func (p *OutputParagraph) Select(s bool) {
	p.Selected = s
	p.SetAvailable(p.Open)
	p.Render()
}

func (p *OutputParagraph) SetAvailable(s bool) {
	p.Open = s
	if p.Selected {
		p.Paragraph.TitleStyle.Fg = ui.ColorWhite
	} else if s {
		p.Paragraph.TitleStyle.Fg = ui.ColorGreen
	} else {
		p.Paragraph.TitleStyle.Fg = ui.ColorRed
	}
}

func (p *OutputParagraph) SetText(data string) {
	p.Paragraph.Text = data
	p.Render()
}

func (p *OutputParagraph) ParseText() {
	lineArray := []string {}
	temp := ""
	
	for _, c := range p.Paragraph.Text {
		temp += string(c) 
		if c == '\n' {
			lineArray = append(lineArray, temp)
			temp = ""
		}
	}
	lineArray = append(lineArray, temp)

	if len(lineArray) > (p.Height-2) {
		lineArray = lineArray[len(lineArray)-(p.Height-2):]
	} else {
		return
	}

	p.Paragraph.Text = ""
	for _, l := range lineArray {
		p.Paragraph.Text += l
	}
}

func (p *OutputParagraph) AddText(data string) {
	p.Paragraph.Text += data

	p.ParseText()
	p.Render()
}

func (p *OutputParagraph) Render() {
	ui.Render(p.Paragraph)
}
