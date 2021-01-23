package component

type Byte struct {
	I1 *Pin
	I2 *Pin
	I3 *Pin
	I4 *Pin
	I5 *Pin
	I6 *Pin
	I7 *Pin
	I8 *Pin
	O1 *Pin
	O2 *Pin
	O3 *Pin
	O4 *Pin
	O5 *Pin
	O6 *Pin
	O7 *Pin
	O8 *Pin
	S  *Pin
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
		I1: &Pin{},
		I2: &Pin{},
		I3: &Pin{},
		I4: &Pin{},
		I5: &Pin{},
		I6: &Pin{},
		I7: &Pin{},
		I8: &Pin{},
		O1: &Pin{},
		O2: &Pin{},
		O3: &Pin{},
		O4: &Pin{},
		O5: &Pin{},
		O6: &Pin{},
		O7: &Pin{},
		O8: &Pin{},
		S:  &Pin{},
	}

	inputs := []*Pin{b.I1, b.I2, b.I3, b.I4, b.I5, b.I6, b.I7, b.I8}
	outputs := b.Outputs()

	for i := 0; i < 8; i++ {
		m := NewMemory()
		Connect(b.S, m.S)
		Connect(inputs[i], m.I)
		Connect(m.O, outputs[i])
	}

	return b
}

func (b *Byte) Outputs() []*Pin {
	return []*Pin{b.O1, b.O2, b.O3, b.O4, b.O5, b.O6, b.O7, b.O8}
}

func (b *Byte) Inputs() []*Pin {
	return []*Pin{b.I1, b.I2, b.I3, b.I4, b.I5, b.I6, b.I7, b.I8}
}
