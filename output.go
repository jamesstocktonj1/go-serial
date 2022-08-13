package main

import (
	"fmt"
	"github.com/fatih/color"
)

type SerialStream struct {
	serialColor color.Attribute
	serialName  string
}

func NewSerialStream(portName string, portColor color.Attribute) SerialStream {
	return SerialStream{portColor, portName}
}

func FormatSerialName(serial SerialStream) string {
	return "[" + serial.serialName + "]"
}

func FormatEndLine() string {
	return "\n"
}

func PrintSimple(data string) {
	fmt.Print(data)
	fmt.Print(FormatEndLine())
}

func PrintFormat(data string, serial SerialStream) {
	col := color.New(serial.serialColor)
	forName := FormatSerialName(serial)

	col.Print(forName)
	fmt.Print(" ")
	fmt.Print(data)
	fmt.Print(FormatEndLine())
}
