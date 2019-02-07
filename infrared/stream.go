package protocol

import (
	"time"
)

type SignalSource interface {
	Read() uint8
}

type Message struct {
	data    string
	isValid bool
}

type ReadState struct {
	start           time.Time
	data            uint8
	previousData    uint8
	finishedReading bool
	message         string
}

func ReceiveIR(source SignalSource, feed chan Message) {
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
			noEdgeOne(&state, feed)
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
			state.message += "0"
		} else {
			state.message += "1"
		}
	}
	state.previousData = 0
}

func noEdgeOne(state *ReadState, feed chan Message) {
	if !state.finishedReading && time.Since(state.start) > time.Duration(time.Millisecond*20) {
		if isValidNECMessage(state.message) {
			feed <- Message{data: state.message, isValid: false}
		} else {
			feed <- Message{data: state.message, isValid: true}
		}
		state.finishedReading = true
		state.message = ""
	}
}
