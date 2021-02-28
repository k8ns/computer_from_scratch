package component

import "github.com/k8ns/computer_from_scratch/wire"

type Byte struct {
	I1 *wire.Pin
	I2 *wire.Pin
	I3 *wire.Pin
	I4 *wire.Pin
	I5 *wire.Pin
	I6 *wire.Pin
	I7 *wire.Pin
	I8 *wire.Pin
	O1 *wire.Pin
	O2 *wire.Pin
	O3 *wire.Pin
	O4 *wire.Pin
	O5 *wire.Pin
	O6 *wire.Pin
	O7 *wire.Pin
	O8 *wire.Pin
	S  *wire.Pin
	m1 *Memory
	m2 *Memory
	m3 *Memory
	m4 *Memory
	m5 *Memory
	m6 *Memory
	m7 *Memory
	m8 *Memory
}

func NewByte() *Byte {
	b := &Byte{
		I1: &wire.Pin{},
		I2: &wire.Pin{},
		I3: &wire.Pin{},
		I4: &wire.Pin{},
		I5: &wire.Pin{},
		I6: &wire.Pin{},
		I7: &wire.Pin{},
		I8: &wire.Pin{},
		O1: &wire.Pin{},
		O2: &wire.Pin{},
		O3: &wire.Pin{},
		O4: &wire.Pin{},
		O5: &wire.Pin{},
		O6: &wire.Pin{},
		O7: &wire.Pin{},
		O8: &wire.Pin{},
		S:  &wire.Pin{},
	}

	inputs := []*wire.Pin{b.I1, b.I2, b.I3, b.I4, b.I5, b.I6, b.I7, b.I8}
	outputs := b.Outputs()

	for i := 0; i < 8; i++ {
		m := NewMemory()
		wire.Connect(b.S, m.S)
		wire.Connect(inputs[i], m.I)
		wire.Connect(m.O, outputs[i])
	}

	return b
}

func (b *Byte) Outputs() []*wire.Pin {
	return []*wire.Pin{b.O1, b.O2, b.O3, b.O4, b.O5, b.O6, b.O7, b.O8}
}

func (b *Byte) Inputs() []*wire.Pin {
	return []*wire.Pin{b.I1, b.I2, b.I3, b.I4, b.I5, b.I6, b.I7, b.I8}
}
