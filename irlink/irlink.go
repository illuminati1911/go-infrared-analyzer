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

type ReadState struct {
	start           time.Time
	data            uint8
	previousData    uint8
	finishedReading bool
	message         []byte
}

func GetHandleToIRSource(signalSource SignalSource) chan Message {
	var feed chan Message
	go receiveFromSource(signalSource, feed)
	return feed
}

func receiveFromSource(source SignalSource, feed chan Message) {
	state := ReadState{
		start:           time.Now(),
		previousData:    1,
		finishedReading: true,
	}

	for {
		state.data = source.Read()
		if state.data == 0 && state.previousData == 1 {
			fallingEdge(&state)
		} else if state.data == 1 && state.previousData == 0 {
			risingEdge(&state)
		} else if state.data == 1 && state.previousData == 1 {
			noEdgeOne(&state)
		}
	}
}

func risingEdge(state *ReadState) {
	state.start = time.Now()
	state.previousData = 1
}

func fallingEdge(state *ReadState) {
	if time.Since(state.start) < time.Duration(time.Millisecond*3) {
		state.finishedReading = false
		if time.Since(state.start) < time.Duration(time.Millisecond) {
			state.message = append(state.message, 0)
		} else {
			state.message = append(state.message, 1)
		}
	}
	state.previousData = 0
}

func noEdgeOne(state *ReadState) {
	if !state.finishedReading && time.Since(state.start) > time.Duration(time.Millisecond*20) {
		fmt.Println(state.message)
		fmt.Println(len(state.message))
		fmt.Println("==============")
		if len(state.message) != validMessageLength {
			// send error
		} else {
			// Construct valid byte array and send
		}
		state.finishedReading = true
		state.message = []byte{}
	}
}
