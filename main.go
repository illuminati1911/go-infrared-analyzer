package main

import (
	"fmt"
	"os"

	gpio "github.com/illuminati1911/go-infrared-analyzer/gpio"
	irlink "github.com/illuminati1911/go-infrared-analyzer/irlink"
)

func main() {
	err, pin := gpio.Open(23)
	defer gpio.Close()
	if err != nil {
		fmt.Println("Error opening GPIO")
		os.Exit(1)
	}
	for data := range irlink.GetHandleToIRSource(pin) {
		fmt.Println(data)
	}
}
