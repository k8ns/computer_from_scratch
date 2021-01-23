package component

type Enabler struct {
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
	E  *Pin
}

func NewEnabler() *Enabler {
	e := &Enabler{
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
		E:  &Pin{},
	}

	inputs := e.Inputs()
	outputs := e.Outputs()

	for i := 0; i < 8; i++ {
		andGate := NewANDGate()
		Connect(e.E, andGate.B)
		Connect(inputs[i], andGate.A)
		Connect(andGate.O, outputs[i])
	}

	return e
}

func (e *Enabler) Inputs() []*Pin {
	return []*Pin{e.I1, e.I2, e.I3, e.I4, e.I5, e.I6, e.I7, e.I8}
}

func (e *Enabler) Outputs() []*Pin {
	return []*Pin{e.O1, e.O2, e.O3, e.O4, e.O5, e.O6, e.O7, e.O8}
}
