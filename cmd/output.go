package output

import (
	"github.com/fatih/color"
)


//const portColors := [4]


func getColorStream(index int) *color.Color {
	return color.New(color.FgCyan).Add(color.Underline)
}