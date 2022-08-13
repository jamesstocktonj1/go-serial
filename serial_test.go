package main

import (
	"testing"
)

func TestIsSerialPort(t *testing.T) {

	testPorts := []string{"COM4", "COM5", "COM7", "COM10"}

	testPort := "COM4"
	testAvailable := IsSerialPort(testPort, testPorts)

	testPort2 := "COM6"
	testAvailable2 := IsSerialPort(testPort2, testPorts)

	if testAvailable != true {
		t.Errorf("expected availability to be false but got %t", testAvailable)
	}

	if testAvailable2 != false {
		t.Errorf("expected availability to be false but got %t", testAvailable2)
	}
}
