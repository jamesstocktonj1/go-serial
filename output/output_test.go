package output

import (
	"github.com/fatih/color"
	"os"
	"strings"
	"testing"
)

func TestNewSerialStream(t *testing.T) {
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

func TestPrintVerbose(t *testing.T) {

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

func TestPrint(t *testing.T) {

	testInput, testOutput, _ := os.Pipe()

	defaultStdout := os.Stdout
	os.Stdout = testOutput

	testData := make([]byte, 25)

	t.Run("SimpleTest", func(t *testing.T) {
		PrintSimple("Hello World!")

		testInput.Read(testData)
		if !strings.Contains(string(testData), "Hello World!") {
			t.Errorf("expect Hello World! but got %s", testData)
		}
	})

	testColor := color.FgRed
	testName := "COMX"
	testStream := NewSerialStream(testName, testColor)

	t.Run("FormatTest", func(t *testing.T) {
		PrintFormat("Hello World!", testStream)

		testInput.Read(testData)
		if !strings.Contains(string(testData), "[COMX] Hello World!") {
			t.Errorf("expect [COMX] Hello World! but got %s", testData)
		}
	})

	os.Stdout = defaultStdout
}

func TestLoggingVerbose(t *testing.T) {

	if testing.Verbose() {

		PrintLogging("Hello World!")

		testColor := color.FgRed
		testName := "COMX"
		testStream := NewSerialStream(testName, testColor)

		PrintSerialLogging("Hello World!", testStream)
	}
}

func TestLogging(t *testing.T) {

	testInput, testOutput, _ := os.Pipe()

	defaultStdout := os.Stdout
	os.Stdout = testOutput

	testData := make([]byte, 25)

	t.Run("LoggingTest", func(t *testing.T) {
		PrintLogging("Hello World!")

		testInput.Read(testData)
		if !strings.Contains(string(testData), "[log]  Hello World!\n") {
			t.Errorf("expect Hello World! but got %s", testData)
		}
	})

	testColor := color.FgRed
	testName := "COMX"
	testStream := NewSerialStream(testName, testColor)

	t.Run("LoggingSerialTest", func(t *testing.T) {
		PrintSerialLogging("Hello World!", testStream)

		testInput.Read(testData)
		if !strings.Contains(string(testData), "[COMX] Hello World!\n") {
			t.Errorf("expect [COMX] Hello World! but got %s", testData)
		}
	})

	os.Stdout = defaultStdout
}
