package main

import (
	"fmt"
	"os"

	gpio "github.com/illuminati1911/go-infrared-analyzer/gpio"
	irlink "github.com/illuminati1911/go-infrared-analyzer/irlink"
)

func main() {
	// Get signal source from Raspberry Pi GPIO port.
	err, pin := gpio.Open(23)
	defer gpio.Close()
	if err != nil {
		fmt.Println("Error opening GPIO")
		os.Exit(1)
	}

	// Get Changhong IR message chunks
	var irFeed chan irlink.Message
	go irlink.GetHandleToIRSource(pin, irFeed)
	for data := range irFeed {
		fmt.Println(data)
	}
}
