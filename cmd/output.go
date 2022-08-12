package output

import (
	"github.com/fatih/color"
)


const portColors := [4]


func (index int) *Color {
	return color.New(color.FgCyan).Add(color.Underline)
}