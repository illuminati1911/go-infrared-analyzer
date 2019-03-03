// Package for NEC infrared protocol constructing from stream
// and parsing from binary.

package protocol

import (
	"time"
)

// SignalSource is any source used as a source for IR data.
type SignalSource interface {
	Read() uint8
}

// Message is final format of the IR/NEC data received from SignalSource.
// Will be returned to invoker through received channel.
type Message struct {
	Data    string
	IsValid bool
}

// ReadState holds the state of the message constructing procedure
// when receiving data from SignalSource.
type ReadState struct {
	start           time.Time
	data            uint8
	previousData    uint8
	finishedReading bool
	message         string
}

// ReceiveIR function is used to initiate IR message construction process.
// SignalSource will provide logical (binary) changes from the source port/BUS
// and after each message is constructed, they will be sent through the `feed`
// channel receiced as second parameter.
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
	close(feed)
}

// Time will be recorded for each "Logical 1" received from signal source.
// According to NEC protocol, the actual data gets constructed following way
// based on the changes in the signal source:
//
// Logical '0': A 562.5µs pulse burst followed by a 562.5µs space
// Logical '1': A 562.5µs pulse burst followed by a 1.6875ms space
//
// While not fully following the protocol rules, we will get accurate enough results by just
// measuring the length of the space following the initial burst and simplifying the rule:
//
// Logical '1': space longer than 1.0ms
// Logical '0': space shorter than 1.0ms
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
		feed <- Message{Data: state.message, IsValid: isValidNECMessage(state.message)}
		state.finishedReading = true
		state.message = ""
	}
}
