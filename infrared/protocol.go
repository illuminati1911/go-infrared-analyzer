// Package for NEC infrared protocol constructing from stream
// and parsing from binary.

package protocol

import (
	"errors"
	"strconv"

	"github.com/illuminati1911/go-infrared-analyzer/utils"
)

const validChangHongMessageLength int = 120

// NECMessage consisting of 15 bytes.
// Each byte is transformed to a binary string and
// appended to the chunks slice.
type NECMessage struct {
	chunks []string
}

// BitOrder type alias to tell whether to parse the binary string to
// bytes with Least or Most Significant Bit first.
type BitOrder bool

const (
	LSB BitOrder = true
	MSB BitOrder = false
)

// Bytes returns the byte presentation of each binary string as string.
// This should be only used for visual demonstration purposes.
func (m NECMessage) Bytes(b BitOrder) []string {
	bytes := make([]string, 0)
	for _, chunk := range m.chunks {
		orderedChunk := chunk
		if b == LSB {
			orderedChunk = utils.Reverse(orderedChunk)
		}
		n, err := strconv.ParseUint(orderedChunk, 2, 8)
		if err == nil {
			bytes = append(bytes, strconv.FormatUint(n, 10))
		}
	}
	return bytes
}

// ToStringTable generates 2D array of the NEC message for simplyfying render in
// UI or shell.
func (m NECMessage) ToStringTable() [][]string {
	var table [][]string
	lsb := m.Bytes(LSB)
	msb := m.Bytes(MSB)
	for i, bin := range m.chunks {
		var s []string
		s = append(s, bin)
		s = append(s, lsb[i])
		s = append(s, msb[i])
		table = append(table, s)
	}
	return table
}

// GenerateNECMessage returns NECMessage generated from the binary string
// if possible.
func GenerateNECMessage(message Message) (NECMessage, error) {
	if !message.IsValid {
		return NECMessage{}, errors.New("cannot parse: Message size invalid")
	}
	return NECMessage{chunks: utils.SplitStringToChunks(message.Data, 8)}, nil
}

func isValidMessage(message string, useCHFormat bool) bool {
	if useCHFormat {
		return len(message) == validChangHongMessageLength
	}
	return len(message)%8 == 0
}
