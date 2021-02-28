package component

import (
	"github.com/k8ns/computer_from_scratch/gate"
	"github.com/k8ns/computer_from_scratch/wire"
)

type RAM struct {
	SA *wire.Pin
	S  *wire.Pin
	E  *wire.Pin
}

func NewRam16(bus *Bus) *RAM {
	ram := &RAM{
		SA: &wire.Pin{},
		S:  &wire.Pin{},
		E:  &wire.Pin{},
	}

	addressRegister := NewRegister()
	bus.ConnectRegisterInputs(addressRegister)
	wire.Connect(ram.SA, addressRegister.S)
	addressRegister.E.Voltage(true)

	d1 := NewDecoder2x4()
	d2 := NewDecoder2x4()

	wire.Connect(addressRegister.O1, d1.I1)
	wire.Connect(addressRegister.O2, d1.I2)
	wire.Connect(addressRegister.O3, d2.I1)
	wire.Connect(addressRegister.O4, d2.I2)

	for _, dOut1 := range d1.Outputs() {
		for _, dOut2 := range d2.Outputs() {
			addressGate := gate.NewANDGate()
			wire.Connect(dOut1, addressGate.A)
			wire.Connect(dOut2, addressGate.B)

			enableGate := gate.NewANDGate()
			setGate := gate.NewANDGate()
			wire.Connect(addressGate.O, enableGate.A, setGate.A)
			wire.Connect(ram.E, enableGate.B)
			wire.Connect(ram.S, setGate.B)

			register := NewRegister()
			wire.Connect(enableGate.O, register.E)
			wire.Connect(setGate.O, register.S)

			bus.ConnectRegister(register)
		}
	}

	return ram
}
