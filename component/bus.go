package component

import "github.com/k8ns/computer_from_scratch/wire"

type BusInput struct {
	I1 *wire.Pin
	I2 *wire.Pin
	I3 *wire.Pin
	I4 *wire.Pin
	I5 *wire.Pin
	I6 *wire.Pin
	I7 *wire.Pin
	I8 *wire.Pin
}

func NewBusInput() *BusInput {
	return &BusInput{
		I1: &wire.Pin{},
		I2: &wire.Pin{},
		I3: &wire.Pin{},
		I4: &wire.Pin{},
		I5: &wire.Pin{},
		I6: &wire.Pin{},
		I7: &wire.Pin{},
		I8: &wire.Pin{},
	}
}

func (i *BusInput) Inputs() []*wire.Pin {
	return []*wire.Pin{i.I1, i.I2, i.I3, i.I4, i.I5, i.I6, i.I7, i.I8}
}

func (i *BusInput) Set(bits string) {
	for n, bit := range StringToByte(bits) {
		i.Inputs()[n].Voltage(bit)
	}
}

func (i *BusInput) Reset() {
	for _, input := range i.Inputs() {
		input.Voltage(false)
	}
}

type Bus struct {
	W1 *wire.Wire
	W2 *wire.Wire
	W3 *wire.Wire
	W4 *wire.Wire
	W5 *wire.Wire
	W6 *wire.Wire
	W7 *wire.Wire
	W8 *wire.Wire
}

func NewBus() *Bus {
	b := &Bus{
		W1: &wire.Wire{},
		W2: &wire.Wire{},
		W3: &wire.Wire{},
		W4: &wire.Wire{},
		W5: &wire.Wire{},
		W6: &wire.Wire{},
		W7: &wire.Wire{},
		W8: &wire.Wire{},
	}
	return b
}

func (b *Bus) Wires() []*wire.Wire {
	return []*wire.Wire{b.W1, b.W2, b.W3, b.W4, b.W5, b.W6, b.W7, b.W8}
}

func (b *Bus) Connect(pins ...*wire.Pin) {
	for i, p := range pins {
		b.Wires()[i].Add(p)
	}
}

func (b *Bus) ConnectRegister(registers ...*Register) {
	for _, r := range registers {
		b.Connect(r.Inputs()...)
		b.Connect(r.Outputs()...)
	}
}

func (b *Bus) ConnectRegisterInputs(registers ...*Register) {
	for _, r := range registers {
		b.Connect(r.Inputs()...)
	}
}

func (b *Bus) Get() []bool {
	bits := make([]bool, 0)
	for _, w := range b.Wires() {
		bits = append(bits, w.Get())
	}
	return bits
}
