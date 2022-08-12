package main

import (
	"github.com/fatih/color"
)

type SerialStream struct {
	serialColor color.Attribute
	serialName string
}


//const portColors := [4]


func NewSerialStream(portName string, portColor color.Attribute) *SerialStream  {
	return &SerialStream { portColor, portName }
}