package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jamesstockonj1/go-serial/output"
)

func main() {
	c := getColorStream(0)
	c.Println("Hello World!")
}