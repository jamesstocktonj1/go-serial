package main

import (
	"testing"
	"github.com/fatih/color"
)


func TestNewSerialStream(t *testing.T)  {
	testColor := color.FgRed
	testName := "COMX"

	testStream := NewSerialStream(testName, testColor)

	if testStream.serialColor != testColor {
		t.Errorf("expected serial color FgRed")
	}

	if testStream.serialName != testName {
		t.Errorf("expected serial name %s but got %s", testName, testStream.serialName)
	}
}


func TestFormatSerialName(t *testing.T) {
	testColor := color.FgRed
	testName := "COMX"

	testStream := NewSerialStream(testName, testColor)

	formattedString := FormatSerialName(testStream)

	if formattedString != "[COMX]" {
		t.Errorf("expected formatted string [COMX] but got %s", formattedString)
	}
}


func TestFormatEndLine(t *testing.T) {
	
}



func TestPrint(t *testing.T) {

	if testing.Verbose() {

		PrintSimple("Hello World!")


		testColor := color.FgRed
		testName := "COMX"
		testColor2 := color.FgGreen
		testName2 := "COMY"

		testStream := NewSerialStream(testName, testColor)
		testStream2 := NewSerialStream(testName2, testColor2)

		PrintFormat("Hello World!", testStream)
		PrintFormat("Hello World!", testStream2)
	}
}
