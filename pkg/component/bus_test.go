package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBusCanConnectToRegister(t *testing.T) {
	r := NewRegister()
	bus := NewBus()
	bus.ConnectRegister(r)

	busInput := NewBusInput()
	bus.Connect(busInput.Inputs()...)

	msg := "10100110"
	for i, bit := range StringToByte(msg) {
		busInput.Inputs()[i].Voltage(bit)
	}
	r.S.Voltage(true)
	r.S.Voltage(false)
	busInput.Reset()

	//chech there is nothing on the bus
	assert.Equal(t, "00000000", ByteToString(bus.Get()))

	r.E.Voltage(true)
	assert.Equal(t, msg, ByteToString(bus.Get()))

	r.E.Voltage(false)
	assert.Equal(t, "00000000", ByteToString(bus.Get()))
}

func TestBusCanTransmitDataFromOneRegisterToAnother(t *testing.T) {
	r1 := NewRegister()
	r2 := NewRegister()
	bus := NewBus()
	bus.ConnectRegister(r1, r2)
	busInput := NewBusInput()
	bus.Connect(busInput.Inputs()...)

	msg := "10100110"
	for i, bit := range StringToByte(msg) {
		busInput.Inputs()[i].Voltage(bit)
	}
	r1.S.Voltage(true)
	r1.S.Voltage(false)
	busInput.Reset()

	r1.E.Voltage(true)
	r2.S.Voltage(true)
	r2.S.Voltage(false)
	r1.E.Voltage(false)

	r2.E.Voltage(true)
	assert.Equal(t, msg, ByteToString(bus.Get()))

	r2.E.Voltage(false)
	assert.Equal(t, "00000000", ByteToString(bus.Get()))

}
