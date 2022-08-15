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
	return ""
}

func PrintSimple(data string) {
	fmt.Print(data)
	fmt.Print(FormatEndLine())
}

func PrintFormat(data string, serial SerialStream) {
	forName := FormatSerialName(serial)

	color.Set(serial.serialColor)
	fmt.Print(forName)

	color.Unset()
	fmt.Print(" ")
	fmt.Print(data)
	fmt.Print(FormatEndLine())
}

func PrintLogging(log string) {

	color.Set(color.FgYellow)
	fmt.Print("[log]  ")

	color.Unset()
	fmt.Println(log)
}

func PrintSerialLogging(log string, serial SerialStream) {
	forName := FormatSerialName(serial)

	color.Set(color.FgYellow)
	fmt.Printf("%s ", forName)

	color.Unset()
	fmt.Println(log)
}
