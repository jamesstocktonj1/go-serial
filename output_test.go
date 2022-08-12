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
