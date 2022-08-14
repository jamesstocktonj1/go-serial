package main

import (
	"testing"
	"os"
	"strings"
	"flag"
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

func TestFormatListSerialPort(t *testing.T) {

	testInput, testOutput, _ := os.Pipe()

	defaultStdout := os.Stdout
	os.Stdout = testOutput

	testData := make([]byte, 25)

	FormatListSerialPort()
	testInput.Read(testData)

	if !strings.Contains(string(testData), "Ports:\n") {
		t.Errorf("expected list serial to start with Ports: but got %s", testData)
	}

	os.Stdout = defaultStdout
}

func TestOpenSerialPort(t *testing.T) {

	testInput, testOutput, _ := os.Pipe()

	defaultStdout := os.Stdout
	os.Stdout = testOutput

	testData := make([]byte, 35)


	testFlagSet := flag.NewFlagSet("TestFlagSet", flag.ContinueOnError)
	testFlagSet.Parse([]string{})

	t.Run("TestListPorts", func(t *testing.T) {
		*listPorts = true
		*defaultPort = false

		OpenSerialPort(testFlagSet)
		testInput.Read(testData)

		if !strings.Contains(string(testData), "Ports:\n") {
			t.Errorf("expected list serial to start with Ports: but got %s", testData)
		}
	})

	t.Run("TestNoDevice", func(t *testing.T) {
		*listPorts = false
		*defaultPort = false

		OpenSerialPort(testFlagSet)
		testInput.Read(testData)

		if !strings.Contains(string(testData), "Error: No serial device specified\n") {
			t.Errorf("expected no device specified to display error but got %s", testData)
		}
	})

	t.Run("TestBadDevice", func(t *testing.T) {

		if os.Geteuid() == -1 {
			testFlagSet.Parse([]string{"com1"})
		} else {
			testFlagSet.Parse([]string{"/dev/ttx1"})
		}
		
		OpenSerialPort(testFlagSet)
		testInput.Read(testData)

		if !strings.Contains(string(testData), "Error: Invalid serial device,") {
			t.Errorf("expected invalid device to display error but got %s", testData)
		}
	})

	os.Stdout = defaultStdout
}