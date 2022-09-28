package serial

import (
	"fmt"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
	"os"
	"flag"
	"go-serial/settings"
	"go-serial/parse"
	"go-serial/output"
	//"bufio"
)

func ListSerialPorts() []string {

	ports, err := enumerator.GetDetailedPortsList()

	if err != nil {
		fmt.Println("Error Getting Serial Ports")
	}

	var portList []string
	for _, port := range ports {
		portList = append(portList, port.Name)
	}

	return portList
}

func FormatListSerialPort() {
	
	ports, err := enumerator.GetDetailedPortsList()

	if err != nil {
		fmt.Println("Error Getting Serial Ports")
	}

	fmt.Println("Ports:")
	for _, port := range ports {

		portUsb := ""
		if port.IsUSB {
			portUsb = "USB"
		}
		fmt.Printf("  %s\t\t%s\n", port.Name, portUsb)
	}
}

func IsSerialPort(port string, ports []string) bool {

	for _, p := range ports {

		if p == port {
			return true
		}
	}

	return false
}


func OpenSerialPort(flagSet *flag.FlagSet) serial.Port {

	if *settings.PrintVersion {
		fmt.Printf("Version: %s\n", settings.Version)
		return nil
	}

	if *settings.ListPorts {
		FormatListSerialPort()
		return nil
	}
	
	if *settings.DefaultPort {
		portsList := ListSerialPorts()

		if len(portsList) == 0 {
			fmt.Println("Error: No serial devices connected")
			return nil
		} else {
			*settings.Port = portsList[0]
		}

	} else {

		if len(flagSet.Args()) == 0 {
			fmt.Println("Error: No serial device specified")
			return nil
		}
		
		*settings.Port = flagSet.Args()[0]

		if !parse.ParseComPort() {
			fmt.Print("Error: Invalid serial device, device should take form ")

			if os.Geteuid() == -1 {
				fmt.Println("COMx")
			} else {
				fmt.Println("/dev/tty")
			}
			return nil
		}
	}

	var localSetting = serial.Mode{*settings.Baud, *settings.DataBits, serial.NoParity, serial.OneStopBit}
	localPort, err := serial.Open(*settings.Port, &localSetting)

	if err != nil {
		fmt.Println("Error: Failed to open serial port")
		return nil
	}

	return localPort
}

func SerialInputHandler(ser serial.Port) {
	defer settings.SerialSync.Done()
	
	ser.ResetInputBuffer()
	buff := make([]byte, 100)

	for {

		n, err := ser.Read(buff)

		if err != nil {
			output.PrintLogging("Error: Unable to read Serial")
			break
		}

		if n == 0 {
			fmt.Println("\nEOF")
			break
		}

		output.PrintSimple(string(buff[:n]))
	}

	ser.Close()
}

func SerialOutputHandler(ser serial.Port) {
	defer settings.SerialSync.Done()

	//consoleReader := bufio.NewReader(os.Stdin)
	buff := make([]byte, 100)

	
	for {
		n, err := os.Stdin.Read(buff)
		//buff, err := consoleReader.ReadByte()
		output.PrintLogging(string(buff[:n-1]))

		if err != nil {
			output.PrintLogging("Error: Unable to read console terminal")
		}

		if n < 3 {
			output.PrintLogging("Error: No data read from console")
		}


		n, err = ser.Write(buff[:n-1])

		if err != nil {
			output.PrintLogging("Error: Unable to write serial")
		}
	}
}