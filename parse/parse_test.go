package parse

import (
	"os"
	"testing"
	"go-serial/settings"
)

func TestHandleCommandParsing(t *testing.T) {

	testFlagSet := HandleCommandParsing()

	t.Run("Default", func(t *testing.T) {

		//test default state
		testFlagSet.Parse([]string{})

		if *settings.Baud != 9600 {
			t.Errorf("default baud value expected 9600 but got %d", *settings.Baud)
		}

		if *settings.DataBits != 8 {
			t.Errorf("default data bits expected 8 but got %d", *settings.DataBits)
		}

		if *settings.StopBits != 1 {
			t.Errorf("default stop bits expected 1 but got %d", *settings.StopBits)
		}

		if *settings.Parity != "None" {
			t.Errorf("default parity expected None but got %s", *settings.Parity)
		}

		if *settings.Verbose != false {
			t.Errorf("default verbose value expected false but got %t", *settings.Verbose)
		}

		if *settings.ListPorts != false {
			t.Errorf("default list port value expected false but got %t", *settings.ListPorts)
		}

		if *settings.DefaultPort != false {
			t.Errorf("default port select flag value expected false but got %t", *settings.DefaultPort)
		}
	})

	t.Run("BaudTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-baud", "115200"})
		if *settings.Baud != 115200 {
			t.Errorf("expected baud value 115200 but got %d", *settings.Baud)
		}
	})

	t.Run("DataTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-data", "9"})
		if *settings.DataBits != 9 {
			t.Errorf("expected data bits value 9 but got %d", *settings.DataBits)
		}
	})

	t.Run("StopTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-stop", "2"})
		if *settings.StopBits != 2 {
			t.Errorf("expected stop bits value 2 but got %d", *settings.StopBits)
		}
	})

	t.Run("ParityTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-parity", "Odd"})
		if *settings.Parity != "Odd" {
			t.Errorf("expected parity value Odd but got %s", *settings.Parity)
		}
	})

	t.Run("Verbose", func(t *testing.T) {
		testFlagSet.Parse([]string{"-v"})
		if *settings.Verbose != true {
			t.Errorf("expected verbose value true but got %t", *settings.Verbose)
		}
	})

	t.Run("ListPortTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-l"})
		if *settings.ListPorts != true {
			t.Errorf("expected port list value true but got %t", *settings.ListPorts)
		}
	})

	t.Run("DefaultTest", func(t *testing.T) {
		testFlagSet.Parse([]string{"-d"})
		if *settings.DefaultPort != true {
			t.Errorf("expected port select flag value true but got %t", *settings.DefaultPort)
		}
	})
}

func TestParseParity(t *testing.T) {

	*settings.Parity = "Even"
	testResult := ParseParity()
	if testResult != true {
		t.Errorf("expected Even to be a valid parity")
	}

	*settings.Parity = "Bad"
	testResult = ParseParity()
	if testResult != false {
		t.Errorf("expected Bad to be an invalid parity")
	}
}

func TestParseComPort(t *testing.T) {

	// linux
	*settings.Port = "/dev/ttyUSB0"

	//windows
	if os.Geteuid() == -1 {
		*settings.Port = "COM5"
	}

	testResult := ParseComPort()
	if testResult != true {
		t.Errorf("expected %s to be a valid Serial Port", *settings.Port)
	}

	// linux
	*settings.Port = "/dev/ttx"

	//windows
	if os.Geteuid() == -1 {
		*settings.Port = "COXX"
	}

	testResult = ParseComPort()
	if testResult != false {
		t.Errorf("expected %s to not be a valid Serial Port", *settings.Port)
	}
}
