package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/illuminati1911/go-infrared-analyzer/gpio"
	ir "github.com/illuminati1911/go-infrared-analyzer/infrared"
	"github.com/olekukonko/tablewriter"
)

type config struct {
	Pin            int  `json:pin`
	EnableCHFormat bool `json:changhong_format`
}

func main() {
	c, err := readConfig()
	if err != nil {
		fmt.Println("Could not read config.json")
		fmt.Println(err)
		os.Exit(1)
	}
	err, source := gpio.Open(uint8(c.Pin))
	defer gpio.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	recvFromSource(source, c.EnableCHFormat)
}

// Read config.json file which contains gpio pin number
// and enabling flag for ChangHong IR format.
//
func readConfig() (config, error) {
	f, err := os.Open("config.json")
	defer f.Close()
	if err != nil {
		return config{}, err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return config{}, err
	}

	var c config
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return config{}, err
	}
	return c, nil
}

// Receive binary messages from source
// and parse them to NEC protocol.
//
func recvFromSource(source ir.SignalSource, useCH bool) {
	irFeed := make(chan ir.Message)
	go ir.ReceiveIR(source, irFeed, useCH)
	for message := range irFeed {
		nec, err := ir.GenerateNECMessage(message)
		if err != nil {
			fmt.Println(err)
			continue
		}
		renderNEC(nec)
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
	table.Render()
}
