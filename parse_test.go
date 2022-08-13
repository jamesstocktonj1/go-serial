package main

import (
	"testing"
)

func TestHandleCommandParsing(t *testing.T) {

	testFlagSet := HandleCommandParsing()

	t.Run("Default", func(t *testing.T) {

		//test default state
		testFlagSet.Parse([]string{})

		if *port != "COMX" {
			t.Errorf("default port value expected COMX but got %s", *port)
		}
		
		if *baud != 9600 {
			t.Errorf("default baud value expected 9600 but got %d", *baud)
		}

		if *dataBits != 8 {
			t.Errorf("default data bits expected 8 but got %d", *dataBits)
		}

		if *stopBits != 1 {
			t.Errorf("default stop bits expected 1 but got %d", *stopBits)
		}

		if *parity != "None" {
			t.Errorf("default parity expected None but got %s", *parity)
		}

		if *verbose != false {
			t.Errorf("default verbose value expected false but got %t", *verbose)
		}
	})


	t.Run("PortTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-com", "COM5"})
		if *port != "COM5" {
			t.Errorf("expected port value COM5 but got %s", *port)
		}
	})

	t.Run("BaudTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-baud", "115200"})
		if *baud != 115200 {
			t.Errorf("expected baud value 115200 but got %d", *baud)
		}
	})

	t.Run("DataTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-data", "9"})
		if *dataBits != 9 {
			t.Errorf("expected data bits value 9 but got %d", *dataBits)
		}
	})

	t.Run("StopTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-stop", "2"})
		if *stopBits != 2 {
			t.Errorf("expected stop bits value 2 but got %d", *stopBits)
		}
	})

	t.Run("ParityTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-parity", "Odd"})
		if *parity != "Odd" {
			t.Errorf("expected parity Odd but got %s", *parity)
		}
	})
}