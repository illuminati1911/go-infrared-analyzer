package protocol

import (
	"errors"
	"strconv"
)

const validNECMessageLength int = 120

type NECMessage struct {
	chunks []string
}

type BitOrder bool

const (
	LSB BitOrder = true
	MSB BitOrder = false
)

func (m NECMessage) Bytes(b BitOrder) []string {
	bytes := make([]string, 0)
	for _, chunk := range m.chunks {
		orderedChunk := chunk
		if b == LSB {
			orderedChunk = reverse(orderedChunk)
		}
		n, err := strconv.ParseUint(orderedChunk, 2, 8)
		if err == nil {
			bytes = append(bytes, strconv.FormatUint(n, 10))
		}
	}
	return bytes
}

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

func reverse(s string) string {
	rs := ""
	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}
	return rs
}

func GenerateNECMessage(message string) (NECMessage, error) {
	if !isValidNECMessage(message) {
		return NECMessage{}, errors.New("cannot parse: Message size invalid")
	}
	return NECMessage{chunks: splitStringToChunks(message, 8)}, nil
}

func isValidNECMessage(message string) bool {
	return len(message) == validNECMessageLength
}

func splitStringToChunks(s string, size uint) []string {
	if uint(len(s)) < size {
		return []string{s}
	}
	return append([]string{s[:size]}, splitStringToChunks(s[size:], size)...)
}
