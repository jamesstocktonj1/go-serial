package screen

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type InputParagraph struct {
	Paragraph *widgets.Paragraph
	Width     int
	Send      bool
}

func CreateInput() InputParagraph {
	p := widgets.NewParagraph()
	p.Title = "Console"

	p.SetRect(0, screen_height-INPUT_HEIGHT, screen_width, screen_height)

	return InputParagraph{
		Paragraph: p,
	}
}

func (p *InputParagraph) UpdateInput(e ui.Event) {
	switch e.ID {
	case "<C-c>":
	case "<Resize>":
	case "<C-<Backspace>>":
		if len(p.Paragraph.Text) > 1 {
			p.Paragraph.Text = p.Paragraph.Text[:len(p.Paragraph.Text)-1]
		}
	case "<Space>":
		p.Paragraph.Text += " "
	case "<Enter>":
		p.Send = true
	default:
		p.Paragraph.Text += e.ID
	}
}

func (p *InputParagraph) ReadBuffer() string {
	buf := p.Paragraph.Text + "\n"
	p.Send = false
	p.Paragraph.Text = ""

	return buf
}

func (p *InputParagraph) RenderInput() {
	ui.Render(p.Paragraph)
}
