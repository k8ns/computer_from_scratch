package component

type Decoder2x4 struct {
	I1 *Pin
	I2 *Pin
	O1 *Pin
	O2 *Pin
	O3 *Pin
	O4 *Pin
}

func NewDecoder2x4() *Decoder2x4 {
	d := &Decoder2x4{
		I1: &Pin{},
		I2: &Pin{},
		O1: &Pin{},
		O2: &Pin{},
		O3: &Pin{},
		O4: &Pin{},
	}

	not1 := NewNOTGate()
	not2 := NewNOTGate()
	and1 := NewANDGate()
	and2 := NewANDGate()
	and3 := NewANDGate()
	and4 := NewANDGate()

	Connect(d.I1, not1.I, and3.A, and4.A)
	Connect(d.I2, not2.I, and2.B, and4.B)
	Connect(not1.O, and1.A, and2.A)
	Connect(not2.O, and1.B, and3.B)
	Connect(and1.O, d.O1)
	Connect(and2.O, d.O2)
	Connect(and3.O, d.O3)
	Connect(and4.O, d.O4)

	return d
}

func (d *Decoder2x4) Outputs() []*Pin {
	return []*Pin{d.O1, d.O2, d.O3, d.O4}
}
