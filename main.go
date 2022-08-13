package main

import (
	"fmt"
	"flag"
)

func main() {

	HandleCommandParsing()
	flag.Parse()

	fmt.Println(*port)
}