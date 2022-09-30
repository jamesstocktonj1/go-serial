package screen

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type InputParagraph struct {
	Paragraph *widgets.Paragraph
	Width     int
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
	case "<C-<Backspace>>":
		if len(p.Paragraph.Text) > 0 {
			p.Paragraph.Text = p.Paragraph.Text[:len(p.Paragraph.Text)-1]
		}
	case "<Space>":
		p.Paragraph.Text += " "
	default:
		if len(e.ID) == 1 {
			p.Paragraph.Text += e.ID
		}
	}
}

func (p *InputParagraph) ReadBuffer() string {
	buf := p.Paragraph.Text + "\n"
	p.Paragraph.Text = ""
	p.Render()

	return buf
}

func (p *InputParagraph) Render() {
	ui.Render(p.Paragraph)
}
