package main

import (
	"fmt"
	"os"

	"github.com/illuminati1911/go-infrared-analyzer/gpio"
	ir "github.com/illuminati1911/go-infrared-analyzer/infrared"
	"github.com/olekukonko/tablewriter"
)

func main() {
	err, source := gpio.Open(23)
	defer gpio.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	recvFromSource(source)
}

// Receive binary messages from source
// and parse them to NEC protocol.
//
func recvFromSource(source ir.SignalSource) {
	var irFeed chan ir.Message
	go ir.ReceiveIR(source, irFeed)
	for message := range irFeed {
		if message.isValid {
			nec, err := ir.GenerateNECMessage(message)
			if err != nil {
				fmt.Println(err)
				continue
			}
			renderNEC(nec)
		}
	}
}

// Render NEC protocol messages to stdout.
//
func renderNEC(nm ir.NECMessage) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Binary", "Bytes LSB", "Bytes MSB"})

	for _, v := range nm.ToStringTable() {
		table.Append(v)
	}
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render() // Send output
}
