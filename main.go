package main

import (
	"fmt"
	"os"
	//"flag"
)

func main() {

	flagSet := HandleCommandParsing()
	flagSet.Parse(os.Args[1:])

	fmt.Println(*port)
}