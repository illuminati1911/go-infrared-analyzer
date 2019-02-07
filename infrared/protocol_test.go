package protocol

import (
	"testing"
)

func TestIsValidNECMessage(t *testing.T) {
	if isValidNECMessage("abc") {
		t.Error("NEC Message validation failure.")
	}
}

func TestGenerateNECMessage(t *testing.T) {
	testMessage1 := "011010101001111000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000111000"
	t1, _ := GenerateNECMessage(testMessage1)
	t.Error(t1.Bytes(LSB))

	testMessage2 := "011010100001111000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000011011000"
	t2, _ := GenerateNECMessage(testMessage2)
	t.Error(t2.Bytes(LSB))

	testMessageOff := "011010101001111000000000000000000000100000000011000000000000000000000000000000000000000000000000000000000000000000010100"
	to, _ := GenerateNECMessage(testMessageOff)
	t.Error(to.Bytes(LSB))

	y := "011010101010111000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000011000"
	ty, _ := GenerateNECMessage(y)
	t.Error(ty.Bytes(LSB))

	x := "011010101010111000000000000000000000100000000011000000000000000000000000000000000000000000000000000000000000000000100100"
	tx, _ := GenerateNECMessage(x)
	t.Error(tx.Bytes(LSB))
}

func TestReverse(t *testing.T) {
	rText := reverse("abc123")
	if rText != "321cba" {
		t.Error("String reverse failure")
	}
}