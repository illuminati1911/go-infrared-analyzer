package gpio

import (
	"github.com/stianeikeland/go-rpio"
)

type GPIO struct {
	pin rpio.Pin
}

func (g GPIO) Read() uint8 {
	return uint8(g.pin.Read())
}

func Open(busNumber uint8) (error, GPIO) {
	err := rpio.Open()
	if err != nil {
		return err, GPIO{}
	}
	pin := rpio.Pin(busNumber)
	pin.Mode(rpio.Input)
	return nil, GPIO{pin: pin}
}

func Close() {
	rpio.Close()
}
