package irlink

import (
	"fmt"
	"time"
)

const validMessageLength int = 120

type SignalSource interface {
	Read() uint8
}

type Message struct {
	data    [15]byte
	isValid bool
}

func GetHandleToIRSource(signalSource SignalSource) chan Message {
	var feed chan Message
	go receiveFromSource(signalSource, feed)
	return feed
}

func receiveFromSource(source SignalSource, feed chan Message) {
	var start = time.Now()
	var res uint8
	previousValue := 1
	finishedReading := true
	var message []byte
	for {
		res = source.Read()
		if res == 0 && previousValue == 1 {
			if time.Since(start) < time.Duration(time.Millisecond*3) {
				finishedReading = false
				if time.Since(start) < time.Duration(time.Millisecond) {
					message = append(message, 0)
				} else {
					message = append(message, 1)
				}
			}
			previousValue = 0
		} else if res == 1 && previousValue == 0 {
			start = time.Now()
			previousValue = 1
		} else if res == 1 && previousValue == 1 {
			if !finishedReading && time.Since(start) > time.Duration(time.Millisecond*20) {
				fmt.Println(message)
				fmt.Println(len(message))
				fmt.Println("==============")
				if len(message) != validMessageLength {
					// send error
				} else {
					// Construct valid byte array and send
				}
				finishedReading = true
				message = []byte{}
			}
		}
	}
}
